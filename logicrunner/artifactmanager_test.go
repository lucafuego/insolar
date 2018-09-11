/*
 *    Copyright 2018 INS Ecosystem
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package logicrunner

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"

	"github.com/insolar/insolar/ledger/ledgertestutil"
	"github.com/insolar/insolar/logicrunner/goplugin"
	"github.com/insolar/insolar/logicrunner/goplugin/testutil"
)

func TestGoPlugin_Hello(t *testing.T) {
	l, cleaner := ledgertestutil.TmpLedger(t, "")
	defer cleaner()

	var helloCode = `
package main

import (
	"fmt"

	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Hello struct {
	foundation.BaseContract
}

func New() *Hello {
	return &Hello{};
}

func (b *Hello) String() string {
	return fmt.Sprint("Hello, Go is there!")
}
	`

	lr, err := NewLogicRunner(configuration.LogicRunner{})
	assert.NoError(t, err, "Initialize runner")

	am := l.GetManager()
	lr.ArtifactManager = am
	mr := &testMessageRouter{LogicRunner: lr}

	insiderStorage, err := ioutil.TempDir("", "test-")
	assert.NoError(t, err)
	defer os.RemoveAll(insiderStorage) // nolint: errcheck

	// assert.NoError(t, lr.Start(core.Components{
	// 	"core.Ledger":        l,
	// 	"core.MessageRouter": &testMessageRouter{},
	// }), "starting logicrunner")

	gp, err := goplugin.NewGoPlugin(
		&configuration.GoPlugin{
			MainListen:     "127.0.0.1:7778",
			RunnerListen:   "127.0.0.1:7777",
			RunnerPath:     "./goplugin/ginsider-cli/ginsider-cli",
			RunnerCodePath: insiderStorage,
		},
		mr,
		am,
	)
	assert.NoError(t, err)
	defer gp.Stop()

	err = lr.RegisterExecutor(core.MachineTypeGoPlugin, gp)
	assert.NoError(t, err)

	cb := testutil.NewContractBuilder(am, icc)
	err = cb.Build(map[string]string{"hello": helloCode})
	assert.NoError(t, err)

	_, res, err := gp.CallMethod(
		*cb.Codes["hello"],
		mustEncodeCBOR(t, &struct{}{}),
		"String",
		mustEncodeCBOR(t, []interface{}{}),
	)
	if err != nil {
		panic(err)
	}
	resParsed := mustDecodeCBOR(t, res)
	assert.Equal(t, "Hello, Go is there!", resParsed[0])
}

func mustDecodeCBOR(t *testing.T, in []byte) []interface{} {
	ch := new(codec.CborHandle)
	var resParsed []interface{}
	err := codec.NewDecoderBytes(in, ch).Decode(&resParsed)
	if err != nil {
		t.Fatal(err)
	}
	return resParsed
}

func mustEncodeCBOR(t *testing.T, in interface{}) []byte {
	ch := new(codec.CborHandle)
	var data []byte
	err := codec.NewEncoderBytes(&data, ch).Encode(in)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
