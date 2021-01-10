/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package common

import (
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	core "github.com/libp2p/go-libp2p-core"
	"time"
)

const (
	// port for RPC
	DefaultRPCPort = 8915
	// port for query and register meycoinsvr
	DefaultSrvPort = 8916
)

// subprotocol for wezen
const (
	WezenMapSub  core.ProtocolID = "/wezen/0.1"
	WezenPingSub core.ProtocolID = "/ping/0.1"
)
const (
	MapQuery p2pcommon.SubProtocol = 0x0100 + iota
	MapResponse
)

const WezenConnectionTTL = time.Second * 30

// Additional messages of wezen response
const (
	TooOldVersionMsg = "too old version"
)

