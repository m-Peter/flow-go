// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	flow "github.com/onflow/flow-go/model/flow"
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"

	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// GossipSubMsgValidationRpcInspector is an autogenerated mock type for the GossipSubMsgValidationRpcInspector type
type GossipSubMsgValidationRpcInspector struct {
	mock.Mock
}

// Done provides a mock function with given fields:
func (_m *GossipSubMsgValidationRpcInspector) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Inspect provides a mock function with given fields: _a0, _a1
func (_m *GossipSubMsgValidationRpcInspector) Inspect(_a0 peer.ID, _a1 *pubsub.RPC) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(peer.ID, *pubsub.RPC) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *GossipSubMsgValidationRpcInspector) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OnClusterIdsUpdated provides a mock function with given fields: _a0
func (_m *GossipSubMsgValidationRpcInspector) OnClusterIdsUpdated(_a0 flow.ChainIDList) {
	_m.Called(_a0)
}

// Ready provides a mock function with given fields:
func (_m *GossipSubMsgValidationRpcInspector) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *GossipSubMsgValidationRpcInspector) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewGossipSubMsgValidationRpcInspector interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipSubMsgValidationRpcInspector creates a new instance of GossipSubMsgValidationRpcInspector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipSubMsgValidationRpcInspector(t mockConstructorTestingTNewGossipSubMsgValidationRpcInspector) *GossipSubMsgValidationRpcInspector {
	mock := &GossipSubMsgValidationRpcInspector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
