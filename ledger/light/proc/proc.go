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

package proc

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/bus"
)

type Dependencies struct {
	FetchJet               func(*FetchJet)
	CheckJet               func(*CheckJet)
	WaitHot                func(*WaitHot)
	WaitHotWM              func(*WaitHotWM)
	GetIndex               func(*GetIndex)
	GetIndexWM             func(*GetIndexWM)
	SendObject             func(*SendObject)
	GetCode                func(*GetCode)
	GetRequest             func(*GetRequest)
	UpdateObject           func(*UpdateObject)
	SetBlob                func(*SetBlob)
	SetRecord              func(*SetRecord)
	RegisterChild          func(*RegisterChild)
	GetPendingRequests     func(*GetPendingRequests)
	GetPendingRequestID    func(*GetPendingRequestID)
	GetJet                 func(*GetJet)
	GetChildren            func(*GetChildren)
	HotData                func(*HotData)
	PassState              func(*PassState)
	CalculateID            func(*CalculateID)
	SetCode                func(*SetCode)
	GetPendingFilament     func(*GetPendingFilament)
	RefreshPendingFilament func(*RefreshPendingFilament)
	Sender                 bus.Sender
}

type ReturnReply struct {
	Message *message.Message
	Err     error
	Reply   insolar.Reply
	Sender  bus.Sender
}

func (p *ReturnReply) Proceed(ctx context.Context) error {
	select {
	case <-ctx.Done():
	}
	var msg *message.Message
	if p.Err != nil {
		msg = bus.ErrorAsMessage(ctx, p.Err)
	} else {
		msg = bus.ReplyAsMessage(ctx, p.Reply)
	}
	p.Sender.Reply(ctx, p.Message, msg)
	return nil
}
