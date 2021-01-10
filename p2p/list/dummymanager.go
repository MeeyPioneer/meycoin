/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package list

import (
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"github.com/meeypioneer/meycoin/types"
	"time"
)

// dummyListManager allows all remote nodes
type dummyListManager struct {

}

func newDummyListManager() p2pcommon.ListManager {
	return &dummyListManager{}
}


func (*dummyListManager) Start() {
}

func (*dummyListManager) Stop() {
}


func (*dummyListManager) IsBanned(addr string, pid types.PeerID) (bool, time.Time) {
	return false, UndefinedTime
}

func (*dummyListManager) RefineList() {
}

func (*dummyListManager) Summary() map[string]interface{} {
	sum := make(map[string]interface{})
	return sum
}
