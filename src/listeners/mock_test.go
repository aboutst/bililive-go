// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bililive-go/bililive-go/src/listeners (interfaces: Listener,Manager)
//
// Generated by this command:
//
//	mockgen -package listeners -destination mock_test.go github.com/bililive-go/bililive-go/src/listeners Listener,Manager
//

// Package listeners is a generated GoMock package.
package listeners

import (
	context "context"
	reflect "reflect"

	live "github.com/bililive-go/bililive-go/src/live"
	types "github.com/bililive-go/bililive-go/src/types"
	gomock "go.uber.org/mock/gomock"
)

// MockListener is a mock of Listener interface.
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
	isgomock struct{}
}

// MockListenerMockRecorder is the mock recorder for MockListener.
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance.
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockListener) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockListenerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockListener)(nil).Close))
}

// Start mocks base method.
func (m *MockListener) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockListenerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockListener)(nil).Start))
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
	isgomock struct{}
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddListener mocks base method.
func (m *MockManager) AddListener(ctx context.Context, arg1 live.Live) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListener", ctx, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddListener indicates an expected call of AddListener.
func (mr *MockManagerMockRecorder) AddListener(ctx, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListener", reflect.TypeOf((*MockManager)(nil).AddListener), ctx, arg1)
}

// Close mocks base method.
func (m *MockManager) Close(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", ctx)
}

// Close indicates an expected call of Close.
func (mr *MockManagerMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close), ctx)
}

// GetListener mocks base method.
func (m *MockManager) GetListener(ctx context.Context, liveId types.LiveID) (Listener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListener", ctx, liveId)
	ret0, _ := ret[0].(Listener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListener indicates an expected call of GetListener.
func (mr *MockManagerMockRecorder) GetListener(ctx, liveId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListener", reflect.TypeOf((*MockManager)(nil).GetListener), ctx, liveId)
}

// HasListener mocks base method.
func (m *MockManager) HasListener(ctx context.Context, liveId types.LiveID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasListener", ctx, liveId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasListener indicates an expected call of HasListener.
func (mr *MockManagerMockRecorder) HasListener(ctx, liveId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasListener", reflect.TypeOf((*MockManager)(nil).HasListener), ctx, liveId)
}

// RemoveListener mocks base method.
func (m *MockManager) RemoveListener(ctx context.Context, liveId types.LiveID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveListener", ctx, liveId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveListener indicates an expected call of RemoveListener.
func (mr *MockManagerMockRecorder) RemoveListener(ctx, liveId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveListener", reflect.TypeOf((*MockManager)(nil).RemoveListener), ctx, liveId)
}

// Start mocks base method.
func (m *MockManager) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start), ctx)
}
