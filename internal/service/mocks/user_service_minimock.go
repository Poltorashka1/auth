// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i auth/internal/service.UserService -o user_service_minimock.go -n UserServiceMock -p mocks

import (
	serviceModel "auth/internal/service/user/model"
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// UserServiceMock implements service.UserService
type UserServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSignUp          func(ctx context.Context, user serviceModel.SignUpUser) (i1 int64, err error)
	inspectFuncSignUp   func(ctx context.Context, user serviceModel.SignUpUser)
	afterSignUpCounter  uint64
	beforeSignUpCounter uint64
	SignUpMock          mUserServiceMockSignUp
}

// NewUserServiceMock returns a mock for service.UserService
func NewUserServiceMock(t minimock.Tester) *UserServiceMock {
	m := &UserServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SignUpMock = mUserServiceMockSignUp{mock: m}
	m.SignUpMock.callArgs = []*UserServiceMockSignUpParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserServiceMockSignUp struct {
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockSignUpExpectation
	expectations       []*UserServiceMockSignUpExpectation

	callArgs []*UserServiceMockSignUpParams
	mutex    sync.RWMutex
}

// UserServiceMockSignUpExpectation specifies expectation struct of the UserService.SignUp
type UserServiceMockSignUpExpectation struct {
	mock    *UserServiceMock
	params  *UserServiceMockSignUpParams
	results *UserServiceMockSignUpResults
	Counter uint64
}

// UserServiceMockSignUpParams contains parameters of the UserService.SignUp
type UserServiceMockSignUpParams struct {
	ctx  context.Context
	user serviceModel.SignUpUser
}

// UserServiceMockSignUpResults contains results of the UserService.SignUp
type UserServiceMockSignUpResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for UserService.SignUp
func (mmSignUp *mUserServiceMockSignUp) Expect(ctx context.Context, user serviceModel.SignUpUser) *mUserServiceMockSignUp {
	if mmSignUp.mock.funcSignUp != nil {
		mmSignUp.mock.t.Fatalf("UserServiceMock.SignUp mock is already set by Set")
	}

	if mmSignUp.defaultExpectation == nil {
		mmSignUp.defaultExpectation = &UserServiceMockSignUpExpectation{}
	}

	mmSignUp.defaultExpectation.params = &UserServiceMockSignUpParams{ctx, user}
	for _, e := range mmSignUp.expectations {
		if minimock.Equal(e.params, mmSignUp.defaultExpectation.params) {
			mmSignUp.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSignUp.defaultExpectation.params)
		}
	}

	return mmSignUp
}

// Inspect accepts an inspector function that has same arguments as the UserService.SignUp
func (mmSignUp *mUserServiceMockSignUp) Inspect(f func(ctx context.Context, user serviceModel.SignUpUser)) *mUserServiceMockSignUp {
	if mmSignUp.mock.inspectFuncSignUp != nil {
		mmSignUp.mock.t.Fatalf("Inspect function is already set for UserServiceMock.SignUp")
	}

	mmSignUp.mock.inspectFuncSignUp = f

	return mmSignUp
}

// Return sets up results that will be returned by UserService.SignUp
func (mmSignUp *mUserServiceMockSignUp) Return(i1 int64, err error) *UserServiceMock {
	if mmSignUp.mock.funcSignUp != nil {
		mmSignUp.mock.t.Fatalf("UserServiceMock.SignUp mock is already set by Set")
	}

	if mmSignUp.defaultExpectation == nil {
		mmSignUp.defaultExpectation = &UserServiceMockSignUpExpectation{mock: mmSignUp.mock}
	}
	mmSignUp.defaultExpectation.results = &UserServiceMockSignUpResults{i1, err}
	return mmSignUp.mock
}

// Set uses given function f to mock the UserService.SignUp method
func (mmSignUp *mUserServiceMockSignUp) Set(f func(ctx context.Context, user serviceModel.SignUpUser) (i1 int64, err error)) *UserServiceMock {
	if mmSignUp.defaultExpectation != nil {
		mmSignUp.mock.t.Fatalf("Default expectation is already set for the UserService.SignUp method")
	}

	if len(mmSignUp.expectations) > 0 {
		mmSignUp.mock.t.Fatalf("Some expectations are already set for the UserService.SignUp method")
	}

	mmSignUp.mock.funcSignUp = f
	return mmSignUp.mock
}

// When sets expectation for the UserService.SignUp which will trigger the result defined by the following
// Then helper
func (mmSignUp *mUserServiceMockSignUp) When(ctx context.Context, user serviceModel.SignUpUser) *UserServiceMockSignUpExpectation {
	if mmSignUp.mock.funcSignUp != nil {
		mmSignUp.mock.t.Fatalf("UserServiceMock.SignUp mock is already set by Set")
	}

	expectation := &UserServiceMockSignUpExpectation{
		mock:   mmSignUp.mock,
		params: &UserServiceMockSignUpParams{ctx, user},
	}
	mmSignUp.expectations = append(mmSignUp.expectations, expectation)
	return expectation
}

// Then sets up UserService.SignUp return parameters for the expectation previously defined by the When method
func (e *UserServiceMockSignUpExpectation) Then(i1 int64, err error) *UserServiceMock {
	e.results = &UserServiceMockSignUpResults{i1, err}
	return e.mock
}

// SignUp implements service.UserService
func (mmSignUp *UserServiceMock) SignUp(ctx context.Context, user serviceModel.SignUpUser) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmSignUp.beforeSignUpCounter, 1)
	defer mm_atomic.AddUint64(&mmSignUp.afterSignUpCounter, 1)

	if mmSignUp.inspectFuncSignUp != nil {
		mmSignUp.inspectFuncSignUp(ctx, user)
	}

	mm_params := UserServiceMockSignUpParams{ctx, user}

	// Record call args
	mmSignUp.SignUpMock.mutex.Lock()
	mmSignUp.SignUpMock.callArgs = append(mmSignUp.SignUpMock.callArgs, &mm_params)
	mmSignUp.SignUpMock.mutex.Unlock()

	for _, e := range mmSignUp.SignUpMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmSignUp.SignUpMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSignUp.SignUpMock.defaultExpectation.Counter, 1)
		mm_want := mmSignUp.SignUpMock.defaultExpectation.params
		mm_got := UserServiceMockSignUpParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSignUp.t.Errorf("UserServiceMock.SignUp got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSignUp.SignUpMock.defaultExpectation.results
		if mm_results == nil {
			mmSignUp.t.Fatal("No results are set for the UserServiceMock.SignUp")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmSignUp.funcSignUp != nil {
		return mmSignUp.funcSignUp(ctx, user)
	}
	mmSignUp.t.Fatalf("Unexpected call to UserServiceMock.SignUp. %v %v", ctx, user)
	return
}

// SignUpAfterCounter returns a count of finished UserServiceMock.SignUp invocations
func (mmSignUp *UserServiceMock) SignUpAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSignUp.afterSignUpCounter)
}

// SignUpBeforeCounter returns a count of UserServiceMock.SignUp invocations
func (mmSignUp *UserServiceMock) SignUpBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSignUp.beforeSignUpCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.SignUp.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSignUp *mUserServiceMockSignUp) Calls() []*UserServiceMockSignUpParams {
	mmSignUp.mutex.RLock()

	argCopy := make([]*UserServiceMockSignUpParams, len(mmSignUp.callArgs))
	copy(argCopy, mmSignUp.callArgs)

	mmSignUp.mutex.RUnlock()

	return argCopy
}

// MinimockSignUpDone returns true if the count of the SignUp invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockSignUpDone() bool {
	for _, e := range m.SignUpMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SignUpMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSignUpCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSignUp != nil && mm_atomic.LoadUint64(&m.afterSignUpCounter) < 1 {
		return false
	}
	return true
}

// MinimockSignUpInspect logs each unmet expectation
func (m *UserServiceMock) MinimockSignUpInspect() {
	for _, e := range m.SignUpMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.SignUp with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SignUpMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSignUpCounter) < 1 {
		if m.SignUpMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserServiceMock.SignUp")
		} else {
			m.t.Errorf("Expected call to UserServiceMock.SignUp with params: %#v", *m.SignUpMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSignUp != nil && mm_atomic.LoadUint64(&m.afterSignUpCounter) < 1 {
		m.t.Error("Expected call to UserServiceMock.SignUp")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSignUpInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserServiceMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSignUpDone()
}
