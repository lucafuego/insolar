//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/insolar/insolar/api/requester"
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/testutils"
)

const CallUrl = "http://localhost:19192/api/call"

type TimeoutSuite struct {
	suite.Suite

	mc    *minimock.Controller
	ctx   context.Context
	api   *Runner
	user  *requester.UserConfigJSON
	delay chan struct{}
}

func (suite *TimeoutSuite) TestRunner_callHandler_NoTimeout() {
	seed, err := suite.api.SeedGenerator.Next()
	suite.NoError(err)
	suite.api.SeedManager.Add(*seed)

	close(suite.delay)
	suite.api.cfg.Timeout = 60

	seeds := base64.StdEncoding.EncodeToString(seed[:])

	resp, err := requester.SendWithSeed(
		suite.ctx,
		CallUrl,
		suite.user,
		&requester.Request{},
		seeds,
	)
	suite.NoError(err)

	var result requester.ContractAnswer
	err = json.Unmarshal(resp, &result)
	suite.NoError(err)
	suite.Equal("", result.Error.Message)
	suite.Equal("OK", result.Result.ContractResult)
}

func (suite *TimeoutSuite) TestRunner_callHandler_Timeout() {
	seed, err := suite.api.SeedGenerator.Next()
	suite.NoError(err)
	suite.api.SeedManager.Add(*seed)

	suite.api.cfg.Timeout = 1

	resp, err := requester.SendWithSeed(
		suite.ctx,
		CallUrl,
		suite.user,
		&requester.Request{},
		string(seed[:]),
	)
	suite.NoError(err)

	close(suite.delay)

	var result requester.ContractAnswer
	err = json.Unmarshal(resp, &result)
	suite.NoError(err)
	suite.Equal("Messagebus timeout exceeded", result.Error.Message)
	suite.Equal("", result.Result.ContractResult)
}

func TestTimeoutSuite(t *testing.T) {
	timeoutSuite := new(TimeoutSuite)
	timeoutSuite.ctx, _ = inslogger.WithTraceField(context.Background(), "APItests")
	timeoutSuite.mc = minimock.NewController(t)

	ks := platformpolicy.NewKeyProcessor()
	sKey, err := ks.GeneratePrivateKey()
	require.NoError(t, err)
	sKeyString, err := ks.ExportPrivateKeyPEM(sKey)
	require.NoError(t, err)
	pKey := ks.ExtractPublicKey(sKey)
	pKeyString, err := ks.ExportPublicKeyPEM(pKey)
	require.NoError(t, err)

	userRef := testutils.RandomRef().String()
	timeoutSuite.user, err = requester.CreateUserConfig(userRef, string(sKeyString))

	http.DefaultServeMux = new(http.ServeMux)
	cfg := configuration.NewAPIRunner()
	cfg.Address = "localhost:19192"
	cfg.Timeout = 1
	timeoutSuite.api, err = NewRunner(&cfg)
	require.NoError(t, err)

	cert := testutils.NewCertificateMock(timeoutSuite.mc)
	cert.GetRootDomainReferenceFunc = func() (r *insolar.Reference) {
		ref := testutils.RandomRef()
		return &ref
	}

	cm := testutils.NewCertificateManagerMock(timeoutSuite.mc)
	cm.GetCertificateFunc = func() (r insolar.Certificate) {
		return cert
	}

	// TODO: refactor this mock
	cr := testutils.NewContractRequesterMock(timeoutSuite.mc)
	cr.SendRequestFunc = func(p context.Context, p1 *insolar.Reference, method string, p3 []interface{}) (insolar.Reply, error) {
		switch method {
		case "GetPublicKey":
			var result = string(pKeyString)
			var contractErr *foundation.Error
			data, _ := insolar.MarshalArgs(result, contractErr)
			return &reply.CallMethod{
				Result: data,
			}, nil
		default:
			<-timeoutSuite.delay
			var result = "OK"
			var contractErr *foundation.Error
			data, _ := insolar.MarshalArgs(result, contractErr)
			return &reply.CallMethod{
				Result: data,
			}, nil
		}
	}

	timeoutSuite.api.ContractRequester = cr
	timeoutSuite.api.CertificateManager = cm
	timeoutSuite.api.Start(timeoutSuite.ctx)

	requester.SetTimeout(25)
	suite.Run(t, timeoutSuite)

	timeoutSuite.api.Stop(timeoutSuite.ctx)
}

func (suite *TimeoutSuite) BeforeTest(suiteName, testName string) {
	suite.delay = make(chan struct{}, 0)
}

func (suite *TimeoutSuite) AfterTest(suiteName, testName string) {
	suite.mc.Wait(1*time.Minute)
	suite.mc.Finish()
}
