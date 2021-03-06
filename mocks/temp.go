// Code generated by MockGen. DO NOT EDIT.
// Source: ./temp.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTemperature is a mock of Temperature interface.
type MockTemperature struct {
	ctrl     *gomock.Controller
	recorder *MockTemperatureMockRecorder
}

// MockTemperatureMockRecorder is the mock recorder for MockTemperature.
type MockTemperatureMockRecorder struct {
	mock *MockTemperature
}

// NewMockTemperature creates a new mock instance.
func NewMockTemperature(ctrl *gomock.Controller) *MockTemperature {
	mock := &MockTemperature{ctrl: ctrl}
	mock.recorder = &MockTemperatureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemperature) EXPECT() *MockTemperatureMockRecorder {
	return m.recorder
}

// ExtractTemp mocks base method.
func (m *MockTemperature) ExtractTemp() *float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractTemp")
	ret0, _ := ret[0].(*float64)
	return ret0
}

// ExtractTemp indicates an expected call of ExtractTemp.
func (mr *MockTemperatureMockRecorder) ExtractTemp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractTemp", reflect.TypeOf((*MockTemperature)(nil).ExtractTemp))
}
