// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: FetchBSF)

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFetchBSF is a mock of FetchBSF interface
type MockFetchBSF struct {
	ctrl     *gomock.Controller
	recorder *MockFetchBSFMockRecorder
}

// MockFetchBSFMockRecorder is the mock recorder for MockFetchBSF
type MockFetchBSFMockRecorder struct {
	mock *MockFetchBSF
}

// NewMockFetchBSF creates a new mock instance
func NewMockFetchBSF(ctrl *gomock.Controller) *MockFetchBSF {
	mock := &MockFetchBSF{ctrl: ctrl}
	mock.recorder = &MockFetchBSFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetchBSF) EXPECT() *MockFetchBSFMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockFetchBSF) Fetch(arg0 bool) ([][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].([][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockFetchBSFMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockFetchBSF)(nil).Fetch), arg0)
}