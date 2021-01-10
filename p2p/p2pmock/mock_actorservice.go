// Code generated by MockGen. DO NOT EDIT.
// Source: actorservice.go

// Package p2pmock is a generated GoMock package.
package p2pmock

import (
	actor "github.com/meeypioneer/mey-actor/actor"
	types "github.com/meeypioneer/meycoin/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockActorService is a mock of ActorService interface
type MockActorService struct {
	ctrl     *gomock.Controller
	recorder *MockActorServiceMockRecorder
}

// MockActorServiceMockRecorder is the mock recorder for MockActorService
type MockActorServiceMockRecorder struct {
	mock *MockActorService
}

// NewMockActorService creates a new mock instance
func NewMockActorService(ctrl *gomock.Controller) *MockActorService {
	mock := &MockActorService{ctrl: ctrl}
	mock.recorder = &MockActorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActorService) EXPECT() *MockActorServiceMockRecorder {
	return m.recorder
}

// TellRequest mocks base method
func (m *MockActorService) TellRequest(actorName string, msg interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TellRequest", actorName, msg)
}

// TellRequest indicates an expected call of TellRequest
func (mr *MockActorServiceMockRecorder) TellRequest(actorName, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TellRequest", reflect.TypeOf((*MockActorService)(nil).TellRequest), actorName, msg)
}

// SendRequest mocks base method
func (m *MockActorService) SendRequest(actorName string, msg interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendRequest", actorName, msg)
}

// SendRequest indicates an expected call of SendRequest
func (mr *MockActorServiceMockRecorder) SendRequest(actorName, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRequest", reflect.TypeOf((*MockActorService)(nil).SendRequest), actorName, msg)
}

// CallRequest mocks base method
func (m *MockActorService) CallRequest(actorName string, msg interface{}, timeout time.Duration) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRequest", actorName, msg, timeout)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallRequest indicates an expected call of CallRequest
func (mr *MockActorServiceMockRecorder) CallRequest(actorName, msg, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRequest", reflect.TypeOf((*MockActorService)(nil).CallRequest), actorName, msg, timeout)
}

// CallRequestDefaultTimeout mocks base method
func (m *MockActorService) CallRequestDefaultTimeout(actorName string, msg interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRequestDefaultTimeout", actorName, msg)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallRequestDefaultTimeout indicates an expected call of CallRequestDefaultTimeout
func (mr *MockActorServiceMockRecorder) CallRequestDefaultTimeout(actorName, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRequestDefaultTimeout", reflect.TypeOf((*MockActorService)(nil).CallRequestDefaultTimeout), actorName, msg)
}

// FutureRequest mocks base method
func (m *MockActorService) FutureRequest(actorName string, msg interface{}, timeout time.Duration) *actor.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FutureRequest", actorName, msg, timeout)
	ret0, _ := ret[0].(*actor.Future)
	return ret0
}

// FutureRequest indicates an expected call of FutureRequest
func (mr *MockActorServiceMockRecorder) FutureRequest(actorName, msg, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FutureRequest", reflect.TypeOf((*MockActorService)(nil).FutureRequest), actorName, msg, timeout)
}

// FutureRequestDefaultTimeout mocks base method
func (m *MockActorService) FutureRequestDefaultTimeout(actorName string, msg interface{}) *actor.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FutureRequestDefaultTimeout", actorName, msg)
	ret0, _ := ret[0].(*actor.Future)
	return ret0
}

// FutureRequestDefaultTimeout indicates an expected call of FutureRequestDefaultTimeout
func (mr *MockActorServiceMockRecorder) FutureRequestDefaultTimeout(actorName, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FutureRequestDefaultTimeout", reflect.TypeOf((*MockActorService)(nil).FutureRequestDefaultTimeout), actorName, msg)
}

// GetChainAccessor mocks base method
func (m *MockActorService) GetChainAccessor() types.ChainAccessor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainAccessor")
	ret0, _ := ret[0].(types.ChainAccessor)
	return ret0
}

// GetChainAccessor indicates an expected call of GetChainAccessor
func (mr *MockActorServiceMockRecorder) GetChainAccessor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainAccessor", reflect.TypeOf((*MockActorService)(nil).GetChainAccessor))
}
