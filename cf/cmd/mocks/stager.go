// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cflocal/cf/cmd (interfaces: Stager)

// Package mocks is a generated GoMock package.
package mocks

import (
	forge "github.com/buildpack/forge"
	engine "github.com/buildpack/forge/engine"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStager) EXPECT() *MockStagerMockRecorder {
	return m.recorder
}

// Stage mocks base method
func (m *MockStager) Stage(arg0 *forge.StageConfig) (engine.Stream, error) {
	ret := m.ctrl.Call(m, "Stage", arg0)
	ret0, _ := ret[0].(engine.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stage indicates an expected call of Stage
func (mr *MockStagerMockRecorder) Stage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stage", reflect.TypeOf((*MockStager)(nil).Stage), arg0)
}
