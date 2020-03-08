// Code generated by MockGen. DO NOT EDIT.
// Source: firebase/firebase.go

// Package firebase is a generated GoMock package.
package firebase

import (
	applications "github.com/elko-dev/spawn/applications"
	gcp "github.com/elko-dev/spawn/gcp"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProject is a mock of Project interface
type MockProject struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMockRecorder
}

// MockProjectMockRecorder is the mock recorder for MockProject
type MockProjectMockRecorder struct {
	mock *MockProject
}

// NewMockProject creates a new mock instance
func NewMockProject(ctrl *gomock.Controller) *MockProject {
	mock := &MockProject{ctrl: ctrl}
	mock.recorder = &MockProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProject) EXPECT() *MockProjectMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockProject) Create(project gcp.ProjectRequest) (gcp.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", project)
	ret0, _ := ret[0].(gcp.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProjectMockRecorder) Create(project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProject)(nil).Create), project)
}

// MockFirebaseProject is a mock of FirebaseProject interface
type MockFirebaseProject struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseProjectMockRecorder
}

// MockFirebaseProjectMockRecorder is the mock recorder for MockFirebaseProject
type MockFirebaseProjectMockRecorder struct {
	mock *MockFirebaseProject
}

// NewMockFirebaseProject creates a new mock instance
func NewMockFirebaseProject(ctrl *gomock.Controller) *MockFirebaseProject {
	mock := &MockFirebaseProject{ctrl: ctrl}
	mock.recorder = &MockFirebaseProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirebaseProject) EXPECT() *MockFirebaseProjectMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFirebaseProject) Create(gcpProjectID string) (FirebaseProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", gcpProjectID)
	ret0, _ := ret[0].(FirebaseProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFirebaseProjectMockRecorder) Create(gcpProjectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFirebaseProject)(nil).Create), gcpProjectID)
}

// MockIosApp is a mock of IosApp interface
type MockIosApp struct {
	ctrl     *gomock.Controller
	recorder *MockIosAppMockRecorder
}

// MockIosAppMockRecorder is the mock recorder for MockIosApp
type MockIosAppMockRecorder struct {
	mock *MockIosApp
}

// NewMockIosApp creates a new mock instance
func NewMockIosApp(ctrl *gomock.Controller) *MockIosApp {
	mock := &MockIosApp{ctrl: ctrl}
	mock.recorder = &MockIosAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIosApp) EXPECT() *MockIosAppMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIosApp) Create(projectID string, request IOSRequest) (applications.IOSApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectID, request)
	ret0, _ := ret[0].(applications.IOSApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockIosAppMockRecorder) Create(projectID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIosApp)(nil).Create), projectID, request)
}

// MockAndroidApp is a mock of AndroidApp interface
type MockAndroidApp struct {
	ctrl     *gomock.Controller
	recorder *MockAndroidAppMockRecorder
}

// MockAndroidAppMockRecorder is the mock recorder for MockAndroidApp
type MockAndroidAppMockRecorder struct {
	mock *MockAndroidApp
}

// NewMockAndroidApp creates a new mock instance
func NewMockAndroidApp(ctrl *gomock.Controller) *MockAndroidApp {
	mock := &MockAndroidApp{ctrl: ctrl}
	mock.recorder = &MockAndroidAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAndroidApp) EXPECT() *MockAndroidAppMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAndroidApp) Create(projectID string, request AndroidRequest) (applications.AndroidApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectID, request)
	ret0, _ := ret[0].(applications.AndroidApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAndroidAppMockRecorder) Create(projectID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAndroidApp)(nil).Create), projectID, request)
}
