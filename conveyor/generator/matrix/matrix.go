/*
*    Copyright 2019 Insolar Technologies
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

package matrix

import (
	"github.com/insolar/insolar/conveyor/interfaces/statemachine"
	"github.com/insolar/insolar/conveyor/generator/state_machines/get_object"
	
)

type StateMachineSet struct{
	stateMachines []statemachine.StateMachine
}

func newStateMachineSet() *StateMachineSet {
	return &StateMachineSet{
		stateMachines: make([]statemachine.StateMachine, 1),
	}
}

func (s *StateMachineSet) addMachine(machine statemachine.StateMachine) {
	s.stateMachines = append(s.stateMachines, machine)
}

func ( s *StateMachineSet ) GetStateMachineByID(id int) statemachine.StateMachine{
	return s.stateMachines[id]
}

type Matrix struct {
	future *StateMachineSet
	present *StateMachineSet
	past *StateMachineSet
}

type MachineType int

const (
	GetObjectStateMachine MachineType = iota + 1
	
)

func NewMatrix() *Matrix {
	m := Matrix{
		future: newStateMachineSet(),
		present: newStateMachineSet(),
		past: newStateMachineSet(),
	}

	m.future.addMachine(getobject.RawGetObjectStateMachineFutureFactory())
	m.present.addMachine(getobject.RawGetObjectStateMachinePresentFactory())
	m.past.addMachine(getobject.RawGetObjectStateMachinePastFactory())
	
	return &m
}

func (m *Matrix) GetInitialStateMachine() statemachine.StateMachine {
	return m.present.stateMachines[1]
}

func (m *Matrix) GetFutureConfig() statemachine.SetAccessor{
	return m.future
}

func (m *Matrix) GetPresentConfig() statemachine.SetAccessor{
	return m.present
}

func (m *Matrix) GetPastConfig() statemachine.SetAccessor{
	return m.past
}
