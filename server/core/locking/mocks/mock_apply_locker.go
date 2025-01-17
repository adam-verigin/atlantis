// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/locking (interfaces: ApplyLocker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	locking "github.com/runatlantis/atlantis/server/core/locking"
	"reflect"
	"time"
)

type MockApplyLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockApplyLocker(options ...pegomock.Option) *MockApplyLocker {
	mock := &MockApplyLocker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockApplyLocker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockApplyLocker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockApplyLocker) CheckApplyLock() (locking.ApplyCommandLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockApplyLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CheckApplyLock", params, []reflect.Type{reflect.TypeOf((*locking.ApplyCommandLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 locking.ApplyCommandLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(locking.ApplyCommandLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockApplyLocker) LockApply() (locking.ApplyCommandLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockApplyLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("LockApply", params, []reflect.Type{reflect.TypeOf((*locking.ApplyCommandLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 locking.ApplyCommandLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(locking.ApplyCommandLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockApplyLocker) UnlockApply() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockApplyLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UnlockApply", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockApplyLocker) VerifyWasCalledOnce() *VerifierMockApplyLocker {
	return &VerifierMockApplyLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockApplyLocker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockApplyLocker {
	return &VerifierMockApplyLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockApplyLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockApplyLocker {
	return &VerifierMockApplyLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockApplyLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockApplyLocker {
	return &VerifierMockApplyLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockApplyLocker struct {
	mock                   *MockApplyLocker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockApplyLocker) CheckApplyLock() *MockApplyLocker_CheckApplyLock_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CheckApplyLock", params, verifier.timeout)
	return &MockApplyLocker_CheckApplyLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockApplyLocker_CheckApplyLock_OngoingVerification struct {
	mock              *MockApplyLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockApplyLocker_CheckApplyLock_OngoingVerification) GetCapturedArguments() {
}

func (c *MockApplyLocker_CheckApplyLock_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockApplyLocker) LockApply() *MockApplyLocker_LockApply_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "LockApply", params, verifier.timeout)
	return &MockApplyLocker_LockApply_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockApplyLocker_LockApply_OngoingVerification struct {
	mock              *MockApplyLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockApplyLocker_LockApply_OngoingVerification) GetCapturedArguments() {
}

func (c *MockApplyLocker_LockApply_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockApplyLocker) UnlockApply() *MockApplyLocker_UnlockApply_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UnlockApply", params, verifier.timeout)
	return &MockApplyLocker_UnlockApply_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockApplyLocker_UnlockApply_OngoingVerification struct {
	mock              *MockApplyLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockApplyLocker_UnlockApply_OngoingVerification) GetCapturedArguments() {
}

func (c *MockApplyLocker_UnlockApply_OngoingVerification) GetAllCapturedArguments() {
}
