// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
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

// HighestQC provides a mock function with given fields:
func (_m *PaceMaker) HighestQC() *flow.QuorumCertificate {
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

// OnPartialTC provides a mock function with given fields: curView
func (_m *PaceMaker) OnPartialTC(curView uint64) {
	_m.Called(curView)
}

// ProcessQC provides a mock function with given fields: qc
func (_m *PaceMaker) ProcessQC(qc *flow.QuorumCertificate) (*model.NewViewEvent, bool) {
	ret := _m.Called(qc)

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func(*flow.QuorumCertificate) *model.NewViewEvent); ok {
		r0 = rf(qc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*flow.QuorumCertificate) bool); ok {
		r1 = rf(qc)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ProcessTC provides a mock function with given fields: tc
func (_m *PaceMaker) ProcessTC(tc *flow.TimeoutCertificate) (*model.NewViewEvent, bool) {
	ret := _m.Called(tc)

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func(*flow.TimeoutCertificate) *model.NewViewEvent); ok {
		r0 = rf(tc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*flow.TimeoutCertificate) bool); ok {
		r1 = rf(tc)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *PaceMaker) Start() {
	_m.Called()
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
