// Automatically generated by pegomock. DO NOT EDIT!
// Source: second (interfaces: Person)

package second

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockPerson struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPerson() *MockPerson {
	return &MockPerson{fail: pegomock.GlobalFailHandler}
}

func (mock *MockPerson) GetName() string {
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockPerson) IsOlderThan(_param0 int) bool {
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("IsOlderThan", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockPerson) VerifyWasCalledOnce() *VerifierPerson {
	return &VerifierPerson{mock, pegomock.Times(1), nil}
}

func (mock *MockPerson) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierPerson {
	return &VerifierPerson{mock, invocationCountMatcher, nil}
}

func (mock *MockPerson) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierPerson {
	return &VerifierPerson{mock, invocationCountMatcher, inOrderContext}
}

type VerifierPerson struct {
	mock                   *MockPerson
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierPerson) GetName() *Person_GetName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetName", params)
	return &Person_GetName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Person_GetName_OngoingVerification struct {
	mock              *MockPerson
	methodInvocations []pegomock.MethodInvocation
}

func (c *Person_GetName_OngoingVerification) GetCapturedArguments() {
}

func (c *Person_GetName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierPerson) IsOlderThan(_param0 int) *Person_IsOlderThan_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "IsOlderThan", params)
	return &Person_IsOlderThan_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Person_IsOlderThan_OngoingVerification struct {
	mock              *MockPerson
	methodInvocations []pegomock.MethodInvocation
}

func (c *Person_IsOlderThan_OngoingVerification) GetCapturedArguments() int {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Person_IsOlderThan_OngoingVerification) GetAllCapturedArguments() (_param0 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
	}
	return
}