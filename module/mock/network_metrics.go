// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// NetworkMetrics is an autogenerated mock type for the NetworkMetrics type
type NetworkMetrics struct {
	mock.Mock
}

// DNSLookupDuration provides a mock function with given fields: duration
func (_m *NetworkMetrics) DNSLookupDuration(duration time.Duration) {
	_m.Called(duration)
}

// DirectMessageFinished provides a mock function with given fields: topic
func (_m *NetworkMetrics) DirectMessageFinished(topic string) {
	_m.Called(topic)
}

// DirectMessageStarted provides a mock function with given fields: topic
func (_m *NetworkMetrics) DirectMessageStarted(topic string) {
	_m.Called(topic)
}

// InboundConnections provides a mock function with given fields: connectionCount
func (_m *NetworkMetrics) InboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// MessageAdded provides a mock function with given fields: priority
func (_m *NetworkMetrics) MessageAdded(priority int) {
	_m.Called(priority)
}

// MessageProcessingFinished provides a mock function with given fields: topic, duration
func (_m *NetworkMetrics) MessageProcessingFinished(topic string, duration time.Duration) {
	_m.Called(topic, duration)
}

// MessageProcessingStarted provides a mock function with given fields: topic
func (_m *NetworkMetrics) MessageProcessingStarted(topic string) {
	_m.Called(topic)
}

// MessageRemoved provides a mock function with given fields: priority
func (_m *NetworkMetrics) MessageRemoved(priority int) {
	_m.Called(priority)
}

// NetworkDuplicateMessagesDropped provides a mock function with given fields: topic, messageType
func (_m *NetworkMetrics) NetworkDuplicateMessagesDropped(topic string, messageType string) {
	_m.Called(topic, messageType)
}

// NetworkMessageReceived provides a mock function with given fields: sizeBytes, topic, messageType
func (_m *NetworkMetrics) NetworkMessageReceived(sizeBytes int, topic string, messageType string) {
	_m.Called(sizeBytes, topic, messageType)
}

// NetworkMessageSent provides a mock function with given fields: sizeBytes, topic, messageType
func (_m *NetworkMetrics) NetworkMessageSent(sizeBytes int, topic string, messageType string) {
	_m.Called(sizeBytes, topic, messageType)
}

// OnDNSCacheHit provides a mock function with given fields:
func (_m *NetworkMetrics) OnDNSCacheHit() {
	_m.Called()
}

// OnDNSCacheInvalidated provides a mock function with given fields:
func (_m *NetworkMetrics) OnDNSCacheInvalidated() {
	_m.Called()
}

// OnDNSCacheMiss provides a mock function with given fields:
func (_m *NetworkMetrics) OnDNSCacheMiss() {
	_m.Called()
}

// OnDNSLookupRequestDropped provides a mock function with given fields:
func (_m *NetworkMetrics) OnDNSLookupRequestDropped() {
	_m.Called()
}

// OutboundConnections provides a mock function with given fields: connectionCount
func (_m *NetworkMetrics) OutboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// QueueDuration provides a mock function with given fields: duration, priority
func (_m *NetworkMetrics) QueueDuration(duration time.Duration, priority int) {
	_m.Called(duration, priority)
}

type NewNetworkMetricsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNetworkMetrics creates a new instance of NetworkMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetworkMetrics(t NewNetworkMetricsT) *NetworkMetrics {
	mock := &NetworkMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
