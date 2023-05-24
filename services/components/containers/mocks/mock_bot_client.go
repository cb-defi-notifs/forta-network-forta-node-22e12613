// Code generated by MockGen. DO NOT EDIT.
// Source: services/components/containers/bot_client.go

// Package mock_containers is a generated GoMock package.
package mock_containers

import (
	context "context"
	reflect "reflect"

	types "github.com/docker/docker/api/types"
	config "github.com/forta-network/forta-node/config"
	gomock "github.com/golang/mock/gomock"
)

// MockBotClient is a mock of BotClient interface.
type MockBotClient struct {
	ctrl     *gomock.Controller
	recorder *MockBotClientMockRecorder
}

// MockBotClientMockRecorder is the mock recorder for MockBotClient.
type MockBotClientMockRecorder struct {
	mock *MockBotClient
}

// NewMockBotClient creates a new mock instance.
func NewMockBotClient(ctrl *gomock.Controller) *MockBotClient {
	mock := &MockBotClient{ctrl: ctrl}
	mock.recorder = &MockBotClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBotClient) EXPECT() *MockBotClientMockRecorder {
	return m.recorder
}

// EnsureBotImages mocks base method.
func (m *MockBotClient) EnsureBotImages(ctx context.Context, botConfigs []config.AgentConfig) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBotImages", ctx, botConfigs)
	ret0, _ := ret[0].([]error)
	return ret0
}

// EnsureBotImages indicates an expected call of EnsureBotImages.
func (mr *MockBotClientMockRecorder) EnsureBotImages(ctx, botConfigs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBotImages", reflect.TypeOf((*MockBotClient)(nil).EnsureBotImages), ctx, botConfigs)
}

// LaunchBot mocks base method.
func (m *MockBotClient) LaunchBot(ctx context.Context, botConfig config.AgentConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchBot", ctx, botConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// LaunchBot indicates an expected call of LaunchBot.
func (mr *MockBotClientMockRecorder) LaunchBot(ctx, botConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchBot", reflect.TypeOf((*MockBotClient)(nil).LaunchBot), ctx, botConfig)
}

// LoadBotContainers mocks base method.
func (m *MockBotClient) LoadBotContainers(ctx context.Context) ([]types.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBotContainers", ctx)
	ret0, _ := ret[0].([]types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadBotContainers indicates an expected call of LoadBotContainers.
func (mr *MockBotClientMockRecorder) LoadBotContainers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBotContainers", reflect.TypeOf((*MockBotClient)(nil).LoadBotContainers), ctx)
}

// StartWaitBotContainer mocks base method.
func (m *MockBotClient) StartWaitBotContainer(ctx context.Context, containerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWaitBotContainer", ctx, containerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWaitBotContainer indicates an expected call of StartWaitBotContainer.
func (mr *MockBotClientMockRecorder) StartWaitBotContainer(ctx, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWaitBotContainer", reflect.TypeOf((*MockBotClient)(nil).StartWaitBotContainer), ctx, containerID)
}

// StopBot mocks base method.
func (m *MockBotClient) StopBot(ctx context.Context, botConfig config.AgentConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopBot", ctx, botConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopBot indicates an expected call of StopBot.
func (mr *MockBotClientMockRecorder) StopBot(ctx, botConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopBot", reflect.TypeOf((*MockBotClient)(nil).StopBot), ctx, botConfig)
}

// TearDownBot mocks base method.
func (m *MockBotClient) TearDownBot(ctx context.Context, botConfig config.AgentConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TearDownBot", ctx, botConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// TearDownBot indicates an expected call of TearDownBot.
func (mr *MockBotClientMockRecorder) TearDownBot(ctx, botConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TearDownBot", reflect.TypeOf((*MockBotClient)(nil).TearDownBot), ctx, botConfig)
}
