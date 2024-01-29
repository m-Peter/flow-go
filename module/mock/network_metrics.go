// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	channels "github.com/onflow/flow-go/network/channels"
	mock "github.com/stretchr/testify/mock"

	network "github.com/libp2p/go-libp2p/core/network"

	p2pmsg "github.com/onflow/flow-go/network/p2p/message"

	peer "github.com/libp2p/go-libp2p/core/peer"

	protocol "github.com/libp2p/go-libp2p/core/protocol"

	time "time"
)

// NetworkMetrics is an autogenerated mock type for the NetworkMetrics type
type NetworkMetrics struct {
	mock.Mock
}

// AllowConn provides a mock function with given fields: dir, usefd
func (_m *NetworkMetrics) AllowConn(dir network.Direction, usefd bool) {
	_m.Called(dir, usefd)
}

// AllowMemory provides a mock function with given fields: size
func (_m *NetworkMetrics) AllowMemory(size int) {
	_m.Called(size)
}

// AllowPeer provides a mock function with given fields: p
func (_m *NetworkMetrics) AllowPeer(p peer.ID) {
	_m.Called(p)
}

// AllowProtocol provides a mock function with given fields: proto
func (_m *NetworkMetrics) AllowProtocol(proto protocol.ID) {
	_m.Called(proto)
}

// AllowService provides a mock function with given fields: svc
func (_m *NetworkMetrics) AllowService(svc string) {
	_m.Called(svc)
}

// AllowStream provides a mock function with given fields: p, dir
func (_m *NetworkMetrics) AllowStream(p peer.ID, dir network.Direction) {
	_m.Called(p, dir)
}

// AsyncProcessingFinished provides a mock function with given fields: duration
func (_m *NetworkMetrics) AsyncProcessingFinished(duration time.Duration) {
	_m.Called(duration)
}

// AsyncProcessingStarted provides a mock function with given fields:
func (_m *NetworkMetrics) AsyncProcessingStarted() {
	_m.Called()
}

// BlockConn provides a mock function with given fields: dir, usefd
func (_m *NetworkMetrics) BlockConn(dir network.Direction, usefd bool) {
	_m.Called(dir, usefd)
}

// BlockMemory provides a mock function with given fields: size
func (_m *NetworkMetrics) BlockMemory(size int) {
	_m.Called(size)
}

// BlockPeer provides a mock function with given fields: p
func (_m *NetworkMetrics) BlockPeer(p peer.ID) {
	_m.Called(p)
}

// BlockProtocol provides a mock function with given fields: proto
func (_m *NetworkMetrics) BlockProtocol(proto protocol.ID) {
	_m.Called(proto)
}

// BlockProtocolPeer provides a mock function with given fields: proto, p
func (_m *NetworkMetrics) BlockProtocolPeer(proto protocol.ID, p peer.ID) {
	_m.Called(proto, p)
}

// BlockService provides a mock function with given fields: svc
func (_m *NetworkMetrics) BlockService(svc string) {
	_m.Called(svc)
}

// BlockServicePeer provides a mock function with given fields: svc, p
func (_m *NetworkMetrics) BlockServicePeer(svc string, p peer.ID) {
	_m.Called(svc, p)
}

// BlockStream provides a mock function with given fields: p, dir
func (_m *NetworkMetrics) BlockStream(p peer.ID, dir network.Direction) {
	_m.Called(p, dir)
}

// DNSLookupDuration provides a mock function with given fields: duration
func (_m *NetworkMetrics) DNSLookupDuration(duration time.Duration) {
	_m.Called(duration)
}

// DuplicateInboundMessagesDropped provides a mock function with given fields: topic, _a1, messageType
func (_m *NetworkMetrics) DuplicateInboundMessagesDropped(topic string, _a1 string, messageType string) {
	_m.Called(topic, _a1, messageType)
}

// DuplicateMessagePenalties provides a mock function with given fields: penalty
func (_m *NetworkMetrics) DuplicateMessagePenalties(penalty float64) {
	_m.Called(penalty)
}

// DuplicateMessagesCounts provides a mock function with given fields: count
func (_m *NetworkMetrics) DuplicateMessagesCounts(count float64) {
	_m.Called(count)
}

// InboundConnections provides a mock function with given fields: connectionCount
func (_m *NetworkMetrics) InboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// InboundMessageReceived provides a mock function with given fields: sizeBytes, topic, _a2, messageType
func (_m *NetworkMetrics) InboundMessageReceived(sizeBytes int, topic string, _a2 string, messageType string) {
	_m.Called(sizeBytes, topic, _a2, messageType)
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

// OnAppSpecificScoreUpdated provides a mock function with given fields: _a0
func (_m *NetworkMetrics) OnAppSpecificScoreUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnBehaviourPenaltyUpdated provides a mock function with given fields: _a0
func (_m *NetworkMetrics) OnBehaviourPenaltyUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnControlMessagesTruncated provides a mock function with given fields: messageType, diff
func (_m *NetworkMetrics) OnControlMessagesTruncated(messageType p2pmsg.ControlMessageType, diff int) {
	_m.Called(messageType, diff)
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

// OnDialRetryBudgetResetToDefault provides a mock function with given fields:
func (_m *NetworkMetrics) OnDialRetryBudgetResetToDefault() {
	_m.Called()
}

// OnDialRetryBudgetUpdated provides a mock function with given fields: budget
func (_m *NetworkMetrics) OnDialRetryBudgetUpdated(budget uint64) {
	_m.Called(budget)
}

// OnEstablishStreamFailure provides a mock function with given fields: duration, attempts
func (_m *NetworkMetrics) OnEstablishStreamFailure(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnFirstMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *NetworkMetrics) OnFirstMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnIHaveControlMessageIdsTruncated provides a mock function with given fields: diff
func (_m *NetworkMetrics) OnIHaveControlMessageIdsTruncated(diff int) {
	_m.Called(diff)
}

// OnIHaveMessageIDsReceived provides a mock function with given fields: channel, msgIdCount
func (_m *NetworkMetrics) OnIHaveMessageIDsReceived(channel string, msgIdCount int) {
	_m.Called(channel, msgIdCount)
}

// OnIPColocationFactorUpdated provides a mock function with given fields: _a0
func (_m *NetworkMetrics) OnIPColocationFactorUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnIWantControlMessageIdsTruncated provides a mock function with given fields: diff
func (_m *NetworkMetrics) OnIWantControlMessageIdsTruncated(diff int) {
	_m.Called(diff)
}

// OnIWantMessageIDsReceived provides a mock function with given fields: msgIdCount
func (_m *NetworkMetrics) OnIWantMessageIDsReceived(msgIdCount int) {
	_m.Called(msgIdCount)
}

// OnIncomingRpcReceived provides a mock function with given fields: iHaveCount, iWantCount, graftCount, pruneCount, msgCount
func (_m *NetworkMetrics) OnIncomingRpcReceived(iHaveCount int, iWantCount int, graftCount int, pruneCount int, msgCount int) {
	_m.Called(iHaveCount, iWantCount, graftCount, pruneCount, msgCount)
}

// OnInvalidMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *NetworkMetrics) OnInvalidMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnLocalMeshSizeUpdated provides a mock function with given fields: topic, size
func (_m *NetworkMetrics) OnLocalMeshSizeUpdated(topic string, size int) {
	_m.Called(topic, size)
}

// OnLocalPeerJoinedTopic provides a mock function with given fields:
func (_m *NetworkMetrics) OnLocalPeerJoinedTopic() {
	_m.Called()
}

// OnLocalPeerLeftTopic provides a mock function with given fields:
func (_m *NetworkMetrics) OnLocalPeerLeftTopic() {
	_m.Called()
}

// OnMeshMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *NetworkMetrics) OnMeshMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnMessageDeliveredToAllSubscribers provides a mock function with given fields: size
func (_m *NetworkMetrics) OnMessageDeliveredToAllSubscribers(size int) {
	_m.Called(size)
}

// OnMessageDuplicate provides a mock function with given fields: size
func (_m *NetworkMetrics) OnMessageDuplicate(size int) {
	_m.Called(size)
}

// OnMessageEnteredValidation provides a mock function with given fields: size
func (_m *NetworkMetrics) OnMessageEnteredValidation(size int) {
	_m.Called(size)
}

// OnMessageRejected provides a mock function with given fields: size, reason
func (_m *NetworkMetrics) OnMessageRejected(size int, reason string) {
	_m.Called(size, reason)
}

// OnMisbehaviorReported provides a mock function with given fields: channel, misbehaviorType
func (_m *NetworkMetrics) OnMisbehaviorReported(channel string, misbehaviorType string) {
	_m.Called(channel, misbehaviorType)
}

// OnOutboundRpcDropped provides a mock function with given fields:
func (_m *NetworkMetrics) OnOutboundRpcDropped() {
	_m.Called()
}

// OnOverallPeerScoreUpdated provides a mock function with given fields: _a0
func (_m *NetworkMetrics) OnOverallPeerScoreUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnPeerAddedToProtocol provides a mock function with given fields: _a0
func (_m *NetworkMetrics) OnPeerAddedToProtocol(_a0 string) {
	_m.Called(_a0)
}

// OnPeerDialFailure provides a mock function with given fields: duration, attempts
func (_m *NetworkMetrics) OnPeerDialFailure(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnPeerDialed provides a mock function with given fields: duration, attempts
func (_m *NetworkMetrics) OnPeerDialed(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnPeerGraftTopic provides a mock function with given fields: topic
func (_m *NetworkMetrics) OnPeerGraftTopic(topic string) {
	_m.Called(topic)
}

// OnPeerPruneTopic provides a mock function with given fields: topic
func (_m *NetworkMetrics) OnPeerPruneTopic(topic string) {
	_m.Called(topic)
}

// OnPeerRemovedFromProtocol provides a mock function with given fields:
func (_m *NetworkMetrics) OnPeerRemovedFromProtocol() {
	_m.Called()
}

// OnPeerThrottled provides a mock function with given fields:
func (_m *NetworkMetrics) OnPeerThrottled() {
	_m.Called()
}

// OnRateLimitedPeer provides a mock function with given fields: pid, role, msgType, topic, reason
func (_m *NetworkMetrics) OnRateLimitedPeer(pid peer.ID, role string, msgType string, topic string, reason string) {
	_m.Called(pid, role, msgType, topic, reason)
}

// OnRpcReceived provides a mock function with given fields: msgCount, iHaveCount, iWantCount, graftCount, pruneCount
func (_m *NetworkMetrics) OnRpcReceived(msgCount int, iHaveCount int, iWantCount int, graftCount int, pruneCount int) {
	_m.Called(msgCount, iHaveCount, iWantCount, graftCount, pruneCount)
}

// OnRpcSent provides a mock function with given fields: msgCount, iHaveCount, iWantCount, graftCount, pruneCount
func (_m *NetworkMetrics) OnRpcSent(msgCount int, iHaveCount int, iWantCount int, graftCount int, pruneCount int) {
	_m.Called(msgCount, iHaveCount, iWantCount, graftCount, pruneCount)
}

// OnStreamCreated provides a mock function with given fields: duration, attempts
func (_m *NetworkMetrics) OnStreamCreated(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnStreamCreationFailure provides a mock function with given fields: duration, attempts
func (_m *NetworkMetrics) OnStreamCreationFailure(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnStreamCreationRetryBudgetResetToDefault provides a mock function with given fields:
func (_m *NetworkMetrics) OnStreamCreationRetryBudgetResetToDefault() {
	_m.Called()
}

// OnStreamCreationRetryBudgetUpdated provides a mock function with given fields: budget
func (_m *NetworkMetrics) OnStreamCreationRetryBudgetUpdated(budget uint64) {
	_m.Called(budget)
}

// OnStreamEstablished provides a mock function with given fields: duration, attempts
func (_m *NetworkMetrics) OnStreamEstablished(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnTimeInMeshUpdated provides a mock function with given fields: _a0, _a1
func (_m *NetworkMetrics) OnTimeInMeshUpdated(_a0 channels.Topic, _a1 time.Duration) {
	_m.Called(_a0, _a1)
}

// OnUnauthorizedMessage provides a mock function with given fields: role, msgType, topic, offense
func (_m *NetworkMetrics) OnUnauthorizedMessage(role string, msgType string, topic string, offense string) {
	_m.Called(role, msgType, topic, offense)
}

// OnUndeliveredMessage provides a mock function with given fields:
func (_m *NetworkMetrics) OnUndeliveredMessage() {
	_m.Called()
}

// OnViolationReportSkipped provides a mock function with given fields:
func (_m *NetworkMetrics) OnViolationReportSkipped() {
	_m.Called()
}

// OutboundConnections provides a mock function with given fields: connectionCount
func (_m *NetworkMetrics) OutboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// OutboundMessageSent provides a mock function with given fields: sizeBytes, topic, _a2, messageType
func (_m *NetworkMetrics) OutboundMessageSent(sizeBytes int, topic string, _a2 string, messageType string) {
	_m.Called(sizeBytes, topic, _a2, messageType)
}

// QueueDuration provides a mock function with given fields: duration, priority
func (_m *NetworkMetrics) QueueDuration(duration time.Duration, priority int) {
	_m.Called(duration, priority)
}

// RoutingTablePeerAdded provides a mock function with given fields:
func (_m *NetworkMetrics) RoutingTablePeerAdded() {
	_m.Called()
}

// RoutingTablePeerRemoved provides a mock function with given fields:
func (_m *NetworkMetrics) RoutingTablePeerRemoved() {
	_m.Called()
}

// SetWarningStateCount provides a mock function with given fields: _a0
func (_m *NetworkMetrics) SetWarningStateCount(_a0 uint) {
	_m.Called(_a0)
}

// UnicastMessageSendingCompleted provides a mock function with given fields: topic
func (_m *NetworkMetrics) UnicastMessageSendingCompleted(topic string) {
	_m.Called(topic)
}

// UnicastMessageSendingStarted provides a mock function with given fields: topic
func (_m *NetworkMetrics) UnicastMessageSendingStarted(topic string) {
	_m.Called(topic)
}

type mockConstructorTestingTNewNetworkMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewNetworkMetrics creates a new instance of NetworkMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetworkMetrics(t mockConstructorTestingTNewNetworkMetrics) *NetworkMetrics {
	mock := &NetworkMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
