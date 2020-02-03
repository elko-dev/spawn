// Code generated by MockGen. DO NOT EDIT.
// Source: applicationtype/applicationtypefactory.go

// Package applicationtype is a generated GoMock package.
package applicationtype

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPrompt is a mock of Prompt interface
type MockPrompt struct {
	ctrl     *gomock.Controller
	recorder *MockPromptMockRecorder
}

// MockPromptMockRecorder is the mock recorder for MockPrompt
type MockPromptMockRecorder struct {
	mock *MockPrompt
}

// NewMockPrompt creates a new mock instance
func NewMockPrompt(ctrl *gomock.Controller) *MockPrompt {
	mock := &MockPrompt{ctrl: ctrl}
	mock.recorder = &MockPromptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrompt) EXPECT() *MockPromptMockRecorder {
	return m.recorder
}

// ForType mocks base method
func (m *MockPrompt) ForType() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForType")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForType indicates an expected call of ForType
func (mr *MockPromptMockRecorder) ForType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForType", reflect.TypeOf((*MockPrompt)(nil).ForType))
}

// MockApplicationType is a mock of ApplicationType interface
type MockApplicationType struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationTypeMockRecorder
}

// MockApplicationTypeMockRecorder is the mock recorder for MockApplicationType
type MockApplicationTypeMockRecorder struct {
	mock *MockApplicationType
}

// NewMockApplicationType creates a new mock instance
func NewMockApplicationType(ctrl *gomock.Controller) *MockApplicationType {
	mock := &MockApplicationType{ctrl: ctrl}
	mock.recorder = &MockApplicationTypeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationType) EXPECT() *MockApplicationTypeMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockApplicationType) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockApplicationTypeMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApplicationType)(nil).Create))
}