package hot

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// WriteManagerMock implements WriteManager
type WriteManagerMock struct {
	t minimock.Tester

	funcCloseAndWait          func(ctx context.Context, p1 insolar.PulseNumber) (err error)
	inspectFuncCloseAndWait   func(ctx context.Context, p1 insolar.PulseNumber)
	afterCloseAndWaitCounter  uint64
	beforeCloseAndWaitCounter uint64
	CloseAndWaitMock          mWriteManagerMockCloseAndWait

	funcOpen          func(ctx context.Context, p1 insolar.PulseNumber) (err error)
	inspectFuncOpen   func(ctx context.Context, p1 insolar.PulseNumber)
	afterOpenCounter  uint64
	beforeOpenCounter uint64
	OpenMock          mWriteManagerMockOpen
}

// NewWriteManagerMock returns a mock for WriteManager
func NewWriteManagerMock(t minimock.Tester) *WriteManagerMock {
	m := &WriteManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseAndWaitMock = mWriteManagerMockCloseAndWait{mock: m}
	m.CloseAndWaitMock.callArgs = []*WriteManagerMockCloseAndWaitParams{}

	m.OpenMock = mWriteManagerMockOpen{mock: m}
	m.OpenMock.callArgs = []*WriteManagerMockOpenParams{}

	return m
}

type mWriteManagerMockCloseAndWait struct {
	mock               *WriteManagerMock
	defaultExpectation *WriteManagerMockCloseAndWaitExpectation
	expectations       []*WriteManagerMockCloseAndWaitExpectation

	callArgs []*WriteManagerMockCloseAndWaitParams
	mutex    sync.RWMutex
}

// WriteManagerMockCloseAndWaitExpectation specifies expectation struct of the WriteManager.CloseAndWait
type WriteManagerMockCloseAndWaitExpectation struct {
	mock    *WriteManagerMock
	params  *WriteManagerMockCloseAndWaitParams
	results *WriteManagerMockCloseAndWaitResults
	Counter uint64
}

// WriteManagerMockCloseAndWaitParams contains parameters of the WriteManager.CloseAndWait
type WriteManagerMockCloseAndWaitParams struct {
	ctx context.Context
	p1  insolar.PulseNumber
}

// WriteManagerMockCloseAndWaitResults contains results of the WriteManager.CloseAndWait
type WriteManagerMockCloseAndWaitResults struct {
	err error
}

// Expect sets up expected params for WriteManager.CloseAndWait
func (mmCloseAndWait *mWriteManagerMockCloseAndWait) Expect(ctx context.Context, p1 insolar.PulseNumber) *mWriteManagerMockCloseAndWait {
	if mmCloseAndWait.mock.funcCloseAndWait != nil {
		mmCloseAndWait.mock.t.Fatalf("WriteManagerMock.CloseAndWait mock is already set by Set")
	}

	if mmCloseAndWait.defaultExpectation == nil {
		mmCloseAndWait.defaultExpectation = &WriteManagerMockCloseAndWaitExpectation{}
	}

	mmCloseAndWait.defaultExpectation.params = &WriteManagerMockCloseAndWaitParams{ctx, p1}
	for _, e := range mmCloseAndWait.expectations {
		if minimock.Equal(e.params, mmCloseAndWait.defaultExpectation.params) {
			mmCloseAndWait.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCloseAndWait.defaultExpectation.params)
		}
	}

	return mmCloseAndWait
}

// Inspect accepts an inspector function that has same arguments as the WriteManager.CloseAndWait
func (mmCloseAndWait *mWriteManagerMockCloseAndWait) Inspect(f func(ctx context.Context, p1 insolar.PulseNumber)) *mWriteManagerMockCloseAndWait {
	if mmCloseAndWait.mock.inspectFuncCloseAndWait != nil {
		mmCloseAndWait.mock.t.Fatalf("Inspect function is already set for WriteManagerMock.CloseAndWait")
	}

	mmCloseAndWait.mock.inspectFuncCloseAndWait = f

	return mmCloseAndWait
}

// Return sets up results that will be returned by WriteManager.CloseAndWait
func (mmCloseAndWait *mWriteManagerMockCloseAndWait) Return(err error) *WriteManagerMock {
	if mmCloseAndWait.mock.funcCloseAndWait != nil {
		mmCloseAndWait.mock.t.Fatalf("WriteManagerMock.CloseAndWait mock is already set by Set")
	}

	if mmCloseAndWait.defaultExpectation == nil {
		mmCloseAndWait.defaultExpectation = &WriteManagerMockCloseAndWaitExpectation{mock: mmCloseAndWait.mock}
	}
	mmCloseAndWait.defaultExpectation.results = &WriteManagerMockCloseAndWaitResults{err}
	return mmCloseAndWait.mock
}

//Set uses given function f to mock the WriteManager.CloseAndWait method
func (mmCloseAndWait *mWriteManagerMockCloseAndWait) Set(f func(ctx context.Context, p1 insolar.PulseNumber) (err error)) *WriteManagerMock {
	if mmCloseAndWait.defaultExpectation != nil {
		mmCloseAndWait.mock.t.Fatalf("Default expectation is already set for the WriteManager.CloseAndWait method")
	}

	if len(mmCloseAndWait.expectations) > 0 {
		mmCloseAndWait.mock.t.Fatalf("Some expectations are already set for the WriteManager.CloseAndWait method")
	}

	mmCloseAndWait.mock.funcCloseAndWait = f
	return mmCloseAndWait.mock
}

// When sets expectation for the WriteManager.CloseAndWait which will trigger the result defined by the following
// Then helper
func (mmCloseAndWait *mWriteManagerMockCloseAndWait) When(ctx context.Context, p1 insolar.PulseNumber) *WriteManagerMockCloseAndWaitExpectation {
	if mmCloseAndWait.mock.funcCloseAndWait != nil {
		mmCloseAndWait.mock.t.Fatalf("WriteManagerMock.CloseAndWait mock is already set by Set")
	}

	expectation := &WriteManagerMockCloseAndWaitExpectation{
		mock:   mmCloseAndWait.mock,
		params: &WriteManagerMockCloseAndWaitParams{ctx, p1},
	}
	mmCloseAndWait.expectations = append(mmCloseAndWait.expectations, expectation)
	return expectation
}

// Then sets up WriteManager.CloseAndWait return parameters for the expectation previously defined by the When method
func (e *WriteManagerMockCloseAndWaitExpectation) Then(err error) *WriteManagerMock {
	e.results = &WriteManagerMockCloseAndWaitResults{err}
	return e.mock
}

// CloseAndWait implements WriteManager
func (mmCloseAndWait *WriteManagerMock) CloseAndWait(ctx context.Context, p1 insolar.PulseNumber) (err error) {
	mm_atomic.AddUint64(&mmCloseAndWait.beforeCloseAndWaitCounter, 1)
	defer mm_atomic.AddUint64(&mmCloseAndWait.afterCloseAndWaitCounter, 1)

	if mmCloseAndWait.inspectFuncCloseAndWait != nil {
		mmCloseAndWait.inspectFuncCloseAndWait(ctx, p1)
	}

	params := &WriteManagerMockCloseAndWaitParams{ctx, p1}

	// Record call args
	mmCloseAndWait.CloseAndWaitMock.mutex.Lock()
	mmCloseAndWait.CloseAndWaitMock.callArgs = append(mmCloseAndWait.CloseAndWaitMock.callArgs, params)
	mmCloseAndWait.CloseAndWaitMock.mutex.Unlock()

	for _, e := range mmCloseAndWait.CloseAndWaitMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCloseAndWait.CloseAndWaitMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCloseAndWait.CloseAndWaitMock.defaultExpectation.Counter, 1)
		want := mmCloseAndWait.CloseAndWaitMock.defaultExpectation.params
		got := WriteManagerMockCloseAndWaitParams{ctx, p1}
		if want != nil && !minimock.Equal(*want, got) {
			mmCloseAndWait.t.Errorf("WriteManagerMock.CloseAndWait got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmCloseAndWait.CloseAndWaitMock.defaultExpectation.results
		if results == nil {
			mmCloseAndWait.t.Fatal("No results are set for the WriteManagerMock.CloseAndWait")
		}
		return (*results).err
	}
	if mmCloseAndWait.funcCloseAndWait != nil {
		return mmCloseAndWait.funcCloseAndWait(ctx, p1)
	}
	mmCloseAndWait.t.Fatalf("Unexpected call to WriteManagerMock.CloseAndWait. %v %v", ctx, p1)
	return
}

// CloseAndWaitAfterCounter returns a count of finished WriteManagerMock.CloseAndWait invocations
func (mmCloseAndWait *WriteManagerMock) CloseAndWaitAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCloseAndWait.afterCloseAndWaitCounter)
}

// CloseAndWaitBeforeCounter returns a count of WriteManagerMock.CloseAndWait invocations
func (mmCloseAndWait *WriteManagerMock) CloseAndWaitBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCloseAndWait.beforeCloseAndWaitCounter)
}

// Calls returns a list of arguments used in each call to WriteManagerMock.CloseAndWait.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCloseAndWait *mWriteManagerMockCloseAndWait) Calls() []*WriteManagerMockCloseAndWaitParams {
	mmCloseAndWait.mutex.RLock()

	argCopy := make([]*WriteManagerMockCloseAndWaitParams, len(mmCloseAndWait.callArgs))
	copy(argCopy, mmCloseAndWait.callArgs)

	mmCloseAndWait.mutex.RUnlock()

	return argCopy
}

// MinimockCloseAndWaitDone returns true if the count of the CloseAndWait invocations corresponds
// the number of defined expectations
func (m *WriteManagerMock) MinimockCloseAndWaitDone() bool {
	for _, e := range m.CloseAndWaitMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseAndWaitMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseAndWaitCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCloseAndWait != nil && mm_atomic.LoadUint64(&m.afterCloseAndWaitCounter) < 1 {
		return false
	}
	return true
}

// MinimockCloseAndWaitInspect logs each unmet expectation
func (m *WriteManagerMock) MinimockCloseAndWaitInspect() {
	for _, e := range m.CloseAndWaitMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to WriteManagerMock.CloseAndWait with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseAndWaitMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseAndWaitCounter) < 1 {
		if m.CloseAndWaitMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to WriteManagerMock.CloseAndWait")
		} else {
			m.t.Errorf("Expected call to WriteManagerMock.CloseAndWait with params: %#v", *m.CloseAndWaitMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCloseAndWait != nil && mm_atomic.LoadUint64(&m.afterCloseAndWaitCounter) < 1 {
		m.t.Error("Expected call to WriteManagerMock.CloseAndWait")
	}
}

type mWriteManagerMockOpen struct {
	mock               *WriteManagerMock
	defaultExpectation *WriteManagerMockOpenExpectation
	expectations       []*WriteManagerMockOpenExpectation

	callArgs []*WriteManagerMockOpenParams
	mutex    sync.RWMutex
}

// WriteManagerMockOpenExpectation specifies expectation struct of the WriteManager.Open
type WriteManagerMockOpenExpectation struct {
	mock    *WriteManagerMock
	params  *WriteManagerMockOpenParams
	results *WriteManagerMockOpenResults
	Counter uint64
}

// WriteManagerMockOpenParams contains parameters of the WriteManager.Open
type WriteManagerMockOpenParams struct {
	ctx context.Context
	p1  insolar.PulseNumber
}

// WriteManagerMockOpenResults contains results of the WriteManager.Open
type WriteManagerMockOpenResults struct {
	err error
}

// Expect sets up expected params for WriteManager.Open
func (mmOpen *mWriteManagerMockOpen) Expect(ctx context.Context, p1 insolar.PulseNumber) *mWriteManagerMockOpen {
	if mmOpen.mock.funcOpen != nil {
		mmOpen.mock.t.Fatalf("WriteManagerMock.Open mock is already set by Set")
	}

	if mmOpen.defaultExpectation == nil {
		mmOpen.defaultExpectation = &WriteManagerMockOpenExpectation{}
	}

	mmOpen.defaultExpectation.params = &WriteManagerMockOpenParams{ctx, p1}
	for _, e := range mmOpen.expectations {
		if minimock.Equal(e.params, mmOpen.defaultExpectation.params) {
			mmOpen.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmOpen.defaultExpectation.params)
		}
	}

	return mmOpen
}

// Inspect accepts an inspector function that has same arguments as the WriteManager.Open
func (mmOpen *mWriteManagerMockOpen) Inspect(f func(ctx context.Context, p1 insolar.PulseNumber)) *mWriteManagerMockOpen {
	if mmOpen.mock.inspectFuncOpen != nil {
		mmOpen.mock.t.Fatalf("Inspect function is already set for WriteManagerMock.Open")
	}

	mmOpen.mock.inspectFuncOpen = f

	return mmOpen
}

// Return sets up results that will be returned by WriteManager.Open
func (mmOpen *mWriteManagerMockOpen) Return(err error) *WriteManagerMock {
	if mmOpen.mock.funcOpen != nil {
		mmOpen.mock.t.Fatalf("WriteManagerMock.Open mock is already set by Set")
	}

	if mmOpen.defaultExpectation == nil {
		mmOpen.defaultExpectation = &WriteManagerMockOpenExpectation{mock: mmOpen.mock}
	}
	mmOpen.defaultExpectation.results = &WriteManagerMockOpenResults{err}
	return mmOpen.mock
}

//Set uses given function f to mock the WriteManager.Open method
func (mmOpen *mWriteManagerMockOpen) Set(f func(ctx context.Context, p1 insolar.PulseNumber) (err error)) *WriteManagerMock {
	if mmOpen.defaultExpectation != nil {
		mmOpen.mock.t.Fatalf("Default expectation is already set for the WriteManager.Open method")
	}

	if len(mmOpen.expectations) > 0 {
		mmOpen.mock.t.Fatalf("Some expectations are already set for the WriteManager.Open method")
	}

	mmOpen.mock.funcOpen = f
	return mmOpen.mock
}

// When sets expectation for the WriteManager.Open which will trigger the result defined by the following
// Then helper
func (mmOpen *mWriteManagerMockOpen) When(ctx context.Context, p1 insolar.PulseNumber) *WriteManagerMockOpenExpectation {
	if mmOpen.mock.funcOpen != nil {
		mmOpen.mock.t.Fatalf("WriteManagerMock.Open mock is already set by Set")
	}

	expectation := &WriteManagerMockOpenExpectation{
		mock:   mmOpen.mock,
		params: &WriteManagerMockOpenParams{ctx, p1},
	}
	mmOpen.expectations = append(mmOpen.expectations, expectation)
	return expectation
}

// Then sets up WriteManager.Open return parameters for the expectation previously defined by the When method
func (e *WriteManagerMockOpenExpectation) Then(err error) *WriteManagerMock {
	e.results = &WriteManagerMockOpenResults{err}
	return e.mock
}

// Open implements WriteManager
func (mmOpen *WriteManagerMock) Open(ctx context.Context, p1 insolar.PulseNumber) (err error) {
	mm_atomic.AddUint64(&mmOpen.beforeOpenCounter, 1)
	defer mm_atomic.AddUint64(&mmOpen.afterOpenCounter, 1)

	if mmOpen.inspectFuncOpen != nil {
		mmOpen.inspectFuncOpen(ctx, p1)
	}

	params := &WriteManagerMockOpenParams{ctx, p1}

	// Record call args
	mmOpen.OpenMock.mutex.Lock()
	mmOpen.OpenMock.callArgs = append(mmOpen.OpenMock.callArgs, params)
	mmOpen.OpenMock.mutex.Unlock()

	for _, e := range mmOpen.OpenMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmOpen.OpenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmOpen.OpenMock.defaultExpectation.Counter, 1)
		want := mmOpen.OpenMock.defaultExpectation.params
		got := WriteManagerMockOpenParams{ctx, p1}
		if want != nil && !minimock.Equal(*want, got) {
			mmOpen.t.Errorf("WriteManagerMock.Open got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmOpen.OpenMock.defaultExpectation.results
		if results == nil {
			mmOpen.t.Fatal("No results are set for the WriteManagerMock.Open")
		}
		return (*results).err
	}
	if mmOpen.funcOpen != nil {
		return mmOpen.funcOpen(ctx, p1)
	}
	mmOpen.t.Fatalf("Unexpected call to WriteManagerMock.Open. %v %v", ctx, p1)
	return
}

// OpenAfterCounter returns a count of finished WriteManagerMock.Open invocations
func (mmOpen *WriteManagerMock) OpenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmOpen.afterOpenCounter)
}

// OpenBeforeCounter returns a count of WriteManagerMock.Open invocations
func (mmOpen *WriteManagerMock) OpenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmOpen.beforeOpenCounter)
}

// Calls returns a list of arguments used in each call to WriteManagerMock.Open.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmOpen *mWriteManagerMockOpen) Calls() []*WriteManagerMockOpenParams {
	mmOpen.mutex.RLock()

	argCopy := make([]*WriteManagerMockOpenParams, len(mmOpen.callArgs))
	copy(argCopy, mmOpen.callArgs)

	mmOpen.mutex.RUnlock()

	return argCopy
}

// MinimockOpenDone returns true if the count of the Open invocations corresponds
// the number of defined expectations
func (m *WriteManagerMock) MinimockOpenDone() bool {
	for _, e := range m.OpenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.OpenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterOpenCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcOpen != nil && mm_atomic.LoadUint64(&m.afterOpenCounter) < 1 {
		return false
	}
	return true
}

// MinimockOpenInspect logs each unmet expectation
func (m *WriteManagerMock) MinimockOpenInspect() {
	for _, e := range m.OpenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to WriteManagerMock.Open with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.OpenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterOpenCounter) < 1 {
		if m.OpenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to WriteManagerMock.Open")
		} else {
			m.t.Errorf("Expected call to WriteManagerMock.Open with params: %#v", *m.OpenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcOpen != nil && mm_atomic.LoadUint64(&m.afterOpenCounter) < 1 {
		m.t.Error("Expected call to WriteManagerMock.Open")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *WriteManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCloseAndWaitInspect()

		m.MinimockOpenInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *WriteManagerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *WriteManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseAndWaitDone() &&
		m.MinimockOpenDone()
}
