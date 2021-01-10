/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package message

import (
	"github.com/meeypioneer/meycoin/types"
)

const WezenRPCSvc = "pRpcSvc"

type MapQueryMsg struct {
	Count int
	BestBlock *types.Block
}

type MapQueryRsp struct {
	Peers []*types.PeerAddress
	Err error
}


