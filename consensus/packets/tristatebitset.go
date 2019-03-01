/*
 * The Clear BSD License
 *
 * Copyright (c) 2019 Insolar Technologies
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted (subject to the limitations in the disclaimer below) provided that the following conditions are met:
 *
 *  Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *  Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *  Neither the name of Insolar Technologies nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package packets

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/bits"

	"github.com/pkg/errors"
)

const lowLengthSize = 6
const firstBitMask = 0x80
const last6BitsMask = 0x3f
const lastTwoBitsMask = 0x3

// TriStateBitSet bitset implementation.
type TriStateBitSet struct {
	CompressedSet bool
	array         bitArray
}

func (dbs *TriStateBitSet) GetCells(mapper BitSetMapper) ([]BitSetCell, error) {
	cells := make([]BitSetCell, len(dbs.array))
	for i := 0; i < len(dbs.array); i++ {
		id, err := mapper.IndexToRef(i)
		if err != nil {
			return nil, err
		}
		cells[i] = BitSetCell{NodeID: id, State: dbs.array[i]}
	}
	return cells, nil
}

// NewTriStateBitSet creates and returns a tristatebitset.
func NewTriStateBitSet(size int) (*TriStateBitSet, error) {
	bitset := &TriStateBitSet{
		array: make(bitArray, size),
	}
	for i := 0; i < size; i++ {
		bitset.array[i] = TimedOut
	}
	return bitset, nil
}

func (dbs *TriStateBitSet) GetTristateArray() ([]TriState, error) {
	result := make([]TriState, len(dbs.array))
	copy(result, dbs.array)
	return result, nil
}

func (dbs *TriStateBitSet) ApplyChanges(changes []BitSetCell, mapper BitSetMapper) error {
	for _, cell := range changes {
		index, err := mapper.RefToIndex(cell.NodeID)
		if err != nil {
			return errors.Wrap(err, "[ ApplyChanges ] failed to get index from ref")
		}
		dbs.array[index] = cell.State
	}
	return nil
}

func (dbs *TriStateBitSet) Serialize() ([]byte, error) {
	var firstByte uint8 // compressed and hBitLength bits
	if dbs.CompressedSet {
		firstByte = 0x01
	} else {
		firstByte = 0x00
	}

	data, err := dbs.array.Serialize(dbs.CompressedSet)
	if err != nil {
		return nil, errors.Wrap(err, "[ Serialize ] failed to serialize a bitarray")
	}

	length := len(dbs.array)
	var result bytes.Buffer
	firstByte = firstByte << 1
	if bits.Len(uint(length)) > lowLengthSize {
		err = dbs.serializeWithHLength(firstByte, length, &result)
		if err != nil {
			return nil, errors.Wrap(err, "[ Serialize ] failed to serialize first bytes")
		}
	} else {
		err = dbs.serializeWithLLength(firstByte, length, &result)
		if err != nil {
			return nil, errors.Wrap(err, "[ Serialize ] failed to serialize first bytes")
		}
	}

	err = binary.Write(&result, defaultByteOrder, data)
	if err != nil {
		return nil, errors.Wrap(err, "[ Serialize ] failed to write binary")
	}

	return result.Bytes(), nil
}

func (dbs *TriStateBitSet) serializeWithHLength(firstByte uint8, length int, result *bytes.Buffer) error {
	var secondByte uint8 // hBitLength
	firstByte++
	firstByte = firstByte << lowLengthSize // move compressed and hBitLength bits to right
	secondByte = uint8(length & 0xff)
	lowByte := uint8(length >> 8)
	if lowByte != 0 {
		lowByte &= last6BitsMask
		firstByte |= lowByte
	}
	err := binary.Write(result, defaultByteOrder, firstByte)
	if err != nil {
		return errors.Wrap(err, "[ serializeWithHLength ] failed to write binary")
	}
	err = binary.Write(result, defaultByteOrder, secondByte)
	if err != nil {
		return errors.Wrap(err, "[ serializeWithHLength ] failed to write binary")
	}
	return nil
}

func (dbs *TriStateBitSet) serializeWithLLength(firstByte uint8, length int, result *bytes.Buffer) error {
	firstByte = firstByte << lowLengthSize // move compressed and hbit flags to right
	firstByte += uint8(length)
	err := binary.Write(result, defaultByteOrder, firstByte)
	if err != nil {
		return errors.Wrap(err, "[ serializeWithLLength ] failed to write binary")
	}
	return nil
}

func DeserializeBitSet(data io.Reader) (BitSet, error) {
	firstbyte := uint8(0)
	err := binary.Read(data, defaultByteOrder, &firstbyte)
	if firstbyte == 0 {
		return nil, errors.New("[ DeserializeBitSet ] failed to deserialize: wrong data")
	}

	var array bitArray
	if err != nil {
		return nil, errors.Wrap(err, "[ Deserialize ] failed to read first byte")
	}
	compressed, hbitFlag, lowLength := parseFirstByte(firstbyte)
	var length int
	if hbitFlag {
		var highLength uint8
		err = binary.Read(data, defaultByteOrder, &highLength)
		if err != nil {
			return nil, errors.Wrap(err, "[ Deserialize ] failed to read second byte")
		}
		length = int(lowLength)<<8 | int(highLength)
	} else {
		length = int(lowLength)
	}
	if compressed {
		array, err = deserializeCompressed(data, length)
		if err != nil {
			return nil, errors.Wrap(err, "[ DeserializeBitSet ] failed to deserialize a compressed bitarray")
		}
	} else {
		payload := make([]uint8, div(length, statesInByte))
		err := binary.Read(data, defaultByteOrder, &payload)
		if err != nil {
			return nil, errors.Wrap(err, "[ Deserialize ] failed to read payload")
		}
		array, err = deserialize(payload, length)
		if err != nil {
			return nil, errors.Wrap(err, "[ Deserialize ] failed to parse a bitarray")
		}
	}
	bitset := &TriStateBitSet{
		array: array,
	}
	return bitset, nil
}

func parseFirstByte(b uint8) (compressed bool, hbitFlag bool, lbitLength uint8) {
	lbitLength = uint8(0)
	compressed = false
	hbitFlag = false
	if (b & firstBitMask) == firstBitMask { // check compressed flag bit
		compressed = true
	}
	check := (b << 1) & firstBitMask // check hBitLength flag bit
	if check == firstBitMask {
		hbitFlag = true
	}
	lbitLength = (b << 2) >> 2 // remove 2 first bits
	return
}
