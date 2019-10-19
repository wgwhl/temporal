// The MIT License (MIT)
//
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: stateBuilder.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	shared "github.com/uber/cadence/.gen/go/shared"
	persistence "github.com/uber/cadence/common/persistence"
	reflect "reflect"
)

// MockstateBuilder is a mock of stateBuilder interface
type MockstateBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockstateBuilderMockRecorder
}

// MockstateBuilderMockRecorder is the mock recorder for MockstateBuilder
type MockstateBuilderMockRecorder struct {
	mock *MockstateBuilder
}

// NewMockstateBuilder creates a new mock instance
func NewMockstateBuilder(ctrl *gomock.Controller) *MockstateBuilder {
	mock := &MockstateBuilder{ctrl: ctrl}
	mock.recorder = &MockstateBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstateBuilder) EXPECT() *MockstateBuilderMockRecorder {
	return m.recorder
}

// applyEvents mocks base method
func (m *MockstateBuilder) applyEvents(domainID, requestID string, execution shared.WorkflowExecution, history, newRunHistory []*shared.HistoryEvent, newRunNDC bool) (*shared.HistoryEvent, *decisionInfo, mutableState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "applyEvents", domainID, requestID, execution, history, newRunHistory, newRunNDC)
	ret0, _ := ret[0].(*shared.HistoryEvent)
	ret1, _ := ret[1].(*decisionInfo)
	ret2, _ := ret[2].(mutableState)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// applyEvents indicates an expected call of applyEvents
func (mr *MockstateBuilderMockRecorder) applyEvents(domainID, requestID, execution, history, newRunHistory, newRunNDC interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "applyEvents", reflect.TypeOf((*MockstateBuilder)(nil).applyEvents), domainID, requestID, execution, history, newRunHistory, newRunNDC)
}

// getTransferTasks mocks base method
func (m *MockstateBuilder) getTransferTasks() []persistence.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTransferTasks")
	ret0, _ := ret[0].([]persistence.Task)
	return ret0
}

// getTransferTasks indicates an expected call of getTransferTasks
func (mr *MockstateBuilderMockRecorder) getTransferTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTransferTasks", reflect.TypeOf((*MockstateBuilder)(nil).getTransferTasks))
}

// getTimerTasks mocks base method
func (m *MockstateBuilder) getTimerTasks() []persistence.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTimerTasks")
	ret0, _ := ret[0].([]persistence.Task)
	return ret0
}

// getTimerTasks indicates an expected call of getTimerTasks
func (mr *MockstateBuilderMockRecorder) getTimerTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTimerTasks", reflect.TypeOf((*MockstateBuilder)(nil).getTimerTasks))
}

// getNewRunTransferTasks mocks base method
func (m *MockstateBuilder) getNewRunTransferTasks() []persistence.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNewRunTransferTasks")
	ret0, _ := ret[0].([]persistence.Task)
	return ret0
}

// getNewRunTransferTasks indicates an expected call of getNewRunTransferTasks
func (mr *MockstateBuilderMockRecorder) getNewRunTransferTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNewRunTransferTasks", reflect.TypeOf((*MockstateBuilder)(nil).getNewRunTransferTasks))
}

// getNewRunTimerTasks mocks base method
func (m *MockstateBuilder) getNewRunTimerTasks() []persistence.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNewRunTimerTasks")
	ret0, _ := ret[0].([]persistence.Task)
	return ret0
}

// getNewRunTimerTasks indicates an expected call of getNewRunTimerTasks
func (mr *MockstateBuilderMockRecorder) getNewRunTimerTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNewRunTimerTasks", reflect.TypeOf((*MockstateBuilder)(nil).getNewRunTimerTasks))
}

// getMutableState mocks base method
func (m *MockstateBuilder) getMutableState() mutableState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getMutableState")
	ret0, _ := ret[0].(mutableState)
	return ret0
}

// getMutableState indicates an expected call of getMutableState
func (mr *MockstateBuilderMockRecorder) getMutableState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getMutableState", reflect.TypeOf((*MockstateBuilder)(nil).getMutableState))
}
