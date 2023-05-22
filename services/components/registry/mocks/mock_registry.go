// Code generated by MockGen. DO NOT EDIT.
// Source: services/components/registry/registry.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	reflect "reflect"

	health "github.com/forta-network/forta-core-go/clients/health"
	config "github.com/forta-network/forta-node/config"
	gomock "github.com/golang/mock/gomock"
)

// MockBotRegistry is a mock of BotRegistry interface.
type MockBotRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockBotRegistryMockRecorder
}

// MockBotRegistryMockRecorder is the mock recorder for MockBotRegistry.
type MockBotRegistryMockRecorder struct {
	mock *MockBotRegistry
}

// NewMockBotRegistry creates a new mock instance.
func NewMockBotRegistry(ctrl *gomock.Controller) *MockBotRegistry {
	mock := &MockBotRegistry{ctrl: ctrl}
	mock.recorder = &MockBotRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBotRegistry) EXPECT() *MockBotRegistryMockRecorder {
	return m.recorder
}

// Health mocks base method.
func (m *MockBotRegistry) Health() health.Reports {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(health.Reports)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockBotRegistryMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockBotRegistry)(nil).Health))
}

// LoadAssignedBots mocks base method.
func (m *MockBotRegistry) LoadAssignedBots() ([]config.AgentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAssignedBots")
	ret0, _ := ret[0].([]config.AgentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAssignedBots indicates an expected call of LoadAssignedBots.
func (mr *MockBotRegistryMockRecorder) LoadAssignedBots() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAssignedBots", reflect.TypeOf((*MockBotRegistry)(nil).LoadAssignedBots))
}

// Name mocks base method.
func (m *MockBotRegistry) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockBotRegistryMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBotRegistry)(nil).Name))
}
