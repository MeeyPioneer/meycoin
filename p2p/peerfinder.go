/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package p2p

import (
	"github.com/meeypioneer/mey-library/log"
	"github.com/meeypioneer/meycoin/message"
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"github.com/meeypioneer/meycoin/p2p/p2putil"
	"github.com/meeypioneer/meycoin/types"
	"time"
)

const (
	macConcurrentQueryCount = 4
)

func NewPeerFinder(logger *log.Logger, pm *peerManager, actorService p2pcommon.ActorService, maxCap int, useDiscover, useWezen bool) p2pcommon.PeerFinder {
	var pf p2pcommon.PeerFinder
	if !useDiscover {
		logger.Info().Msg("peer discover option is disabled, so select static peer finder.")
		pf = &staticPeerFinder{pm:pm, logger:logger}
	} else {
		logger.Info().Bool("useWezen",useWezen).Msg("peer discover option is enabled, so select dynamic peer finder.")
		dp := &dynamicPeerFinder{logger: logger, pm: pm, actorService: actorService, maxCap: maxCap, useWezen:useWezen}
		dp.qStats = make(map[types.PeerID]*queryStat)
		pf = dp
	}
	return pf

}

// staticPeerFinder is for BP or backup node. it will not query to wezen or other peers
type staticPeerFinder struct {
	pm     *peerManager
	logger *log.Logger
}

var _ p2pcommon.PeerFinder = (*staticPeerFinder)(nil)

func (dp *staticPeerFinder) OnPeerDisconnect(peer p2pcommon.RemotePeer) {
}

func (dp *staticPeerFinder) OnPeerConnect(pid types.PeerID) {
}

func (dp *staticPeerFinder) CheckAndFill() {
}


// dynamicPeerFinder is triggering map query to Wezen or address query to other connected peer
// to discover peer
// It is not thread-safe. Thread safety is responsible to the caller.
type dynamicPeerFinder struct {
	logger       *log.Logger
	pm           *peerManager
	actorService p2pcommon.ActorService
	useWezen   bool

	// qStats are logs of query. all connected peers must exist queryStat.
	qStats map[types.PeerID]*queryStat
	maxCap int

	wezenTurn time.Time
}

var _ p2pcommon.PeerFinder = (*dynamicPeerFinder)(nil)

func (dp *dynamicPeerFinder) OnPeerDisconnect(peer p2pcommon.RemotePeer) {
	// And check if to connect more peers
	delete(dp.qStats, peer.ID())
}

func (dp *dynamicPeerFinder) OnPeerConnect(pid types.PeerID) {
	dp.logger.Debug().Str(p2putil.LogPeerID, p2putil.ShortForm(pid)).Msg("check and remove peerID in pool")
	if stat := dp.qStats[pid]; stat == nil {
		// first query will be sent quickly
		dp.qStats[pid] = &queryStat{pid: pid, nextTurn: time.Now().Add(p2pcommon.PeerFirstInterval)}
	}
}

func (dp *dynamicPeerFinder) CheckAndFill() {
	// if enough peer is collected already, skip collect
	toConnCount := dp.maxCap - len(dp.pm.waitingPeers)
	if toConnCount <= 0 {
		return
	}
	now := time.Now()
	// query to wezen
	if dp.useWezen && now.After(dp.wezenTurn) {
		dp.wezenTurn = now.Add(p2pcommon.WezenQueryInterval)
		dp.logger.Debug().Time("next_turn", dp.wezenTurn).Msg("querying to wezen")
		dp.actorService.SendRequest(message.P2PSvc, &message.MapQueryMsg{Count: MaxAddrListSizeWezen})
	}
	// query to peers
	queried := 0
	for _, stat := range dp.qStats {
		if stat.nextTurn.Before(now) {
			// slowly collect
			stat.lastCheck = now
			stat.nextTurn = now.Add(p2pcommon.PeerQueryInterval)
			dp.actorService.SendRequest(message.P2PSvc, &message.GetAddressesMsg{ToWhom: stat.pid, Size: MaxAddrListSizePeer, Offset: 0})
			queried++
			if queried >= macConcurrentQueryCount {
				break
			}
		}
	}
}

type queryStat struct {
	pid       types.PeerID
	lastCheck time.Time
	nextTurn  time.Time
}
