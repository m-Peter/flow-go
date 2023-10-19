package state_stream

import (
	"context"
	"time"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/executiondatasync/execution_data"
)

// API represents an interface that defines methods for interacting with a blockchain's execution data and events.
type API interface {
	// GetExecutionDataByBlockID retrieves execution data for a specific block by its block ID.
	GetExecutionDataByBlockID(ctx context.Context, blockID flow.Identifier) (*execution_data.BlockExecutionData, error)
	// SubscribeExecutionData subscribes to execution data starting from a specific block ID and block height.
	SubscribeExecutionData(ctx context.Context, startBlockID flow.Identifier, startBlockHeight uint64) Subscription
	// SubscribeEvents subscribes to events starting from a specific block ID and block height, with an optional event filter.
	SubscribeEvents(ctx context.Context, startBlockID flow.Identifier, startHeight uint64, filter EventFilter) Subscription
}

// Subscription represents a streaming request, and handles the communication between the grpc handler
// and the backend implementation.
type Subscription interface {
	// ID returns the unique identifier for this subscription used for logging
	ID() string

	// Channel returns the channel from which subscription data can be read
	Channel() <-chan interface{}

	// Err returns the error that caused the subscription to fail
	Err() error
}

// Streamable represents a subscription that can be streamed.
type Streamable interface {
	ID() string
	Close()
	Fail(error)
	Send(context.Context, interface{}, time.Duration) error
	Next(context.Context) (interface{}, error)
}
