// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ProtonMail/proton-bridge/pkg/listener (interfaces: Listener)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockListener is a mock of Listener interface.
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
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

// Add mocks base method.
func (m *MockListener) Add(arg0 string, arg1 chan<- string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0, arg1)
}

// Add indicates an expected call of Add.
func (mr *MockListenerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockListener)(nil).Add), arg0, arg1)
}

// Book mocks base method.
func (m *MockListener) Book(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Book", arg0)
}

// Book indicates an expected call of Book.
func (mr *MockListenerMockRecorder) Book(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Book", reflect.TypeOf((*MockListener)(nil).Book), arg0)
}

// Emit mocks base method.
func (m *MockListener) Emit(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Emit", arg0, arg1)
}

// Emit indicates an expected call of Emit.
func (mr *MockListenerMockRecorder) Emit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockListener)(nil).Emit), arg0, arg1)
}

// ProvideChannel mocks base method.
func (m *MockListener) ProvideChannel(arg0 string) <-chan string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideChannel", arg0)
	ret0, _ := ret[0].(<-chan string)
	return ret0
}

// ProvideChannel indicates an expected call of ProvideChannel.
func (mr *MockListenerMockRecorder) ProvideChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideChannel", reflect.TypeOf((*MockListener)(nil).ProvideChannel), arg0)
}

// Remove mocks base method.
func (m *MockListener) Remove(arg0 string, arg1 chan<- string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0, arg1)
}

// Remove indicates an expected call of Remove.
func (mr *MockListenerMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockListener)(nil).Remove), arg0, arg1)
}

// RetryEmit mocks base method.
func (m *MockListener) RetryEmit(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RetryEmit", arg0)
}

// RetryEmit indicates an expected call of RetryEmit.
func (mr *MockListenerMockRecorder) RetryEmit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryEmit", reflect.TypeOf((*MockListener)(nil).RetryEmit), arg0)
}

// SetBuffer mocks base method.
func (m *MockListener) SetBuffer(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBuffer", arg0)
}

// SetBuffer indicates an expected call of SetBuffer.
func (mr *MockListenerMockRecorder) SetBuffer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBuffer", reflect.TypeOf((*MockListener)(nil).SetBuffer), arg0)
}

// SetLimit mocks base method.
func (m *MockListener) SetLimit(arg0 string, arg1 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLimit", arg0, arg1)
}

// SetLimit indicates an expected call of SetLimit.
func (mr *MockListenerMockRecorder) SetLimit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLimit", reflect.TypeOf((*MockListener)(nil).SetLimit), arg0, arg1)
}
