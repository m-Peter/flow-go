// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// RandomGenerator is an autogenerated mock type for the RandomGenerator type
type RandomGenerator struct {
	mock.Mock
}

// UnsafeRandom provides a mock function with given fields:
func (_m *RandomGenerator) UnsafeRandom() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRandomGenerator interface {
	mock.TestingT
	Cleanup(func())
}

// NewRandomGenerator creates a new instance of RandomGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRandomGenerator(t mockConstructorTestingTNewRandomGenerator) *RandomGenerator {
	mock := &RandomGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
