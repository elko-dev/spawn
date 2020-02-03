// Code generated by MockGen. DO NOT EDIT.
// Source: git/factory.go

// Package git is a generated GoMock package.
package git

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

// forGitRepository mocks base method
func (m *MockPrompt) forGitRepository() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "forGitRepository")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// forGitRepository indicates an expected call of forGitRepository
func (mr *MockPromptMockRecorder) forGitRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "forGitRepository", reflect.TypeOf((*MockPrompt)(nil).forGitRepository))
}
