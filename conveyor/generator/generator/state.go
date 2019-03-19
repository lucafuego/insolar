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

package generator

type state struct {
	Name                       string
	Transition                 *handler
	TransitionFuture           *handler
	TransitionPast             *handler
	Migration                  *handler
	MigrationFuturePresent     *handler
	AdapterResponse            *handler
	AdapterResponseFuture      *handler
	AdapterResponsePast        *handler
	ErrorState                 *handler
	ErrorStateFuture           *handler
	ErrorStatePast             *handler
	AdapterResponseError       *handler
	AdapterResponseErrorFuture *handler
	AdapterResponseErrorPast   *handler
}

func (s *state) GetTransitionName() string {
	return s.Transition.Name
}

func (s *state) GetTransitionFutureName() string {
	if s.TransitionFuture != nil {
		return s.TransitionFuture.Name
	}
	return s.Transition.Name
}

func (s *state) GetTransitionPastName() string {
	if s.TransitionPast != nil {
		return s.TransitionPast.Name
	}
	return s.Transition.Name
}

func (s *state) GetMigrationName() string {
	return s.Migration.Name
}

func (s *state) GetMigrationFuturePresentName() string {
	if s.MigrationFuturePresent != nil {
		return s.MigrationFuturePresent.Name
	}
	return s.Migration.Name
}

func (s *state) GetAdapterResponseName() string {
	return s.AdapterResponse.Name
}

func (s *state) GetAdapterResponseFutureName() string {
	if s.AdapterResponseFuture != nil {
		return s.AdapterResponseFuture.Name
	}
	return s.AdapterResponse.Name
}

func (s *state) GetAdapterResponsePastName() string {
	if s.AdapterResponsePast != nil {
		return s.AdapterResponsePast.Name
	}
	return s.AdapterResponse.Name
}

func (s *state) GetErrorStateName() string {
	return s.ErrorState.Name
}

func (s *state) GetErrorStateFutureName() string {
	if s.ErrorStateFuture != nil {
		return s.ErrorStateFuture.Name
	}
	return s.ErrorState.Name
}

func (s *state) GetErrorStatePastName() string {
	if s.ErrorStatePast != nil {
		return s.ErrorStatePast.Name
	}
	return s.ErrorState.Name
}

func (s *state) GetAdapterResponseErrorName() string {
	return s.AdapterResponseError.Name
}

func (s *state) GetAdapterResponseErrorFutureName() string {
	if s.AdapterResponseErrorFuture != nil {
		return s.AdapterResponseErrorFuture.Name
	}
	return s.AdapterResponseError.Name
}

func (s *state) GetAdapterResponseErrorPastName() string {
	if s.AdapterResponseErrorPast != nil {
		return s.AdapterResponseErrorPast.Name
	}
	return s.AdapterResponseError.Name
}
