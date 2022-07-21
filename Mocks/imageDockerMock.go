// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/imageDockerPort.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/docker-generator/api/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockDockerHubRepository is a mock of DockerHubRepository interface.
type MockDockerHubRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDockerHubRepositoryMockRecorder
}

// MockDockerHubRepositoryMockRecorder is the mock recorder for MockDockerHubRepository.
type MockDockerHubRepositoryMockRecorder struct {
	mock *MockDockerHubRepository
}

// NewMockDockerHubRepository creates a new mock instance.
func NewMockDockerHubRepository(ctrl *gomock.Controller) *MockDockerHubRepository {
	mock := &MockDockerHubRepository{ctrl: ctrl}
	mock.recorder = &MockDockerHubRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDockerHubRepository) EXPECT() *MockDockerHubRepositoryMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockDockerHubRepository) Read(image, tag string) (domain.DockerImageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", image, tag)
	ret0, _ := ret[0].(domain.DockerImageResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockDockerHubRepositoryMockRecorder) Read(image, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDockerHubRepository)(nil).Read), image, tag)
}

// MockRedisRepository is a mock of RedisRepository interface.
type MockRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisRepositoryMockRecorder
}

// MockRedisRepositoryMockRecorder is the mock recorder for MockRedisRepository.
type MockRedisRepositoryMockRecorder struct {
	mock *MockRedisRepository
}

// NewMockRedisRepository creates a new mock instance.
func NewMockRedisRepository(ctrl *gomock.Controller) *MockRedisRepository {
	mock := &MockRedisRepository{ctrl: ctrl}
	mock.recorder = &MockRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisRepository) EXPECT() *MockRedisRepositoryMockRecorder {
	return m.recorder
}

// ImageExist mocks base method.
func (m *MockRedisRepository) ImageExist(image, tag string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageExist", image, tag)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ImageExist indicates an expected call of ImageExist.
func (mr *MockRedisRepositoryMockRecorder) ImageExist(image, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageExist", reflect.TypeOf((*MockRedisRepository)(nil).ImageExist), image, tag)
}

// Read mocks base method.
func (m *MockRedisRepository) Read(image, tag string) (domain.DockerImageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", image, tag)
	ret0, _ := ret[0].(domain.DockerImageResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRedisRepositoryMockRecorder) Read(image, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRedisRepository)(nil).Read), image, tag)
}

// MockImageDockerService is a mock of ImageDockerService interface.
type MockImageDockerService struct {
	ctrl     *gomock.Controller
	recorder *MockImageDockerServiceMockRecorder
}

// MockImageDockerServiceMockRecorder is the mock recorder for MockImageDockerService.
type MockImageDockerServiceMockRecorder struct {
	mock *MockImageDockerService
}

// NewMockImageDockerService creates a new mock instance.
func NewMockImageDockerService(ctrl *gomock.Controller) *MockImageDockerService {
	mock := &MockImageDockerService{ctrl: ctrl}
	mock.recorder = &MockImageDockerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageDockerService) EXPECT() *MockImageDockerServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockImageDockerService) Get(image, tag string) (domain.DockerImageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", image, tag)
	ret0, _ := ret[0].(domain.DockerImageResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockImageDockerServiceMockRecorder) Get(image, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockImageDockerService)(nil).Get), image, tag)
}
