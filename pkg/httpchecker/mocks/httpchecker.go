// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/ocm-agent/pkg/httpchecker (interfaces: HTTPChecker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHTTPChecker is a mock of HTTPChecker interface.
type MockHTTPChecker struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPCheckerMockRecorder
}

// MockHTTPCheckerMockRecorder is the mock recorder for MockHTTPChecker.
type MockHTTPCheckerMockRecorder struct {
	mock *MockHTTPChecker
}

// NewMockHTTPChecker creates a new mock instance.
func NewMockHTTPChecker(ctrl *gomock.Controller) *MockHTTPChecker {
	mock := &MockHTTPChecker{ctrl: ctrl}
	mock.recorder = &MockHTTPCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPChecker) EXPECT() *MockHTTPCheckerMockRecorder {
	return m.recorder
}

// UrlAvailabilityCheck mocks base method.
func (m *MockHTTPChecker) UrlAvailabilityCheck(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UrlAvailabilityCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UrlAvailabilityCheck indicates an expected call of UrlAvailabilityCheck.
func (mr *MockHTTPCheckerMockRecorder) UrlAvailabilityCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UrlAvailabilityCheck", reflect.TypeOf((*MockHTTPChecker)(nil).UrlAvailabilityCheck), arg0)
}
