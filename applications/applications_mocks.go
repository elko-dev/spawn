// Code generated by MockGen. DO NOT EDIT.
// Source: applications/applications.go

// Package applications is a generated GoMock package.
package applications

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFactory) Create() (Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFactoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFactory)(nil).Create))
}

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
func (m *MockProject) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockProjectMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProject)(nil).Create))
}

// MockGitRepo is a mock of GitRepo interface
type MockGitRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGitRepoMockRecorder
}

// MockGitRepoMockRecorder is the mock recorder for MockGitRepo
type MockGitRepoMockRecorder struct {
	mock *MockGitRepo
}

// NewMockGitRepo creates a new mock instance
func NewMockGitRepo(ctrl *gomock.Controller) *MockGitRepo {
	mock := &MockGitRepo{ctrl: ctrl}
	mock.recorder = &MockGitRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitRepo) EXPECT() *MockGitRepoMockRecorder {
	return m.recorder
}

// CreateGitRepository mocks base method
func (m *MockGitRepo) CreateGitRepository(repositoryName, templateURL, platformToken string) (GitResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGitRepository", repositoryName, templateURL, platformToken)
	ret0, _ := ret[0].(GitResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGitRepository indicates an expected call of CreateGitRepository
func (mr *MockGitRepoMockRecorder) CreateGitRepository(repositoryName, templateURL, platformToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGitRepository", reflect.TypeOf((*MockGitRepo)(nil).CreateGitRepository), repositoryName, templateURL, platformToken)
}

// MockPlatformRepository is a mock of PlatformRepository interface
type MockPlatformRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPlatformRepositoryMockRecorder
}

// MockPlatformRepositoryMockRecorder is the mock recorder for MockPlatformRepository
type MockPlatformRepositoryMockRecorder struct {
	mock *MockPlatformRepository
}

// NewMockPlatformRepository creates a new mock instance
func NewMockPlatformRepository(ctrl *gomock.Controller) *MockPlatformRepository {
	mock := &MockPlatformRepository{ctrl: ctrl}
	mock.recorder = &MockPlatformRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlatformRepository) EXPECT() *MockPlatformRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPlatformRepository) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockPlatformRepositoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPlatformRepository)(nil).Create))
}

// GetToken mocks base method
func (m *MockPlatformRepository) GetToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetToken indicates an expected call of GetToken
func (mr *MockPlatformRepositoryMockRecorder) GetToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockPlatformRepository)(nil).GetToken))
}

// GetPlatformType mocks base method
func (m *MockPlatformRepository) GetPlatformType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatformType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPlatformType indicates an expected call of GetPlatformType
func (mr *MockPlatformRepositoryMockRecorder) GetPlatformType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatformType", reflect.TypeOf((*MockPlatformRepository)(nil).GetPlatformType))
}

// MockPlatformFactory is a mock of PlatformFactory interface
type MockPlatformFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPlatformFactoryMockRecorder
}

// MockPlatformFactoryMockRecorder is the mock recorder for MockPlatformFactory
type MockPlatformFactoryMockRecorder struct {
	mock *MockPlatformFactory
}

// NewMockPlatformFactory creates a new mock instance
func NewMockPlatformFactory(ctrl *gomock.Controller) *MockPlatformFactory {
	mock := &MockPlatformFactory{ctrl: ctrl}
	mock.recorder = &MockPlatformFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlatformFactory) EXPECT() *MockPlatformFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPlatformFactory) Create(projectName, applicationType string) (PlatformRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectName, applicationType)
	ret0, _ := ret[0].(PlatformRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPlatformFactoryMockRecorder) Create(projectName, applicationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPlatformFactory)(nil).Create), projectName, applicationType)
}

// MockGitFactory is a mock of GitFactory interface
type MockGitFactory struct {
	ctrl     *gomock.Controller
	recorder *MockGitFactoryMockRecorder
}

// MockGitFactoryMockRecorder is the mock recorder for MockGitFactory
type MockGitFactoryMockRecorder struct {
	mock *MockGitFactory
}

// NewMockGitFactory creates a new mock instance
func NewMockGitFactory(ctrl *gomock.Controller) *MockGitFactory {
	mock := &MockGitFactory{ctrl: ctrl}
	mock.recorder = &MockGitFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitFactory) EXPECT() *MockGitFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGitFactory) Create(projectName string) (GitRepo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectName)
	ret0, _ := ret[0].(GitRepo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockGitFactoryMockRecorder) Create(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGitFactory)(nil).Create), projectName)
}

// MockCIFactory is a mock of CIFactory interface
type MockCIFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCIFactoryMockRecorder
}

// MockCIFactoryMockRecorder is the mock recorder for MockCIFactory
type MockCIFactoryMockRecorder struct {
	mock *MockCIFactory
}

// NewMockCIFactory creates a new mock instance
func NewMockCIFactory(ctrl *gomock.Controller) *MockCIFactory {
	mock := &MockCIFactory{ctrl: ctrl}
	mock.recorder = &MockCIFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCIFactory) EXPECT() *MockCIFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCIFactory) Create(projectName string) (CIPlatform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectName)
	ret0, _ := ret[0].(CIPlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCIFactoryMockRecorder) Create(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCIFactory)(nil).Create), projectName)
}

// MockCIPlatform is a mock of CIPlatform interface
type MockCIPlatform struct {
	ctrl     *gomock.Controller
	recorder *MockCIPlatformMockRecorder
}

// MockCIPlatformMockRecorder is the mock recorder for MockCIPlatform
type MockCIPlatformMockRecorder struct {
	mock *MockCIPlatform
}

// NewMockCIPlatform creates a new mock instance
func NewMockCIPlatform(ctrl *gomock.Controller) *MockCIPlatform {
	mock := &MockCIPlatform{ctrl: ctrl}
	mock.recorder = &MockCIPlatformMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCIPlatform) EXPECT() *MockCIPlatformMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCIPlatform) Create(repoURL, latestGitConfig string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", repoURL, latestGitConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCIPlatformMockRecorder) Create(repoURL, latestGitConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCIPlatform)(nil).Create), repoURL, latestGitConfig)
}
