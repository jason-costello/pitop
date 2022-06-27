// Code generated by MockGen. DO NOT EDIT.
// Source: ./disk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	interfaces "github.com/PierreKieffer/pitop/interfaces"
	gomock "github.com/golang/mock/gomock"
)

// MockDiskCollector is a mock of DiskCollector interface.
type MockDiskCollector struct {
	ctrl     *gomock.Controller
	recorder *MockDiskCollectorMockRecorder
}

// MockDiskCollectorMockRecorder is the mock recorder for MockDiskCollector.
type MockDiskCollectorMockRecorder struct {
	mock *MockDiskCollector
}

// NewMockDiskCollector creates a new mock instance.
func NewMockDiskCollector(ctrl *gomock.Controller) *MockDiskCollector {
	mock := &MockDiskCollector{ctrl: ctrl}
	mock.recorder = &MockDiskCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiskCollector) EXPECT() *MockDiskCollectorMockRecorder {
	return m.recorder
}

// ExtractDiskUsage mocks base method.
func (m *MockDiskCollector) ExtractDiskUsage() *[]interfaces.DiskInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractDiskUsage")
	ret0, _ := ret[0].(*[]interfaces.DiskInfo)
	return ret0
}

// ExtractDiskUsage indicates an expected call of ExtractDiskUsage.
func (mr *MockDiskCollectorMockRecorder) ExtractDiskUsage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractDiskUsage", reflect.TypeOf((*MockDiskCollector)(nil).ExtractDiskUsage))
}