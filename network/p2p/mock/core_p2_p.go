// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	host "github.com/libp2p/go-libp2p/core/host"
	component "github.com/onflow/flow-go/module/component"

	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"

	mock "github.com/stretchr/testify/mock"
)

// CoreP2P is an autogenerated mock type for the CoreP2P type
type CoreP2P struct {
	mock.Mock
}

// GetIPPort provides a mock function with given fields:
func (_m *CoreP2P) GetIPPort() (string, string, error) {
	ret := _m.Called()

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func() (string, string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Host provides a mock function with given fields:
func (_m *CoreP2P) Host() host.Host {
	ret := _m.Called()

	var r0 host.Host
	if rf, ok := ret.Get(0).(func() host.Host); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(host.Host)
		}
	}

	return r0
}

// SetComponentManager provides a mock function with given fields: cm
func (_m *CoreP2P) SetComponentManager(cm *component.ComponentManager) {
	_m.Called(cm)
}

// Start provides a mock function with given fields: ctx
func (_m *CoreP2P) Start(ctx irrecoverable.SignalerContext) {
	_m.Called(ctx)
}

// Stop provides a mock function with given fields:
func (_m *CoreP2P) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCoreP2P interface {
	mock.TestingT
	Cleanup(func())
}

// NewCoreP2P creates a new instance of CoreP2P. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCoreP2P(t mockConstructorTestingTNewCoreP2P) *CoreP2P {
	mock := &CoreP2P{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}