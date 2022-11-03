// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockinsecure

import (
	context "context"

	insecure "github.com/onflow/flow-go/insecure"
	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// CorruptNetwork_ConnectAttackerServer is an autogenerated mock type for the CorruptNetwork_ConnectAttackerServer type
type CorruptNetwork_ConnectAttackerServer struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *CorruptNetwork_ConnectAttackerServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// RecvMsg provides a mock function with given fields: m
func (_m *CorruptNetwork_ConnectAttackerServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *CorruptNetwork_ConnectAttackerServer) Send(_a0 *insecure.Message) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*insecure.Message) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendHeader provides a mock function with given fields: _a0
func (_m *CorruptNetwork_ConnectAttackerServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *CorruptNetwork_ConnectAttackerServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetHeader provides a mock function with given fields: _a0
func (_m *CorruptNetwork_ConnectAttackerServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *CorruptNetwork_ConnectAttackerServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewCorruptNetwork_ConnectAttackerServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewCorruptNetwork_ConnectAttackerServer creates a new instance of CorruptNetwork_ConnectAttackerServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCorruptNetwork_ConnectAttackerServer(t mockConstructorTestingTNewCorruptNetwork_ConnectAttackerServer) *CorruptNetwork_ConnectAttackerServer {
	mock := &CorruptNetwork_ConnectAttackerServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}