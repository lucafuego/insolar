/*
 *    Copyright 2018 Insolar
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

package hostnetwork

import (
	"testing"

	"github.com/insolar/insolar/network/hostnetwork/host"
	"github.com/insolar/insolar/network/hostnetwork/id"
	"github.com/insolar/insolar/network/hostnetwork/packet"
)

func TestParseIncomingPacket(t *testing.T) {
	hh := newMockHostHandler()
	builder := packet.NewBuilder()

	pckt := builder.Type(packet.TypeStore).Request(&packet.RequestDataStore{}).Build()
	ParseIncomingPacket(hh, getDefaultCtx(hh), pckt, builder)
}

func TestBuildContext(t *testing.T) {
	hh := newMockHostHandler()
	cb := NewContextBuilder(hh)
	senderAddress, _ := host.NewAddress("0.0.0.0:0")
	sender := host.NewHost(senderAddress)
	sender.ID, _ = id.NewID()
	receiverAddress, _ := host.NewAddress("0.0.0.0:0")
	receiver := host.NewHost(receiverAddress)
	pckt := packet.NewAuthPacket(packet.Unknown, sender, receiver)
	_ = BuildContext(cb, pckt)

	receiver.ID, _ = id.NewID()
	pckt = packet.NewAuthPacket(packet.Unknown, sender, receiver)
	_ = BuildContext(cb, pckt)
}
