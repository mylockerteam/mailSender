// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mylockerteam/mailSender (interfaces: GomailSender)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	gomail_v2 "gopkg.in/gomail.v2"
	reflect "reflect"
)

// MockGomailSender is a mock of GomailSender interface
type MockGomailSender struct {
	ctrl     *gomock.Controller
	recorder *MockGomailSenderMockRecorder
}

// MockGomailSenderMockRecorder is the mock recorder for MockGomailSender
type MockGomailSenderMockRecorder struct {
	mock *MockGomailSender
}

// NewMockGomailSender creates a new mock instance
func NewMockGomailSender(ctrl *gomock.Controller) *MockGomailSender {
	mock := &MockGomailSender{ctrl: ctrl}
	mock.recorder = &MockGomailSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGomailSender) EXPECT() *MockGomailSenderMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockGomailSender) Send(arg0 gomail_v2.Sender, arg1 ...*gomail_v2.Message) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockGomailSenderMockRecorder) Send(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockGomailSender)(nil).Send), varargs...)
}
