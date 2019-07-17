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

package proc_test

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/bus"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/light/executor"
	"github.com/insolar/insolar/ledger/light/hot"
	"github.com/insolar/insolar/ledger/light/proc"
	"github.com/insolar/insolar/ledger/object"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeactivateObject_RecordOverrideErr(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	recordsMock := object.NewRecordModifierMock(t)
	recordsMock.SetMock.Return(object.ErrOverride)

	idxStorage := object.NewIndexStorageMock(t)
	idxStorage.ForIDMock.Return(record.Index{}, nil)

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		recordsMock,
		idxStorage,
		nil,
		nil,
	)

	err := p.Proceed(ctx)
	// Since there is no deduplication yet it's quite possible that there will be
	// two writes by the same key. For this reason currently instead of reporting
	// an error we return OK (nil error). When deduplication will be implemented
	// we should check `ErrOverride` here.
	require.NoError(t, err)
}

func TestDeactivateObject_RecordErr(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	recordsMock := object.NewRecordModifierMock(t)
	recordsMock.SetMock.Return(errors.New("something strange from records.Set"))

	idxStorage := object.NewIndexStorageMock(t)
	idxStorage.ForIDMock.Return(record.Index{}, nil)

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		recordsMock,
		idxStorage,
		nil,
		nil,
	)

	err := p.Proceed(ctx)
	require.Error(t, err)
}

func TestDeactivateObject_IndexForIDErr(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	recordsMock := object.NewRecordModifierMock(t)
	recordsMock.SetMock.Return(nil)

	idxStorageMock := object.NewIndexStorageMock(t)
	idxStorageMock.ForIDMock.Return(record.Index{}, errors.New("something strange from index.ForID"))

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		recordsMock,
		idxStorageMock,
		nil,
		nil,
	)

	err := p.Proceed(ctx)
	require.Error(t, err)
	assert.Equal(t, "can't get index from storage: something strange from index.ForID", err.Error())
}

func TestDeactivateObject_SetIndexErr(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	recordsMock := object.NewRecordModifierMock(t)
	recordsMock.SetMock.Return(nil)

	idxStorageMock := object.NewIndexStorageMock(t)
	idxStorageMock.ForIDMock.Return(record.Index{}, nil)
	idxStorageMock.SetIndexMock.Return(errors.New("something strange from SetIndex"))

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		recordsMock,
		idxStorageMock,
		nil,
		nil,
	)

	err := p.Proceed(ctx)
	require.Error(t, err)
	assert.Equal(t, "something strange from SetIndex", err.Error())
}

func TestDeactivateObject_FilamentSetResultErr(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	recordsMock := object.NewRecordModifierMock(t)
	recordsMock.SetMock.Return(nil)

	idxStorageMock := object.NewIndexStorageMock(t)
	idxStorageMock.ForIDMock.Return(record.Index{}, nil)
	idxStorageMock.SetIndexMock.Return(nil)

	filaments := executor.NewFilamentModifierMock(t)
	filaments.SetResultMock.Return(nil, errors.New("something strange from filament.SetResult"))

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		recordsMock,
		idxStorageMock,
		filaments,
		nil,
	)

	err := p.Proceed(ctx)
	require.Error(t, err)
	assert.Equal(t, "failed to save result: something strange from filament.SetResult", err.Error())
}

func TestDeactivateObject_Proceed(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	recordsMock := object.NewRecordModifierMock(t)
	recordsMock.SetMock.Return(nil)

	idxStorageMock := object.NewIndexStorageMock(t)
	idxStorageMock.ForIDMock.Return(record.Index{}, nil)
	idxStorageMock.SetIndexMock.Return(nil)

	filaments := executor.NewFilamentModifierMock(t)
	filaments.SetResultMock.Return(nil, nil)

	sender := bus.NewSenderMock(t)
	sender.ReplyMock.Return()

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		recordsMock,
		idxStorageMock,
		filaments,
		sender,
	)

	err := p.Proceed(ctx)
	require.NoError(t, err)
}

func TestDeactivateObject_ObjectIsDeactivated(t *testing.T) {
	t.Parallel()

	ctx := flow.TestContextWithPulse(
		inslogger.TestContext(t),
		insolar.GenesisPulse.PulseNumber+10,
	)

	writeAccessor := hot.NewWriteAccessorMock(t)
	writeAccessor.BeginMock.Return(func() {}, nil)

	idxLockMock := object.NewIndexLockerMock(t)
	idxLockMock.LockMock.Return()
	idxLockMock.UnlockMock.Return()

	idxStorageMock := object.NewIndexStorageMock(t)
	idxStorageMock.ForIDMock.Return(record.Index{
		Lifeline: record.Lifeline{
			StateID: record.StateDeactivation,
		},
	}, nil)
	idxStorageMock.SetIndexMock.Return(nil)

	sender := bus.NewSenderMock(t)
	sender.ReplyFunc = func(_ context.Context, _ payload.Meta, inMsg *message.Message) {
		resp, err := payload.Unmarshal(inMsg.Payload)
		require.NoError(t, err)

		res, ok := resp.(*payload.Error)
		require.True(t, ok)
		require.Equal(t, payload.CodeDeactivated, int(res.Code))
	}

	p := proc.NewDeactivateObject(
		payload.Meta{},
		record.Deactivate{},
		gen.ID(),
		record.Result{},
		gen.ID(),
		gen.JetID(),
	)
	p.Dep(
		writeAccessor,
		idxLockMock,
		nil,
		idxStorageMock,
		nil,
		sender,
	)

	err := p.Proceed(ctx)
	require.NoError(t, err)
}
