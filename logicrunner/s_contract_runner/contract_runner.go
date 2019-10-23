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

package s_contract_runner

import (
	"context"

	"github.com/insolar/insolar/conveyor/smachine"
	"github.com/insolar/insolar/logicrunner/machinesmanager"
)

type ContractCallType uint8

const (
	_ ContractCallType = iota
	ContractCallMutable
	ContractCallImmutable
	ContractCallSaga
)

type CallResult interface{}

type ContractRunnerService interface {
	// ClassifyCall(code artifacts.CodeDescriptor, method string) ContractCallType
	// CallImmutableMethod(code ArtifactBinary, method string, state ArtifactBinary) CallResult
}

type ContractRunnerServiceAdapter struct {
	svc  ContractRunnerService
	exec smachine.ExecutionAdapter
}

func (a *ContractRunnerServiceAdapter) PrepareSync(ctx smachine.ExecutionContext, fn func(svc ContractRunnerService)) smachine.SyncCallRequester {
	return a.exec.PrepareSync(ctx, func(_ interface{}) smachine.AsyncResultFunc {
		fn(a.svc)
		return nil
	})
}

func (a *ContractRunnerServiceAdapter) PrepareAsync(ctx smachine.ExecutionContext, fn func(svc ContractRunnerService) smachine.AsyncResultFunc) smachine.AsyncCallRequester {
	return a.exec.PrepareAsync(ctx, func(_ interface{}) smachine.AsyncResultFunc {
		return fn(a.svc)
	})
}

type contractRunnerService struct {
	MachinesManager machinesmanager.MachinesManager
}

func CreateContractRunnerService() *ContractRunnerServiceAdapter {
	ctx := context.Background()
	ae, ch := smachine.NewCallChannelExecutor(ctx, 0, false, 5)
	smachine.StartChannelWorker(ctx, ch, nil)

	return &ContractRunnerServiceAdapter{
		svc:  contractRunnerService{},
		exec: smachine.NewExecutionAdapter("ArtifactClientService", ae),
	}
}

// func (c contractRunnerService) CallMethod()
//
// func (c contractRunnerService) ClassifyCall(code artifacts.CodeDescriptor, method string) ContractCallType {
// 	panic("implement me")
// }
//
// func (c contractRunnerService) CallImmutableMethod(code ArtifactBinary, method string, state ArtifactBinary) CallResult {
// 	panic("implement me")
// }
//
// func (c contractRunnerService) CallMutableMethod(code ArtifactBinary)
