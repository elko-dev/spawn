// Code generated by MockGen. DO NOT EDIT.
// Source: web/factory.go

// Package web is a generated GoMock package.
package web

import (
	applications "github.com/elko-dev/spawn/applications"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppFactory is a mock of AppFactory interface
type MockAppFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAppFactoryMockRecorder
}

// MockAppFactoryMockRecorder is the mock recorder for MockAppFactory
type MockAppFactoryMockRecorder struct {
	mock *MockAppFactory
}

// NewMockAppFactory creates a new mock instance
func NewMockAppFactory(ctrl *gomock.Controller) *MockAppFactory {
	mock := &MockAppFactory{ctrl: ctrl}
	mock.recorder = &MockAppFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppFactory) EXPECT() *MockAppFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAppFactory) Create(applicationType string) (applications.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", applicationType)
	ret0, _ := ret[0].(applications.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAppFactoryMockRecorder) Create(applicationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppFactory)(nil).Create), applicationType)
}

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

// ForClientType mocks base method
func (m *MockPrompt) ForClientType(applicationType string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForClientType", applicationType)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForClientType indicates an expected call of ForClientType
func (mr *MockPromptMockRecorder) ForClientType(applicationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForClientType", reflect.TypeOf((*MockPrompt)(nil).ForClientType), applicationType)
}

// ForServerType mocks base method
func (m *MockPrompt) ForServerType() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForServerType")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForServerType indicates an expected call of ForServerType
func (mr *MockPromptMockRecorder) ForServerType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForServerType", reflect.TypeOf((*MockPrompt)(nil).ForServerType))
}
