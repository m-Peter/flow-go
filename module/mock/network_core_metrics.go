// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	time "time"
)

// NetworkCoreMetrics is an autogenerated mock type for the NetworkCoreMetrics type
type NetworkCoreMetrics struct {
	mock.Mock
}

// DuplicateInboundMessagesDropped provides a mock function with given fields: topic, protocol, messageType
func (_m *NetworkCoreMetrics) DuplicateInboundMessagesDropped(topic string, protocol string, messageType string) {
	_m.Called(topic, protocol, messageType)
}

// InboundMessageReceived provides a mock function with given fields: sizeBytes, topic, protocol, messageType
func (_m *NetworkCoreMetrics) InboundMessageReceived(sizeBytes int, topic string, protocol string, messageType string) {
	_m.Called(sizeBytes, topic, protocol, messageType)
}

// MessageAdded provides a mock function with given fields: priority
func (_m *NetworkCoreMetrics) MessageAdded(priority int) {
	_m.Called(priority)
}

// MessageProcessingFinished provides a mock function with given fields: topic, duration
func (_m *NetworkCoreMetrics) MessageProcessingFinished(topic string, duration time.Duration) {
	_m.Called(topic, duration)
}

// MessageProcessingStarted provides a mock function with given fields: topic
func (_m *NetworkCoreMetrics) MessageProcessingStarted(topic string) {
	_m.Called(topic)
}

// MessageRemoved provides a mock function with given fields: priority
func (_m *NetworkCoreMetrics) MessageRemoved(priority int) {
	_m.Called(priority)
}

// OnMisbehaviorReported provides a mock function with given fields: channel, misbehaviorType
func (_m *NetworkCoreMetrics) OnMisbehaviorReported(channel string, misbehaviorType string) {
	_m.Called(channel, misbehaviorType)
}

// OnRateLimitedPeer provides a mock function with given fields: pid, role, msgType, topic, reason
func (_m *NetworkCoreMetrics) OnRateLimitedPeer(pid peer.ID, role string, msgType string, topic string, reason string) {
	_m.Called(pid, role, msgType, topic, reason)
}

// OnUnauthorizedMessage provides a mock function with given fields: role, msgType, topic, offense
func (_m *NetworkCoreMetrics) OnUnauthorizedMessage(role string, msgType string, topic string, offense string) {
	_m.Called(role, msgType, topic, offense)
}

// OnViolationReportSkipped provides a mock function with given fields:
func (_m *NetworkCoreMetrics) OnViolationReportSkipped() {
	_m.Called()
}

// OutboundMessageSent provides a mock function with given fields: sizeBytes, topic, protocol, messageType
func (_m *NetworkCoreMetrics) OutboundMessageSent(sizeBytes int, topic string, protocol string, messageType string) {
	_m.Called(sizeBytes, topic, protocol, messageType)
}

// QueueDuration provides a mock function with given fields: duration, priority
func (_m *NetworkCoreMetrics) QueueDuration(duration time.Duration, priority int) {
	_m.Called(duration, priority)
}

// UnicastMessageSendingCompleted provides a mock function with given fields: topic
func (_m *NetworkCoreMetrics) UnicastMessageSendingCompleted(topic string) {
	_m.Called(topic)
}

// UnicastMessageSendingStarted provides a mock function with given fields: topic
func (_m *NetworkCoreMetrics) UnicastMessageSendingStarted(topic string) {
	_m.Called(topic)
}

// NewNetworkCoreMetrics creates a new instance of NetworkCoreMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetworkCoreMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetworkCoreMetrics {
	mock := &NetworkCoreMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
