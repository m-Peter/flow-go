// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"

	time "time"
)

// PaceMaker is an autogenerated mock type for the PaceMaker type
type PaceMaker struct {
	mock.Mock
}

// BlockRateDelay provides a mock function with given fields:
func (_m *PaceMaker) BlockRateDelay() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// CurView provides a mock function with given fields:
func (_m *PaceMaker) CurView() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// LastViewTC provides a mock function with given fields:
func (_m *PaceMaker) LastViewTC() *flow.TimeoutCertificate {
	ret := _m.Called()

	var r0 *flow.TimeoutCertificate
	if rf, ok := ret.Get(0).(func() *flow.TimeoutCertificate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.TimeoutCertificate)
		}
	}

	return r0
}

// NewestQC provides a mock function with given fields:
func (_m *PaceMaker) NewestQC() *flow.QuorumCertificate {
	ret := _m.Called()

	var r0 *flow.QuorumCertificate
	if rf, ok := ret.Get(0).(func() *flow.QuorumCertificate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.QuorumCertificate)
		}
	}

	return r0
}

// ProcessQC provides a mock function with given fields: qc
func (_m *PaceMaker) ProcessQC(qc *flow.QuorumCertificate) (*model.NewViewEvent, error) {
	ret := _m.Called(qc)

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func(*flow.QuorumCertificate) *model.NewViewEvent); ok {
		r0 = rf(qc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*flow.QuorumCertificate) error); ok {
		r1 = rf(qc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessTC provides a mock function with given fields: tc
func (_m *PaceMaker) ProcessTC(tc *flow.TimeoutCertificate) (*model.NewViewEvent, error) {
	ret := _m.Called(tc)

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func(*flow.TimeoutCertificate) *model.NewViewEvent); ok {
		r0 = rf(tc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*flow.TimeoutCertificate) error); ok {
		r1 = rf(tc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: ctx
func (_m *PaceMaker) Start(ctx context.Context) {
	_m.Called(ctx)
}

// TimeoutChannel provides a mock function with given fields:
func (_m *PaceMaker) TimeoutChannel() <-chan time.Time {
	ret := _m.Called()

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func() <-chan time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

type mockConstructorTestingTNewPaceMaker interface {
	mock.TestingT
	Cleanup(func())
}

// NewPaceMaker creates a new instance of PaceMaker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPaceMaker(t mockConstructorTestingTNewPaceMaker) *PaceMaker {
	mock := &PaceMaker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
