// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	model "github.com/onflow/flow-go/consensus/hotstuff/model"
	mock "github.com/stretchr/testify/mock"
)

// TimeoutAggregationViolationConsumer is an autogenerated mock type for the TimeoutAggregationViolationConsumer type
type TimeoutAggregationViolationConsumer struct {
	mock.Mock
}

// OnDoubleTimeoutDetected provides a mock function with given fields: _a0, _a1
func (_m *TimeoutAggregationViolationConsumer) OnDoubleTimeoutDetected(_a0 *model.TimeoutObject, _a1 *model.TimeoutObject) {
	_m.Called(_a0, _a1)
}

// OnInvalidTimeoutDetected provides a mock function with given fields: err
func (_m *TimeoutAggregationViolationConsumer) OnInvalidTimeoutDetected(err model.InvalidTimeoutError) {
	_m.Called(err)
}

type mockConstructorTestingTNewTimeoutAggregationViolationConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimeoutAggregationViolationConsumer creates a new instance of TimeoutAggregationViolationConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimeoutAggregationViolationConsumer(t mockConstructorTestingTNewTimeoutAggregationViolationConsumer) *TimeoutAggregationViolationConsumer {
	mock := &TimeoutAggregationViolationConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}