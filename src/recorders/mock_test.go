// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hr3lxphr6j/bililive-go/src/recorders (interfaces: Recorder,Manager)

// Package recorders is a generated GoMock package.
package recorders

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	live "github.com/hr3lxphr6j/bililive-go/src/live"
	reflect "reflect"
)

// MockRecorder is a mock of Recorder interface
type MockRecorder struct {
	ctrl     *gomock.Controller
	recorder *MockRecorderMockRecorder
}

// MockRecorderMockRecorder is the mock recorder for MockRecorder
type MockRecorderMockRecorder struct {
	mock *MockRecorder
}

// NewMockRecorder creates a new mock instance
func NewMockRecorder(ctrl *gomock.Controller) *MockRecorder {
	mock := &MockRecorder{ctrl: ctrl}
	mock.recorder = &MockRecorderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRecorder) EXPECT() *MockRecorderMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockRecorder) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockRecorderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRecorder)(nil).Close))
}

// Start mocks base method
func (m *MockRecorder) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockRecorderMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRecorder)(nil).Start))
}

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddRecorder mocks base method
func (m *MockManager) AddRecorder(arg0 context.Context, arg1 live.Live) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRecorder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRecorder indicates an expected call of AddRecorder
func (mr *MockManagerMockRecorder) AddRecorder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecorder", reflect.TypeOf((*MockManager)(nil).AddRecorder), arg0, arg1)
}

// Close mocks base method
func (m *MockManager) Close(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", arg0)
}

// Close indicates an expected call of Close
func (mr *MockManagerMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close), arg0)
}

// GetRecorder mocks base method
func (m *MockManager) GetRecorder(arg0 context.Context, arg1 live.ID) (Recorder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecorder", arg0, arg1)
	ret0, _ := ret[0].(Recorder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecorder indicates an expected call of GetRecorder
func (mr *MockManagerMockRecorder) GetRecorder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecorder", reflect.TypeOf((*MockManager)(nil).GetRecorder), arg0, arg1)
}

// HasRecorder mocks base method
func (m *MockManager) HasRecorder(arg0 context.Context, arg1 live.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasRecorder", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasRecorder indicates an expected call of HasRecorder
func (mr *MockManagerMockRecorder) HasRecorder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasRecorder", reflect.TypeOf((*MockManager)(nil).HasRecorder), arg0, arg1)
}

// RemoveRecorder mocks base method
func (m *MockManager) RemoveRecorder(arg0 context.Context, arg1 live.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRecorder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRecorder indicates an expected call of RemoveRecorder
func (mr *MockManagerMockRecorder) RemoveRecorder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRecorder", reflect.TypeOf((*MockManager)(nil).RemoveRecorder), arg0, arg1)
}

// RestartRecorder mocks base method
func (m *MockManager) RestartRecorder(arg0 context.Context, arg1 live.Live) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartRecorder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestartRecorder indicates an expected call of RestartRecorder
func (mr *MockManagerMockRecorder) RestartRecorder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartRecorder", reflect.TypeOf((*MockManager)(nil).RestartRecorder), arg0, arg1)
}

// Start mocks base method
func (m *MockManager) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockManagerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start), arg0)
}
