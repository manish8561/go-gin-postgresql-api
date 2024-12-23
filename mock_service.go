// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DoSomething mocks base method.
func (m *MockService) DoSomething(param string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSomething", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSomething indicates an expected call of DoSomething.
func (mr *MockServiceMockRecorder) DoSomething(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSomething", reflect.TypeOf((*MockService)(nil).DoSomething), param)
}
