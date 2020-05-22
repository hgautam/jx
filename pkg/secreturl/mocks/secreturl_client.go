// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/v2/pkg/secreturl (interfaces: Client)

package secreturl_test

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
)

type MockClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClient(options ...pegomock.Option) *MockClient {
	mock := &MockClient{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockClient) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockClient) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockClient) Read(_param0 string) (map[string]interface{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Read", params, []reflect.Type{reflect.TypeOf((*map[string]interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) ReadObject(_param0 string, _param1 interface{}) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ReadObject", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) ReplaceURIs(_param0 string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ReplaceURIs", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) Write(_param0 string, _param1 map[string]interface{}) (map[string]interface{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Write", params, []reflect.Type{reflect.TypeOf((*map[string]interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) WriteObject(_param0 string, _param1 interface{}) (map[string]interface{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("WriteObject", params, []reflect.Type{reflect.TypeOf((*map[string]interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) VerifyWasCalledOnce() *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClient) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockClient struct {
	mock                   *MockClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockClient) Read(_param0 string) *MockClient_Read_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Read", params, verifier.timeout)
	return &MockClient_Read_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_Read_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_Read_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClient_Read_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) ReadObject(_param0 string, _param1 interface{}) *MockClient_ReadObject_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ReadObject", params, verifier.timeout)
	return &MockClient_ReadObject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_ReadObject_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_ReadObject_OngoingVerification) GetCapturedArguments() (string, interface{}) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_ReadObject_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]interface{}, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(interface{})
		}
	}
	return
}

func (verifier *VerifierMockClient) ReplaceURIs(_param0 string) *MockClient_ReplaceURIs_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ReplaceURIs", params, verifier.timeout)
	return &MockClient_ReplaceURIs_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_ReplaceURIs_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_ReplaceURIs_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClient_ReplaceURIs_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) Write(_param0 string, _param1 map[string]interface{}) *MockClient_Write_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Write", params, verifier.timeout)
	return &MockClient_Write_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_Write_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_Write_OngoingVerification) GetCapturedArguments() (string, map[string]interface{}) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_Write_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []map[string]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]map[string]interface{}, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(map[string]interface{})
		}
	}
	return
}

func (verifier *VerifierMockClient) WriteObject(_param0 string, _param1 interface{}) *MockClient_WriteObject_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "WriteObject", params, verifier.timeout)
	return &MockClient_WriteObject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_WriteObject_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_WriteObject_OngoingVerification) GetCapturedArguments() (string, interface{}) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_WriteObject_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]interface{}, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(interface{})
		}
	}
	return
}
