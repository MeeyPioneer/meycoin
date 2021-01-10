/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package client

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/meeypioneer/meycoin/p2p/p2pkey"
	"github.com/meeypioneer/meycoin/p2p/v030"
	"github.com/meeypioneer/meycoin/wezen/common"
	"github.com/libp2p/go-libp2p-core/network"
	"sync"

	"github.com/meeypioneer/mey-actor/actor"
	"github.com/meeypioneer/mey-library/log"
	"github.com/meeypioneer/meycoin/config"
	"github.com/meeypioneer/meycoin/message"
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"github.com/meeypioneer/meycoin/p2p/p2putil"
	"github.com/meeypioneer/meycoin/pkg/component"
	"github.com/meeypioneer/meycoin/types"
)

var ErrTooLowVersion = errors.New("meycoinsvr version is too low")

// PeerMapService is
type WezenConnectSvc struct {
	*component.BaseComponent

	PrivateChain bool

	mapServers []p2pcommon.PeerMeta
	exposeself bool

	ntc p2pcommon.NTContainer
	nt  p2pcommon.NetworkTransport

	rwmutex *sync.RWMutex
}

func NewWezenConnectSvc(cfg *config.P2PConfig, ntc p2pcommon.NTContainer) *WezenConnectSvc {
	pcs := &WezenConnectSvc{
		rwmutex:    &sync.RWMutex{},
		exposeself: cfg.NPExposeSelf,
	}
	pcs.BaseComponent = component.NewBaseComponent(message.MapSvc, pcs, log.NewLogger("pcs"))

	// init
	pcs.ntc = ntc
	pcs.initSvc(cfg)

	return pcs
}

func (pcs *WezenConnectSvc) initSvc(cfg *config.P2PConfig) {
	pcs.PrivateChain = !pcs.ntc.GenesisChainID().PublicNet
	if cfg.NPUseWezen {
		// private network does not use public wezen
		if !pcs.PrivateChain {
			servers := make([]string, 0)
			// add hardcoded built-in servers if net is ONE net.
			if *pcs.ntc.GenesisChainID() == common.ONEMainNet {
				pcs.Logger.Info().Msg("chain is ONE Mainnet so use default wezen for mainnet")
				servers = common.MainnetMapServer
			} else if *pcs.ntc.GenesisChainID() == common.ONETestNet {
				pcs.Logger.Info().Msg("chain is ONE Testnet so use default wezen for testnet")
				servers = common.TestnetMapServer
			} else {
				pcs.Logger.Info().Msg("chain is custom public network so only custom wezen in configuration file will be used")
			}

			for _, addrStr := range servers {
				meta, err := p2putil.FromMultiAddrString(addrStr)
				if err != nil {
					pcs.Logger.Info().Str("addr_str", addrStr).Msg("invalid wezen server address in base setting ")
					continue
				}
				pcs.mapServers = append(pcs.mapServers, meta)
			}
		} else {
			pcs.Logger.Info().Msg("chain is private so only using wezen in config file")
		}
		// append custom wezenes set in configuration file
		for _, addrStr := range cfg.NPAddWezenes {
			meta, err := p2putil.FromMultiAddrString(addrStr)
			if err != nil {
				pcs.Logger.Info().Str("addr_str", addrStr).Msg("invalid wezen server address in config file ")
				continue
			}
			pcs.mapServers = append(pcs.mapServers, meta)
		}

		if len(pcs.mapServers) == 0 {
			pcs.Logger.Warn().Msg("using Wezen is enabled but no active wezen server found. node discovery by wezen will not works well")
		} else {
			pcs.Logger.Info().Array("wezenes", p2putil.NewLogPeerMetasMarshaller(pcs.mapServers, 10)).Msg("using Wezen")
		}
	} else {
		pcs.Logger.Info().Msg("node discovery by Wezen is disabled by configuration.")
	}
}

func (pcs *WezenConnectSvc) BeforeStart() {}

func (pcs *WezenConnectSvc) AfterStart() {
	pcs.nt = pcs.ntc.GetNetworkTransport()
	pcs.nt.AddStreamHandler(common.WezenPingSub, pcs.onPing)

}

func (pcs *WezenConnectSvc) BeforeStop() {
	pcs.nt.RemoveStreamHandler(common.WezenPingSub)
}

func (pcs *WezenConnectSvc) Statistics() *map[string]interface{} {
	return nil
	//dummy := make(map[string]interface{})
	//return &dummy
}

func (pcs *WezenConnectSvc) Receive(context actor.Context) {
	rawMsg := context.Message()
	switch msg := rawMsg.(type) {
	case *message.MapQueryMsg:
		pcs.Hub().Tell(message.P2PSvc, pcs.queryPeers(msg))
	default:
		//		pcs.Logger.Debug().Interface("msg", msg)
	}
}

func (pcs *WezenConnectSvc) queryPeers(msg *message.MapQueryMsg) *message.MapQueryRsp {
	succ := 0
	resultPeers := make([]*types.PeerAddress, 0, msg.Count)
	for _, meta := range pcs.mapServers {
		addrs, err := pcs.connectAndQuery(meta, msg.BestBlock.Hash, msg.BestBlock.Header.BlockNo)
		if err != nil {
			if err == ErrTooLowVersion {
				pcs.Logger.Error().Err(err).Str("wezenID", p2putil.ShortForm(meta.ID)).Msg("Wezen responded this meycoinsvr is too low, check and upgrade meycoinsvr")
			} else {
				pcs.Logger.Warn().Err(err).Str("wezenID", p2putil.ShortForm(meta.ID)).Msg("failed to get peer addresses")
			}
			continue
		}
		// duplicated peers will be filtered out by caller (more precisely, p2p actor)
		resultPeers = append(resultPeers, addrs...)
		succ++
	}
	err := error(nil)
	if succ == 0 {
		err = fmt.Errorf("all servers of wezen are down")
	}
	pcs.Logger.Debug().Int("peer_cnt", len(resultPeers)).Msg("Got map response and send back")
	resp := &message.MapQueryRsp{Peers: resultPeers, Err: err}
	return resp
}

func (pcs *WezenConnectSvc) connectAndQuery(mapServerMeta p2pcommon.PeerMeta, bestHash []byte, bestHeight uint64) ([]*types.PeerAddress, error) {
	s, err := pcs.nt.GetOrCreateStreamWithTTL(mapServerMeta, common.WezenConnectionTTL, common.WezenMapSub)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	peerID := s.Conn().RemotePeer()
	if peerID != mapServerMeta.ID {
		return nil, fmt.Errorf("internal error peerid mismatch, exp %s, actual %s", mapServerMeta.ID.Pretty(), peerID.Pretty())
	}
	pcs.Logger.Debug().Str("wezenID", peerID.String()).Msg("Sending map query")

	rw := v030.NewV030ReadWriter(bufio.NewReader(s), bufio.NewWriter(s), nil)

	peerAddress := pcs.ntc.SelfMeta().ToPeerAddress()
	chainBytes, _ := pcs.ntc.GenesisChainID().Bytes()
	peerStatus := &types.Status{Sender: &peerAddress, BestBlockHash: bestHash, BestHeight: bestHeight, ChainID: chainBytes,
		Version: p2pkey.NodeVersion()}

	return pcs.queryToWezen(mapServerMeta, rw, peerStatus)
}

func (pcs *WezenConnectSvc) queryToWezen(mapServerMeta p2pcommon.PeerMeta, rw p2pcommon.MsgReadWriter, peerStatus *types.Status) ([]*types.PeerAddress, error) {
	// receive input
	err := pcs.sendRequest(peerStatus, mapServerMeta, pcs.exposeself, 100, rw)
	if err != nil {
		return nil, err
	}
	_, resp, err := pcs.readResponse(mapServerMeta, rw)
	if err != nil {
		return nil, err
	}
	switch resp.Status {
	case types.ResultStatus_OK:
		return resp.Addresses, nil
	case types.ResultStatus_FAILED_PRECONDITION:
		if resp.Message == common.TooOldVersionMsg {
			return nil, ErrTooLowVersion
		} else {
			return nil, fmt.Errorf("remote error %s", resp.Status.String())
		}
	default:
		return nil, fmt.Errorf("remote error %s", resp.Status.String())
	}
}

func (pcs *WezenConnectSvc) sendRequest(status *types.Status, mapServerMeta p2pcommon.PeerMeta, register bool, size int, wt p2pcommon.MsgReadWriter) error {
	msgID := p2pcommon.NewMsgID()
	queryReq := &types.MapQuery{Status: status, Size: int32(size), AddMe: register, Excludes: [][]byte{[]byte(mapServerMeta.ID)}}
	bytes, err := p2putil.MarshalMessageBody(queryReq)
	if err != nil {
		return err
	}
	reqMsg := common.NewWezenMessage(msgID, common.MapQuery, bytes)

	return wt.WriteMsg(reqMsg)
}

// tryAddPeer will do check connecting peer and add. it will return peer meta information received from
// remote peer setup some
func (pcs *WezenConnectSvc) readResponse(mapServerMeta p2pcommon.PeerMeta, rd p2pcommon.MsgReadWriter) (p2pcommon.Message, *types.MapResponse, error) {
	data, err := rd.ReadMsg()
	if err != nil {
		return nil, nil, err
	}
	queryResp := &types.MapResponse{}
	err = p2putil.UnmarshalMessageBody(data.Payload(), queryResp)
	if err != nil {
		return data, nil, err
	}
	// old version of wezen will return old formatted PeerAddress, so conversion to new format is needed.
	for _, addr := range queryResp.Addresses {
		if len(addr.Addresses) == 0 {
			ma, err := types.ToMultiAddr(addr.Address, addr.Port)
			if err != nil {
				continue
			}
			addr.Addresses=[]string{ma.String()}
		}
	}
	pcs.Logger.Debug().Str(p2putil.LogPeerID, mapServerMeta.ID.String()).Int("peer_cnt", len(queryResp.Addresses)).Msg("Received map query response")

	return data, queryResp, nil
}

func (pcs *WezenConnectSvc) onPing(s network.Stream) {
	peerID := s.Conn().RemotePeer()
	pcs.Logger.Debug().Str("wezenID", peerID.String()).Msg("Received ping from wezen (maybe)")

	rw := v030.NewV030ReadWriter(bufio.NewReader(s), bufio.NewWriter(s), nil)
	defer s.Close()

	req, err := rw.ReadMsg()
	if err != nil {
		return
	}
	pingReq := &types.Ping{}
	err = p2putil.UnmarshalMessageBody(req.Payload(), pingReq)
	if err != nil {
		return
	}
	// TODO: check if sender is known wezen or peer and it not, ban or write to blacklist .
	pingResp := &types.Ping{}
	bytes, err := p2putil.MarshalMessageBody(pingResp)
	if err != nil {
		return
	}
	msgID := p2pcommon.NewMsgID()
	respMsg := common.NewWezenRespMessage(msgID, req.ID(), p2pcommon.PingResponse, bytes)
	err = rw.WriteMsg(respMsg)
	if err != nil {
		return
	}
}
