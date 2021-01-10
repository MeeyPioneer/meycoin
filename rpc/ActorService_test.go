// Code generated by mockery v1.0.0. DO NOT EDIT.
package rpc

import (
	actor "github.com/meeypioneer/mey-actor/actor"
	"time"
)
import (
	"github.com/meeypioneer/meycoin/types"
	mock "github.com/stretchr/testify/mock"
)

// MockActorService is an autogenerated mock type for the MockActorService type
type MockActorService struct {
	mock.Mock
}

// CallRequest provides a mock function with given fields: _a0, msg, timeout
func (_m *MockActorService) CallRequest(_a0 string, msg interface{}, timeout time.Duration) (interface{}, error) {
	ret := _m.Called(_a0, msg, timeout)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) interface{}); ok {
		r0 = rf(_a0, msg, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}, time.Duration) error); ok {
		r1 = rf(_a0, msg, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallRequestDefaultTimeout provides a mock function with given fields: _a0, msg
func (_m *MockActorService) CallRequestDefaultTimeout(_a0 string, msg interface{}) (interface{}, error) {
	ret := _m.Called(_a0, msg)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(_a0, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(_a0, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FutureRequest provides a mock function with given fields: _a0, msg, timeout
func (_m *MockActorService) FutureRequest(_a0 string, msg interface{}, timeout time.Duration) *actor.Future {
	ret := _m.Called(_a0, msg, timeout)

	var r0 *actor.Future
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) *actor.Future); ok {
		r0 = rf(_a0, msg, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.Future)
		}
	}

	return r0
}

// FutureRequestDefaultTimeout provides a mock function with given fields: _a0, msg
func (_m *MockActorService) FutureRequestDefaultTimeout(_a0 string, msg interface{}) *actor.Future {
	ret := _m.Called(_a0, msg)

	var r0 *actor.Future
	if rf, ok := ret.Get(0).(func(string, interface{}) *actor.Future); ok {
		r0 = rf(_a0, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.Future)
		}
	}

	return r0
}

// SendRequest provides a mock function with given fields: _a0, msg
func (_m *MockActorService) SendRequest(_a0 string, msg interface{}) {
	_m.Called(_a0, msg)
}

// GetChainAccessor implement interface method of ActorService
func (_m *MockActorService) GetChainAccessor() types.ChainAccessor {
	return nil
}

// TellRequest provides a mock function with given fields: _a0, msg
func (_m *MockActorService) TellRequest(_a0 string, msg interface{}) {
	_m.Called(_a0, msg)
}
