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

package message

import (
	"io"

	"github.com/insolar/insolar/core"
)

// GetObjectMessage is a message for calling constructor and obtain its response
type GetObjectMessage struct {
	baseMessage
	Object core.RecordRef
}

// GetOperatingRole returns operating jet role for given message type.
func (m *GetObjectMessage) GetOperatingRole() core.JetRole {
	return core.RoleLightExecutor
}

// GetReference returns referenced object.
func (m *GetObjectMessage) GetReference() core.RecordRef {
	return m.Object
}

// Serialize serializes message.
func (m *GetObjectMessage) Serialize() (io.Reader, error) {
	return serialize(m, GetObjectMessageType)
}
