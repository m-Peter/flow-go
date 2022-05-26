// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PingMetrics is an autogenerated mock type for the PingMetrics type
type PingMetrics struct {
	mock.Mock
}

// NodeInfo provides a mock function with given fields: node, nodeInfo, version, sealedHeight, hotstuffCurView
func (_m *PingMetrics) NodeInfo(node *flow.Identity, nodeInfo string, version string, sealedHeight uint64, hotstuffCurView uint64) {
	_m.Called(node, nodeInfo, version, sealedHeight, hotstuffCurView)
}

// NodeReachable provides a mock function with given fields: node, nodeInfo, rtt
func (_m *PingMetrics) NodeReachable(node *flow.Identity, nodeInfo string, rtt time.Duration) {
	_m.Called(node, nodeInfo, rtt)
}

type NewPingMetricsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewPingMetrics creates a new instance of PingMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPingMetrics(t NewPingMetricsT) *PingMetrics {
	mock := &PingMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
