package first

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPerson is a mock of Person interface
type MockPerson struct {
	ctrl     *gomock.Controller
	recorder *MockPersonMockRecorder
}

// MockPersonMockRecorder is the mock recorder for MockPerson
type MockPersonMockRecorder struct {
	mock *MockPerson
}

// NewMockPerson creates a new mock instance
func NewMockPerson(ctrl *gomock.Controller) *MockPerson {
	mock := &MockPerson{ctrl: ctrl}
	mock.recorder = &MockPersonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPerson) EXPECT() *MockPersonMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockPerson) GetName() string {
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockPersonMockRecorder) GetName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPerson)(nil).GetName))
}

// IsOlderThan mocks base method
func (m *MockPerson) IsOlderThan(arg0 int) bool {
	ret := m.ctrl.Call(m, "IsOlderThan", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOlderThan indicates an expected call of IsOlderThan.
func (mr *MockPersonMockRecorder) IsOlderThan(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOlderThan", reflect.TypeOf((*MockPerson)(nil).IsOlderThan), arg0)
}