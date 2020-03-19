// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/insolar/blob/master/LICENSE.md.

package configuration

// LogicRunner configuration
type LogicRunner struct {
	// BuiltIn - configuration of builtin executor
	BuiltIn *BuiltIn
	// PulseLRUSize - configuration of size of a pulse's cache
	PulseLRUSize int
}

// BuiltIn configuration, no options at the moment
type BuiltIn struct{}

// NewLogicRunner - returns default config of the logic runner
func NewLogicRunner() LogicRunner {
	return LogicRunner{
		BuiltIn:      &BuiltIn{},
		PulseLRUSize: 100,
	}
}
