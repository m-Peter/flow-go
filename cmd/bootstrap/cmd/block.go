package cmd

import (
	"github.com/dapperlabs/flow-go/cmd/bootstrap/run"
	model "github.com/dapperlabs/flow-go/model/bootstrap"
	"github.com/dapperlabs/flow-go/model/flow"
)

func constructGenesisBlock(nodes []model.NodeInfo) *flow.Block {
	identityList := generateIdentityList(nodes)
	block := run.GenerateRootBlock(identityList)

	writeJSON(model.PathGenesisBlock, block)

	return block
}

func generateIdentityList(nodes []model.NodeInfo) flow.IdentityList {

	list := make([]*flow.Identity, 0, len(nodes))

	for _, node := range nodes {
		ident := node.Identity()
		list = append(list, ident)
	}

	return list
}
