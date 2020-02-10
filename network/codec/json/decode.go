// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package json

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/dapperlabs/flow-go/engine/execution"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/libp2p/message"
	"github.com/dapperlabs/flow-go/model/messages"
	"github.com/dapperlabs/flow-go/model/trickle"
)

// decode will decode the envelope into an entity.
func decode(env Envelope) (interface{}, error) {

	// create the desired message
	var v interface{}
	switch env.Code {

	// trickle overlay network
	case CodePing:
		v = &trickle.Ping{}
	case CodePong:
		v = &trickle.Pong{}
	case CodeAuth:
		v = &trickle.Auth{}
	case CodeAnnounce:
		v = &trickle.Announce{}
	case CodeRequest:
		v = &trickle.Request{}
	case CodeResponse:
		v = &trickle.Response{}

	case CodeCollectionGuarantee:
		v = &flow.CollectionGuarantee{}
	case CodeTransaction:
		v = &flow.Transaction{}

	case CodeBlock:
		v = &flow.Block{}

	case CodeCollectionRequest:
		v = &messages.CollectionRequest{}
	case CodeCollectionResponse:
		v = &messages.CollectionResponse{}

	case CodeEcho:
		v = &message.Echo{}

	case CodeExecutionReceipt:
		v = &flow.ExecutionReceipt{}
	case CodeExecutionStateRequest:
		v = &messages.ExecutionStateRequest{}
	case CodeExecutionStateResponse:
		v = &messages.ExecutionStateResponse{}
	case CodeExecutionStateSyncRequest:
		v = &messages.ExecutionStateSyncRequest{}
	case CodeExecutionStateDelta:
		v = &messages.ExecutionStateDelta{}
	case CodeExecutionCompleteBlock:
		v = &execution.CompleteBlock{}
	case CodeExecutionComputationOrder:
		v = &execution.ComputationOrder{}

	default:
		return nil, errors.Errorf("invalid message code (%d)", env.Code)
	}

	// unmarshal the payload
	err := json.Unmarshal(env.Data, v)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode payload")
	}

	return v, nil
}
