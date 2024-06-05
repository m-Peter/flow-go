// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocknetwork

import (
	network "github.com/onflow/flow-go/network"
	mock "github.com/stretchr/testify/mock"
)

// MessageValidator is an autogenerated mock type for the MessageValidator type
type MessageValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: msg
func (_m *MessageValidator) Validate(msg network.IncomingMessageScope) bool {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(network.IncomingMessageScope) bool); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewMessageValidator creates a new instance of MessageValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageValidator {
	mock := &MessageValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
