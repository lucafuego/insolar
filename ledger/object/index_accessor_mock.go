package object

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
)

// IndexAccessorMock implements IndexAccessor
type IndexAccessorMock struct {
	t minimock.Tester

	funcForID          func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (i1 record.Index, err error)
	inspectFuncForID   func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID)
	afterForIDCounter  uint64
	beforeForIDCounter uint64
	ForIDMock          mIndexAccessorMockForID

	funcForPulse          func(ctx context.Context, pn insolar.PulseNumber) (ia1 []record.Index, err error)
	inspectFuncForPulse   func(ctx context.Context, pn insolar.PulseNumber)
	afterForPulseCounter  uint64
	beforeForPulseCounter uint64
	ForPulseMock          mIndexAccessorMockForPulse

	funcLastKnownForID          func(ctx context.Context, objID insolar.ID) (i1 record.Index, err error)
	inspectFuncLastKnownForID   func(ctx context.Context, objID insolar.ID)
	afterLastKnownForIDCounter  uint64
	beforeLastKnownForIDCounter uint64
	LastKnownForIDMock          mIndexAccessorMockLastKnownForID
}

// NewIndexAccessorMock returns a mock for IndexAccessor
func NewIndexAccessorMock(t minimock.Tester) *IndexAccessorMock {
	m := &IndexAccessorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ForIDMock = mIndexAccessorMockForID{mock: m}
	m.ForIDMock.callArgs = []*IndexAccessorMockForIDParams{}

	m.ForPulseMock = mIndexAccessorMockForPulse{mock: m}
	m.ForPulseMock.callArgs = []*IndexAccessorMockForPulseParams{}

	m.LastKnownForIDMock = mIndexAccessorMockLastKnownForID{mock: m}
	m.LastKnownForIDMock.callArgs = []*IndexAccessorMockLastKnownForIDParams{}

	return m
}

type mIndexAccessorMockForID struct {
	mock               *IndexAccessorMock
	defaultExpectation *IndexAccessorMockForIDExpectation
	expectations       []*IndexAccessorMockForIDExpectation

	callArgs []*IndexAccessorMockForIDParams
	mutex    sync.RWMutex
}

// IndexAccessorMockForIDExpectation specifies expectation struct of the IndexAccessor.ForID
type IndexAccessorMockForIDExpectation struct {
	mock    *IndexAccessorMock
	params  *IndexAccessorMockForIDParams
	results *IndexAccessorMockForIDResults
	Counter uint64
}

// IndexAccessorMockForIDParams contains parameters of the IndexAccessor.ForID
type IndexAccessorMockForIDParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	objID insolar.ID
}

// IndexAccessorMockForIDResults contains results of the IndexAccessor.ForID
type IndexAccessorMockForIDResults struct {
	i1  record.Index
	err error
}

// Expect sets up expected params for IndexAccessor.ForID
func (mmForID *mIndexAccessorMockForID) Expect(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) *mIndexAccessorMockForID {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("IndexAccessorMock.ForID mock is already set by Set")
	}

	if mmForID.defaultExpectation == nil {
		mmForID.defaultExpectation = &IndexAccessorMockForIDExpectation{}
	}

	mmForID.defaultExpectation.params = &IndexAccessorMockForIDParams{ctx, pn, objID}
	for _, e := range mmForID.expectations {
		if minimock.Equal(e.params, mmForID.defaultExpectation.params) {
			mmForID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForID.defaultExpectation.params)
		}
	}

	return mmForID
}

// Inspect accepts an inspector function that has same arguments as the IndexAccessor.ForID
func (mmForID *mIndexAccessorMockForID) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID)) *mIndexAccessorMockForID {
	if mmForID.mock.inspectFuncForID != nil {
		mmForID.mock.t.Fatalf("Inspect function is already set for IndexAccessorMock.ForID")
	}

	mmForID.mock.inspectFuncForID = f

	return mmForID
}

// Return sets up results that will be returned by IndexAccessor.ForID
func (mmForID *mIndexAccessorMockForID) Return(i1 record.Index, err error) *IndexAccessorMock {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("IndexAccessorMock.ForID mock is already set by Set")
	}

	if mmForID.defaultExpectation == nil {
		mmForID.defaultExpectation = &IndexAccessorMockForIDExpectation{mock: mmForID.mock}
	}
	mmForID.defaultExpectation.results = &IndexAccessorMockForIDResults{i1, err}
	return mmForID.mock
}

//Set uses given function f to mock the IndexAccessor.ForID method
func (mmForID *mIndexAccessorMockForID) Set(f func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (i1 record.Index, err error)) *IndexAccessorMock {
	if mmForID.defaultExpectation != nil {
		mmForID.mock.t.Fatalf("Default expectation is already set for the IndexAccessor.ForID method")
	}

	if len(mmForID.expectations) > 0 {
		mmForID.mock.t.Fatalf("Some expectations are already set for the IndexAccessor.ForID method")
	}

	mmForID.mock.funcForID = f
	return mmForID.mock
}

// When sets expectation for the IndexAccessor.ForID which will trigger the result defined by the following
// Then helper
func (mmForID *mIndexAccessorMockForID) When(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) *IndexAccessorMockForIDExpectation {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("IndexAccessorMock.ForID mock is already set by Set")
	}

	expectation := &IndexAccessorMockForIDExpectation{
		mock:   mmForID.mock,
		params: &IndexAccessorMockForIDParams{ctx, pn, objID},
	}
	mmForID.expectations = append(mmForID.expectations, expectation)
	return expectation
}

// Then sets up IndexAccessor.ForID return parameters for the expectation previously defined by the When method
func (e *IndexAccessorMockForIDExpectation) Then(i1 record.Index, err error) *IndexAccessorMock {
	e.results = &IndexAccessorMockForIDResults{i1, err}
	return e.mock
}

// ForID implements IndexAccessor
func (mmForID *IndexAccessorMock) ForID(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (i1 record.Index, err error) {
	mm_atomic.AddUint64(&mmForID.beforeForIDCounter, 1)
	defer mm_atomic.AddUint64(&mmForID.afterForIDCounter, 1)

	if mmForID.inspectFuncForID != nil {
		mmForID.inspectFuncForID(ctx, pn, objID)
	}

	params := &IndexAccessorMockForIDParams{ctx, pn, objID}

	// Record call args
	mmForID.ForIDMock.mutex.Lock()
	mmForID.ForIDMock.callArgs = append(mmForID.ForIDMock.callArgs, params)
	mmForID.ForIDMock.mutex.Unlock()

	for _, e := range mmForID.ForIDMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmForID.ForIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForID.ForIDMock.defaultExpectation.Counter, 1)
		want := mmForID.ForIDMock.defaultExpectation.params
		got := IndexAccessorMockForIDParams{ctx, pn, objID}
		if want != nil && !minimock.Equal(*want, got) {
			mmForID.t.Errorf("IndexAccessorMock.ForID got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmForID.ForIDMock.defaultExpectation.results
		if results == nil {
			mmForID.t.Fatal("No results are set for the IndexAccessorMock.ForID")
		}
		return (*results).i1, (*results).err
	}
	if mmForID.funcForID != nil {
		return mmForID.funcForID(ctx, pn, objID)
	}
	mmForID.t.Fatalf("Unexpected call to IndexAccessorMock.ForID. %v %v %v", ctx, pn, objID)
	return
}

// ForIDAfterCounter returns a count of finished IndexAccessorMock.ForID invocations
func (mmForID *IndexAccessorMock) ForIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForID.afterForIDCounter)
}

// ForIDBeforeCounter returns a count of IndexAccessorMock.ForID invocations
func (mmForID *IndexAccessorMock) ForIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForID.beforeForIDCounter)
}

// Calls returns a list of arguments used in each call to IndexAccessorMock.ForID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForID *mIndexAccessorMockForID) Calls() []*IndexAccessorMockForIDParams {
	mmForID.mutex.RLock()

	argCopy := make([]*IndexAccessorMockForIDParams, len(mmForID.callArgs))
	copy(argCopy, mmForID.callArgs)

	mmForID.mutex.RUnlock()

	return argCopy
}

// MinimockForIDDone returns true if the count of the ForID invocations corresponds
// the number of defined expectations
func (m *IndexAccessorMock) MinimockForIDDone() bool {
	for _, e := range m.ForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForID != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockForIDInspect logs each unmet expectation
func (m *IndexAccessorMock) MinimockForIDInspect() {
	for _, e := range m.ForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IndexAccessorMock.ForID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		if m.ForIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IndexAccessorMock.ForID")
		} else {
			m.t.Errorf("Expected call to IndexAccessorMock.ForID with params: %#v", *m.ForIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForID != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		m.t.Error("Expected call to IndexAccessorMock.ForID")
	}
}

type mIndexAccessorMockForPulse struct {
	mock               *IndexAccessorMock
	defaultExpectation *IndexAccessorMockForPulseExpectation
	expectations       []*IndexAccessorMockForPulseExpectation

	callArgs []*IndexAccessorMockForPulseParams
	mutex    sync.RWMutex
}

// IndexAccessorMockForPulseExpectation specifies expectation struct of the IndexAccessor.ForPulse
type IndexAccessorMockForPulseExpectation struct {
	mock    *IndexAccessorMock
	params  *IndexAccessorMockForPulseParams
	results *IndexAccessorMockForPulseResults
	Counter uint64
}

// IndexAccessorMockForPulseParams contains parameters of the IndexAccessor.ForPulse
type IndexAccessorMockForPulseParams struct {
	ctx context.Context
	pn  insolar.PulseNumber
}

// IndexAccessorMockForPulseResults contains results of the IndexAccessor.ForPulse
type IndexAccessorMockForPulseResults struct {
	ia1 []record.Index
	err error
}

// Expect sets up expected params for IndexAccessor.ForPulse
func (mmForPulse *mIndexAccessorMockForPulse) Expect(ctx context.Context, pn insolar.PulseNumber) *mIndexAccessorMockForPulse {
	if mmForPulse.mock.funcForPulse != nil {
		mmForPulse.mock.t.Fatalf("IndexAccessorMock.ForPulse mock is already set by Set")
	}

	if mmForPulse.defaultExpectation == nil {
		mmForPulse.defaultExpectation = &IndexAccessorMockForPulseExpectation{}
	}

	mmForPulse.defaultExpectation.params = &IndexAccessorMockForPulseParams{ctx, pn}
	for _, e := range mmForPulse.expectations {
		if minimock.Equal(e.params, mmForPulse.defaultExpectation.params) {
			mmForPulse.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForPulse.defaultExpectation.params)
		}
	}

	return mmForPulse
}

// Inspect accepts an inspector function that has same arguments as the IndexAccessor.ForPulse
func (mmForPulse *mIndexAccessorMockForPulse) Inspect(f func(ctx context.Context, pn insolar.PulseNumber)) *mIndexAccessorMockForPulse {
	if mmForPulse.mock.inspectFuncForPulse != nil {
		mmForPulse.mock.t.Fatalf("Inspect function is already set for IndexAccessorMock.ForPulse")
	}

	mmForPulse.mock.inspectFuncForPulse = f

	return mmForPulse
}

// Return sets up results that will be returned by IndexAccessor.ForPulse
func (mmForPulse *mIndexAccessorMockForPulse) Return(ia1 []record.Index, err error) *IndexAccessorMock {
	if mmForPulse.mock.funcForPulse != nil {
		mmForPulse.mock.t.Fatalf("IndexAccessorMock.ForPulse mock is already set by Set")
	}

	if mmForPulse.defaultExpectation == nil {
		mmForPulse.defaultExpectation = &IndexAccessorMockForPulseExpectation{mock: mmForPulse.mock}
	}
	mmForPulse.defaultExpectation.results = &IndexAccessorMockForPulseResults{ia1, err}
	return mmForPulse.mock
}

//Set uses given function f to mock the IndexAccessor.ForPulse method
func (mmForPulse *mIndexAccessorMockForPulse) Set(f func(ctx context.Context, pn insolar.PulseNumber) (ia1 []record.Index, err error)) *IndexAccessorMock {
	if mmForPulse.defaultExpectation != nil {
		mmForPulse.mock.t.Fatalf("Default expectation is already set for the IndexAccessor.ForPulse method")
	}

	if len(mmForPulse.expectations) > 0 {
		mmForPulse.mock.t.Fatalf("Some expectations are already set for the IndexAccessor.ForPulse method")
	}

	mmForPulse.mock.funcForPulse = f
	return mmForPulse.mock
}

// When sets expectation for the IndexAccessor.ForPulse which will trigger the result defined by the following
// Then helper
func (mmForPulse *mIndexAccessorMockForPulse) When(ctx context.Context, pn insolar.PulseNumber) *IndexAccessorMockForPulseExpectation {
	if mmForPulse.mock.funcForPulse != nil {
		mmForPulse.mock.t.Fatalf("IndexAccessorMock.ForPulse mock is already set by Set")
	}

	expectation := &IndexAccessorMockForPulseExpectation{
		mock:   mmForPulse.mock,
		params: &IndexAccessorMockForPulseParams{ctx, pn},
	}
	mmForPulse.expectations = append(mmForPulse.expectations, expectation)
	return expectation
}

// Then sets up IndexAccessor.ForPulse return parameters for the expectation previously defined by the When method
func (e *IndexAccessorMockForPulseExpectation) Then(ia1 []record.Index, err error) *IndexAccessorMock {
	e.results = &IndexAccessorMockForPulseResults{ia1, err}
	return e.mock
}

// ForPulse implements IndexAccessor
func (mmForPulse *IndexAccessorMock) ForPulse(ctx context.Context, pn insolar.PulseNumber) (ia1 []record.Index, err error) {
	mm_atomic.AddUint64(&mmForPulse.beforeForPulseCounter, 1)
	defer mm_atomic.AddUint64(&mmForPulse.afterForPulseCounter, 1)

	if mmForPulse.inspectFuncForPulse != nil {
		mmForPulse.inspectFuncForPulse(ctx, pn)
	}

	params := &IndexAccessorMockForPulseParams{ctx, pn}

	// Record call args
	mmForPulse.ForPulseMock.mutex.Lock()
	mmForPulse.ForPulseMock.callArgs = append(mmForPulse.ForPulseMock.callArgs, params)
	mmForPulse.ForPulseMock.mutex.Unlock()

	for _, e := range mmForPulse.ForPulseMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1, e.results.err
		}
	}

	if mmForPulse.ForPulseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForPulse.ForPulseMock.defaultExpectation.Counter, 1)
		want := mmForPulse.ForPulseMock.defaultExpectation.params
		got := IndexAccessorMockForPulseParams{ctx, pn}
		if want != nil && !minimock.Equal(*want, got) {
			mmForPulse.t.Errorf("IndexAccessorMock.ForPulse got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmForPulse.ForPulseMock.defaultExpectation.results
		if results == nil {
			mmForPulse.t.Fatal("No results are set for the IndexAccessorMock.ForPulse")
		}
		return (*results).ia1, (*results).err
	}
	if mmForPulse.funcForPulse != nil {
		return mmForPulse.funcForPulse(ctx, pn)
	}
	mmForPulse.t.Fatalf("Unexpected call to IndexAccessorMock.ForPulse. %v %v", ctx, pn)
	return
}

// ForPulseAfterCounter returns a count of finished IndexAccessorMock.ForPulse invocations
func (mmForPulse *IndexAccessorMock) ForPulseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForPulse.afterForPulseCounter)
}

// ForPulseBeforeCounter returns a count of IndexAccessorMock.ForPulse invocations
func (mmForPulse *IndexAccessorMock) ForPulseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForPulse.beforeForPulseCounter)
}

// Calls returns a list of arguments used in each call to IndexAccessorMock.ForPulse.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForPulse *mIndexAccessorMockForPulse) Calls() []*IndexAccessorMockForPulseParams {
	mmForPulse.mutex.RLock()

	argCopy := make([]*IndexAccessorMockForPulseParams, len(mmForPulse.callArgs))
	copy(argCopy, mmForPulse.callArgs)

	mmForPulse.mutex.RUnlock()

	return argCopy
}

// MinimockForPulseDone returns true if the count of the ForPulse invocations corresponds
// the number of defined expectations
func (m *IndexAccessorMock) MinimockForPulseDone() bool {
	for _, e := range m.ForPulseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForPulseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForPulse != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		return false
	}
	return true
}

// MinimockForPulseInspect logs each unmet expectation
func (m *IndexAccessorMock) MinimockForPulseInspect() {
	for _, e := range m.ForPulseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IndexAccessorMock.ForPulse with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForPulseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		if m.ForPulseMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IndexAccessorMock.ForPulse")
		} else {
			m.t.Errorf("Expected call to IndexAccessorMock.ForPulse with params: %#v", *m.ForPulseMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForPulse != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		m.t.Error("Expected call to IndexAccessorMock.ForPulse")
	}
}

type mIndexAccessorMockLastKnownForID struct {
	mock               *IndexAccessorMock
	defaultExpectation *IndexAccessorMockLastKnownForIDExpectation
	expectations       []*IndexAccessorMockLastKnownForIDExpectation

	callArgs []*IndexAccessorMockLastKnownForIDParams
	mutex    sync.RWMutex
}

// IndexAccessorMockLastKnownForIDExpectation specifies expectation struct of the IndexAccessor.LastKnownForID
type IndexAccessorMockLastKnownForIDExpectation struct {
	mock    *IndexAccessorMock
	params  *IndexAccessorMockLastKnownForIDParams
	results *IndexAccessorMockLastKnownForIDResults
	Counter uint64
}

// IndexAccessorMockLastKnownForIDParams contains parameters of the IndexAccessor.LastKnownForID
type IndexAccessorMockLastKnownForIDParams struct {
	ctx   context.Context
	objID insolar.ID
}

// IndexAccessorMockLastKnownForIDResults contains results of the IndexAccessor.LastKnownForID
type IndexAccessorMockLastKnownForIDResults struct {
	i1  record.Index
	err error
}

// Expect sets up expected params for IndexAccessor.LastKnownForID
func (mmLastKnownForID *mIndexAccessorMockLastKnownForID) Expect(ctx context.Context, objID insolar.ID) *mIndexAccessorMockLastKnownForID {
	if mmLastKnownForID.mock.funcLastKnownForID != nil {
		mmLastKnownForID.mock.t.Fatalf("IndexAccessorMock.LastKnownForID mock is already set by Set")
	}

	if mmLastKnownForID.defaultExpectation == nil {
		mmLastKnownForID.defaultExpectation = &IndexAccessorMockLastKnownForIDExpectation{}
	}

	mmLastKnownForID.defaultExpectation.params = &IndexAccessorMockLastKnownForIDParams{ctx, objID}
	for _, e := range mmLastKnownForID.expectations {
		if minimock.Equal(e.params, mmLastKnownForID.defaultExpectation.params) {
			mmLastKnownForID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLastKnownForID.defaultExpectation.params)
		}
	}

	return mmLastKnownForID
}

// Inspect accepts an inspector function that has same arguments as the IndexAccessor.LastKnownForID
func (mmLastKnownForID *mIndexAccessorMockLastKnownForID) Inspect(f func(ctx context.Context, objID insolar.ID)) *mIndexAccessorMockLastKnownForID {
	if mmLastKnownForID.mock.inspectFuncLastKnownForID != nil {
		mmLastKnownForID.mock.t.Fatalf("Inspect function is already set for IndexAccessorMock.LastKnownForID")
	}

	mmLastKnownForID.mock.inspectFuncLastKnownForID = f

	return mmLastKnownForID
}

// Return sets up results that will be returned by IndexAccessor.LastKnownForID
func (mmLastKnownForID *mIndexAccessorMockLastKnownForID) Return(i1 record.Index, err error) *IndexAccessorMock {
	if mmLastKnownForID.mock.funcLastKnownForID != nil {
		mmLastKnownForID.mock.t.Fatalf("IndexAccessorMock.LastKnownForID mock is already set by Set")
	}

	if mmLastKnownForID.defaultExpectation == nil {
		mmLastKnownForID.defaultExpectation = &IndexAccessorMockLastKnownForIDExpectation{mock: mmLastKnownForID.mock}
	}
	mmLastKnownForID.defaultExpectation.results = &IndexAccessorMockLastKnownForIDResults{i1, err}
	return mmLastKnownForID.mock
}

//Set uses given function f to mock the IndexAccessor.LastKnownForID method
func (mmLastKnownForID *mIndexAccessorMockLastKnownForID) Set(f func(ctx context.Context, objID insolar.ID) (i1 record.Index, err error)) *IndexAccessorMock {
	if mmLastKnownForID.defaultExpectation != nil {
		mmLastKnownForID.mock.t.Fatalf("Default expectation is already set for the IndexAccessor.LastKnownForID method")
	}

	if len(mmLastKnownForID.expectations) > 0 {
		mmLastKnownForID.mock.t.Fatalf("Some expectations are already set for the IndexAccessor.LastKnownForID method")
	}

	mmLastKnownForID.mock.funcLastKnownForID = f
	return mmLastKnownForID.mock
}

// When sets expectation for the IndexAccessor.LastKnownForID which will trigger the result defined by the following
// Then helper
func (mmLastKnownForID *mIndexAccessorMockLastKnownForID) When(ctx context.Context, objID insolar.ID) *IndexAccessorMockLastKnownForIDExpectation {
	if mmLastKnownForID.mock.funcLastKnownForID != nil {
		mmLastKnownForID.mock.t.Fatalf("IndexAccessorMock.LastKnownForID mock is already set by Set")
	}

	expectation := &IndexAccessorMockLastKnownForIDExpectation{
		mock:   mmLastKnownForID.mock,
		params: &IndexAccessorMockLastKnownForIDParams{ctx, objID},
	}
	mmLastKnownForID.expectations = append(mmLastKnownForID.expectations, expectation)
	return expectation
}

// Then sets up IndexAccessor.LastKnownForID return parameters for the expectation previously defined by the When method
func (e *IndexAccessorMockLastKnownForIDExpectation) Then(i1 record.Index, err error) *IndexAccessorMock {
	e.results = &IndexAccessorMockLastKnownForIDResults{i1, err}
	return e.mock
}

// LastKnownForID implements IndexAccessor
func (mmLastKnownForID *IndexAccessorMock) LastKnownForID(ctx context.Context, objID insolar.ID) (i1 record.Index, err error) {
	mm_atomic.AddUint64(&mmLastKnownForID.beforeLastKnownForIDCounter, 1)
	defer mm_atomic.AddUint64(&mmLastKnownForID.afterLastKnownForIDCounter, 1)

	if mmLastKnownForID.inspectFuncLastKnownForID != nil {
		mmLastKnownForID.inspectFuncLastKnownForID(ctx, objID)
	}

	params := &IndexAccessorMockLastKnownForIDParams{ctx, objID}

	// Record call args
	mmLastKnownForID.LastKnownForIDMock.mutex.Lock()
	mmLastKnownForID.LastKnownForIDMock.callArgs = append(mmLastKnownForID.LastKnownForIDMock.callArgs, params)
	mmLastKnownForID.LastKnownForIDMock.mutex.Unlock()

	for _, e := range mmLastKnownForID.LastKnownForIDMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmLastKnownForID.LastKnownForIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLastKnownForID.LastKnownForIDMock.defaultExpectation.Counter, 1)
		want := mmLastKnownForID.LastKnownForIDMock.defaultExpectation.params
		got := IndexAccessorMockLastKnownForIDParams{ctx, objID}
		if want != nil && !minimock.Equal(*want, got) {
			mmLastKnownForID.t.Errorf("IndexAccessorMock.LastKnownForID got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmLastKnownForID.LastKnownForIDMock.defaultExpectation.results
		if results == nil {
			mmLastKnownForID.t.Fatal("No results are set for the IndexAccessorMock.LastKnownForID")
		}
		return (*results).i1, (*results).err
	}
	if mmLastKnownForID.funcLastKnownForID != nil {
		return mmLastKnownForID.funcLastKnownForID(ctx, objID)
	}
	mmLastKnownForID.t.Fatalf("Unexpected call to IndexAccessorMock.LastKnownForID. %v %v", ctx, objID)
	return
}

// LastKnownForIDAfterCounter returns a count of finished IndexAccessorMock.LastKnownForID invocations
func (mmLastKnownForID *IndexAccessorMock) LastKnownForIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLastKnownForID.afterLastKnownForIDCounter)
}

// LastKnownForIDBeforeCounter returns a count of IndexAccessorMock.LastKnownForID invocations
func (mmLastKnownForID *IndexAccessorMock) LastKnownForIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLastKnownForID.beforeLastKnownForIDCounter)
}

// Calls returns a list of arguments used in each call to IndexAccessorMock.LastKnownForID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLastKnownForID *mIndexAccessorMockLastKnownForID) Calls() []*IndexAccessorMockLastKnownForIDParams {
	mmLastKnownForID.mutex.RLock()

	argCopy := make([]*IndexAccessorMockLastKnownForIDParams, len(mmLastKnownForID.callArgs))
	copy(argCopy, mmLastKnownForID.callArgs)

	mmLastKnownForID.mutex.RUnlock()

	return argCopy
}

// MinimockLastKnownForIDDone returns true if the count of the LastKnownForID invocations corresponds
// the number of defined expectations
func (m *IndexAccessorMock) MinimockLastKnownForIDDone() bool {
	for _, e := range m.LastKnownForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LastKnownForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLastKnownForIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLastKnownForID != nil && mm_atomic.LoadUint64(&m.afterLastKnownForIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockLastKnownForIDInspect logs each unmet expectation
func (m *IndexAccessorMock) MinimockLastKnownForIDInspect() {
	for _, e := range m.LastKnownForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IndexAccessorMock.LastKnownForID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LastKnownForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLastKnownForIDCounter) < 1 {
		if m.LastKnownForIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IndexAccessorMock.LastKnownForID")
		} else {
			m.t.Errorf("Expected call to IndexAccessorMock.LastKnownForID with params: %#v", *m.LastKnownForIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLastKnownForID != nil && mm_atomic.LoadUint64(&m.afterLastKnownForIDCounter) < 1 {
		m.t.Error("Expected call to IndexAccessorMock.LastKnownForID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *IndexAccessorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockForIDInspect()

		m.MinimockForPulseInspect()

		m.MinimockLastKnownForIDInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *IndexAccessorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *IndexAccessorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockForIDDone() &&
		m.MinimockForPulseDone() &&
		m.MinimockLastKnownForIDDone()
}
