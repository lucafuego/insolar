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

package blob

import (
	"bytes"
	"context"

	"github.com/insolar/insolar/insolar"
	"github.com/ugorji/go/codec"
)

//go:generate minimock -i github.com/insolar/insolar/ledger/storage/blob.Accessor -o ./ -s _mock.go

// Accessor provides info about Blob-values from storage.
type Accessor interface {
	// ForID returns Blob for a provided id.
	ForID(ctx context.Context, id insolar.ID) (Blob, error)
}

//go:generate minimock -i github.com/insolar/insolar/ledger/storage/blob.SyncAccessor -o ./ -s _mock.go

// SyncAccessor provides methods for querying blobs with specific search conditions.
type SyncAccessor interface {
	// ForPN returns []Blob for a provided jetID and a pulse number.
	ForPN(ctx context.Context, jetID insolar.JetID, pn insolar.PulseNumber) []Blob
}

//go:generate minimock -i github.com/insolar/insolar/ledger/storage/blob.Modifier -o ./ -s _mock.go

// Modifier provides methods for setting Blob-values to storage.
type Modifier interface {
	// Set saves new Blob-value in storage.
	Set(ctx context.Context, id insolar.ID, blob Blob) error
}

//go:generate minimock -i github.com/insolar/insolar/ledger/storage/blob.Cleaner -o ./ -s _mock.go

// Cleaner provides an interface for removing blobs from a storage.
type Cleaner interface {
	Delete(ctx context.Context, pulse insolar.PulseNumber)
}

// Blob represents blob-value with jetID.
type Blob struct {
	Value []byte
	JetID insolar.JetID
}

// Clone returns copy of argument blob.
func Clone(blob Blob) Blob {
	if len(blob.Value) == 0 {
		blob.Value = nil
	} else {
		b := blob.Value
		blob.Value = append([]byte(nil), b...)
	}

	return blob
}

// MustEncode serializes a blob.
func MustEncode(blob *Blob) []byte {
	var buf bytes.Buffer
	enc := codec.NewEncoder(&buf, &codec.CborHandle{})
	err := enc.Encode(blob)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Decode deserializes a blob.
func Decode(buf []byte) (*Blob, error) {
	dec := codec.NewDecoder(bytes.NewReader(buf), &codec.CborHandle{})
	var blob Blob
	err := dec.Decode(&blob)
	if err != nil {
		return nil, err
	}
	return &blob, nil
}
