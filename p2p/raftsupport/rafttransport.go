/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package raftsupport

import (
	"github.com/meeypioneer/mey-library/log"
	"github.com/meeypioneer/meycoin/consensus"
	"github.com/meeypioneer/meycoin/consensus/impl/raftv2"
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"github.com/meeypioneer/meycoin/p2p/p2putil"
	"github.com/meeypioneer/meycoin/types"
	"github.com/meeypioneer/etcd/etcdserver/stats"
	rtypes "github.com/meeypioneer/etcd/pkg/types"
	"github.com/meeypioneer/etcd/raft"
	"github.com/meeypioneer/etcd/raft/raftpb"
	"github.com/meeypioneer/etcd/snap"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// errors
var (
	errInvalidMember     = errors.New("invalid member id")
	errCanNotFoundMember = errors.New("cannot find member")
	errRemovedMember     = errors.New("member was removed")
	errUnreachableMember = errors.New("member is unreachable")
)
// MeyCoinRaftTransport is wrapper of p2p module
type MeyCoinRaftTransport struct {
	logger *log.Logger
	mutex  sync.RWMutex

	nt      p2pcommon.NetworkTransport
	pm      p2pcommon.PeerManager
	mf      p2pcommon.MoFactory
	consAcc consensus.ConsensusAccessor
	raftAcc consensus.MeyCoinRaftAccessor
	snapF   SnapshotIOFactory

	cluster *raftv2.Cluster

	// statuses have connection status of memeber peers
	statuses map[rtypes.ID]*rPeerStatus
	stByPID  map[types.PeerID]*rPeerStatus

	// copied from original transport
	ServerStats *stats.ServerStats
	LeaderStats *stats.LeaderStats
}

var _ raftv2.Transporter = (*MeyCoinRaftTransport)(nil)

func NewMeyCoinRaftTransport(logger *log.Logger, nt p2pcommon.NetworkTransport, pm p2pcommon.PeerManager, mf p2pcommon.MoFactory, consAcc consensus.ConsensusAccessor, cluster interface{}) *MeyCoinRaftTransport {
	t := &MeyCoinRaftTransport{logger: logger, nt: nt, pm: pm, mf: mf, consAcc: consAcc, raftAcc: consAcc.RaftAccessor(), cluster: cluster.(*raftv2.Cluster),
		statuses:    make(map[rtypes.ID]*rPeerStatus),
		stByPID:     make(map[types.PeerID]*rPeerStatus),
		ServerStats: stats.NewServerStats("", ""),
	}
	// TODO need check real id type
	t.LeaderStats = stats.NewLeaderStats(strconv.Itoa(int(t.cluster.NodeID())))
	t.snapF = t
	pm.AddPeerEventListener(t)
	logger.Info().Msg("meycoin raft transport is created")
	return t
}

func (t *MeyCoinRaftTransport) Start() error {
	t.nt.AddStreamHandler(p2pcommon.RaftSnapSubAddr, t.OnRaftSnapshot)

	// do nothing for now
	return nil
}

func (t *MeyCoinRaftTransport) Handler() http.Handler {
	//
	return http.NewServeMux()
}

// Send must send message to target peer or report unreachable if sending peer is failed.
func (t *MeyCoinRaftTransport) Send(msgs []raftpb.Message) {
	for _, m := range msgs {
		if m.To == 0 {
			// ignore intentionally dropped message
			continue
		}

		member := t.raftAcc.GetMemberByID(m.To)
		if member == nil {
			t.logger.Info().Object("raftMsg", &RaftMsgMarshaller{&m}).Msg("ignored message to no raft member")

			continue
		}
		peer, _ := t.pm.GetPeer(member.GetPeerID())
		if peer != nil {
			peer.SendMessage(t.mf.NewRaftMsgOrder(m.Type, &m))
			continue
		} else {
			t.logger.Debug().Str(p2putil.LogPeerID, p2putil.ShortForm(member.GetPeerID())).Msg("peer is unreachable")
			t.raftAcc.ReportUnreachable(member.GetPeerID())
			continue
		}

		t.logger.Debug().Str(p2putil.LogPeerID, p2putil.ShortForm(member.GetPeerID())).Object("raftMsg", &RaftMsgMarshaller{&m}).Msg("can't send message to unconnected peer")
	}
}

func (t *MeyCoinRaftTransport) SendSnapshot(m snap.Message) {
	if m.To == 0 {
		// ignore intentionally dropped message
		t.logger.Warn().Msg("drop snap message: to invalid target")
		m.CloseWithError(errInvalidMember)
		return
	}
	member := t.raftAcc.GetMemberByID(m.To)
	if member == nil {
		// TODO is it ok to ignore?
		t.logger.Warn().Msg("drop snap message: no member")
		m.CloseWithError(errCanNotFoundMember)
		return
	}
	// TODO: member is exists but unreachable should return message to change peer state
	peer, _ := t.pm.GetPeer(member.GetPeerID())
	if peer == nil {
		t.logger.Warn().Msg("drop snap message:no peer")
		t.raftAcc.ReportUnreachable(member.GetPeerID())
		t.raftAcc.ReportSnapshot(member.GetPeerID(), raft.SnapshotFailure)
		m.CloseWithError(errUnreachableMember)
		return
	}

	sender := t.snapF.NewSnapshotSender(peer)
	go sender.Send(&m)
}

func (t *MeyCoinRaftTransport) AddRemote(id rtypes.ID, urls []string) {
	panic("implement me")
}

func (t *MeyCoinRaftTransport) AddPeer(id rtypes.ID, peerID types.PeerID, urls []string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.logger.Debug().Str(p2putil.LogPeerID, p2putil.ShortForm(peerID)).Str("id", id.String()).Msg("Adding member peer")

	member := t.raftAcc.GetMemberByID(uint64(id))
	if member == nil {
		t.logger.Info().Str(p2putil.LogPeerID, p2putil.ShortForm(peerID)).Str("id", id.String()).Msg("can't find member")
		return
	}
	st, exist := t.statuses[id]
	if exist {
		if _, exist := t.pm.GetPeer(member.GetPeerID()); exist {
			t.logger.Info().Str(p2putil.LogPeerID, p2putil.ShortForm(peerID)).Str("id", id.String()).Msg("peer already exists")
			st.activate()
			return
		} else {
			st.deactivate("disconnected")
			t.connectToPeer(member)
		}
	} else {
		// register new status and try to connect
		t.statuses[id] = newPeerStatus(id, peerID, t.logger)
		t.stByPID[peerID] = t.statuses[id]
		t.connectToPeer(member)
	}
}

func (t *MeyCoinRaftTransport) connectToPeer(member *consensus.Member) {
	pid, err := types.IDFromBytes(member.PeerID)
	peerMeta, err := p2putil.FromMultiAddrStringWithPID(member.Address, pid)
	if err != nil {
		t.logger.Panic().Err(err).Str("addr", member.Address).Msg("Address must be valid")
	}

	// member should be add to designated peer
	meta := peerMeta
	t.pm.AddDesignatedPeer(meta)
	t.pm.AddNewPeer(meta)
}

func (t *MeyCoinRaftTransport) RemovePeer(id rtypes.ID) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	st := t.statuses[id]
	if st == nil {
		return
	}
	t.pm.RemoveDesignatedPeer(st.pid)
	delete(t.statuses, id)
	delete(t.stByPID, st.pid)
	if peer, exist := t.pm.GetPeer(st.pid); exist {
		peer.Stop()
	}
	t.logger.Info().Str(p2putil.LogPeerID, p2putil.ShortForm(st.pid)).Uint64("raftID", uint64(id)).Msg("removed raft peer")
}

func (t *MeyCoinRaftTransport) RemoveAllPeers() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for _, peer := range t.pm.GetPeers() {
		peer.Stop()
	}
}

func (t *MeyCoinRaftTransport) UpdatePeer(id rtypes.ID, urls []string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	// To Nothing for now
}

func (t *MeyCoinRaftTransport) ActiveSince(id rtypes.ID) time.Time {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if p, ok := t.statuses[id]; ok {
		return p.activeSince()
	}
	return time.Time{}

}

func (t *MeyCoinRaftTransport) ActivePeers() int {
	return len(t.pm.GetPeers())
}

func (t *MeyCoinRaftTransport) Stop() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	// Lots of works will be done by p2p modules
	t.nt.RemoveStreamHandler(p2pcommon.RaftSnapSubAddr)
}

func (t *MeyCoinRaftTransport) OnRaftSnapshot(s network.Stream) {
	hsresp := p2pcommon.HSHeadResp{Magic: p2pcommon.MAGICRaftSnap}
	peerID := s.Conn().RemotePeer()

	// check validation
	// check if sender is the leader node
	peer, found := t.pm.GetPeer(peerID)
	if !found {
		addr := s.Conn().RemoteMultiaddr()
		t.logger.Info().Str(p2putil.LogPeerID, p2putil.ShortForm(peerID)).Str("multiaddr", addr.String()).Msg("snapshot stream from leader node")
		hsresp.RespCode = p2pcommon.HSCodeAuthFail
		s.Write(hsresp.Marshal())
		s.Close()
		return
	}
	//// TODO raft role is not properly set yet.
	//if peer.AcceptedRole() != p2pcommon.RaftLeader {
	//	t.logger.Warn().Str(p2putil.LogPeerName, peer.Name()).Msg("Closing snapshot stream from follower node")
	//	hsresp.RespCode = p2pcommon.HSCodeNoPermission
	//	s.Write(hsresp.Marshal())
	//	s.Close()
	//	return
	//}
	//
	t.logger.Debug().Str(p2putil.LogPeerName, peerID.Pretty()).Msg("snapshot stream from leader node")

	// read stream and send it to raft
	sr := t.snapF.NewSnapshotReceiver(peer, s)
	sr.Receive()
	t.logger.Debug().Str(p2putil.LogPeerName, peerID.Pretty()).Msg("snapshot receiving finished")
	s.Close()
}

func (t *MeyCoinRaftTransport) OnPeerConnect(pid types.PeerID) {
	go func() {
		t.mutex.Lock()
		defer t.mutex.Unlock()
		if st, exist := t.stByPID[pid]; exist {
			st.activate()
		}
	}()
}

func (t *MeyCoinRaftTransport) OnPeerDisconnect(peer p2pcommon.RemotePeer) {
	go func(peerID types.PeerID) {
		t.mutex.Lock()
		defer t.mutex.Unlock()
		if st, exist := t.stByPID[peer.ID()]; exist {
			st.deactivate("disconnect")
		}
	}(peer.ID())
}

func (t *MeyCoinRaftTransport) NewSnapshotSender(peer p2pcommon.RemotePeer) SnapshotSender {
	return newSnapshotSender(t.logger, t.nt, t.raftAcc, peer)
}

func (t *MeyCoinRaftTransport) NewSnapshotReceiver(peer p2pcommon.RemotePeer, rwc io.ReadWriteCloser) SnapshotReceiver {
	return newSnapshotReceiver(t.logger, t.pm, t.raftAcc, peer, rwc)
}
