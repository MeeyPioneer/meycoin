/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package client

import (
	"errors"
	"github.com/meeypioneer/meycoin/p2p/p2putil"
	"github.com/meeypioneer/meycoin/wezen/common"
	"testing"

	"github.com/meeypioneer/meycoin/config"
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"github.com/meeypioneer/meycoin/p2p/p2pmock"
	"github.com/meeypioneer/meycoin/pkg/component"
	"github.com/meeypioneer/meycoin/types"
	"github.com/golang/mock/gomock"
)


type dummyNTC struct {
	nt      p2pcommon.NetworkTransport
	chainID *types.ChainID
	selfMeta p2pcommon.PeerMeta
}

func (dntc *dummyNTC) SelfMeta() p2pcommon.PeerMeta {
	return dntc.selfMeta
}

func (dntc *dummyNTC) GetNetworkTransport() p2pcommon.NetworkTransport {
	return dntc.nt
}
func (dntc *dummyNTC) GenesisChainID() *types.ChainID {
	return dntc.chainID
}

var (
	pmapDummyCfg = &config.Config{P2P: &config.P2PConfig{}, Wezen: &config.WezenConfig{GenesisFile: "../../examples/genesis.json"}}
	pmapDummyNTC = &dummyNTC{chainID: &types.ChainID{}}
)

// initSvc select Wezenes to connect, or disable wezen
func TestWezenConnectSvc_initSvc(t *testing.T) {
	wezenIDMain, _ := types.IDB58Decode("16Uiu2HAkuxyDkMTQTGFpmnex2SdfTVzYfPztTyK339rqUdsv3ZUa")
	wezenIDTest, _ := types.IDB58Decode("16Uiu2HAkvJTHFuJXxr15rFEHsJWnyn1QvGatW2E9ED9Mvy4HWjVF")
	dummyPeerID2, _ := types.IDB58Decode("16Uiu2HAmFqptXPfcdaCdwipB2fhHATgKGVFVPehDAPZsDKSU7jRm")
	polar2 := "/ip4/172.21.1.2/tcp/8915/p2p/16Uiu2HAmFqptXPfcdaCdwipB2fhHATgKGVFVPehDAPZsDKSU7jRm"
	dummyPeerID3, _ := types.IDB58Decode("16Uiu2HAmU8Wc925gZ5QokM4sGDKjysdPwRCQFoYobvoVnyutccCD")
	polar3 := "/ip4/172.22.2.3/tcp/8915/p2p/16Uiu2HAmU8Wc925gZ5QokM4sGDKjysdPwRCQFoYobvoVnyutccCD"

	customChainID := types.ChainID{Magic: "unittest.blocko.io"}
	type args struct {
		use       bool
		wezenes []string

		chainID *types.ChainID
	}
	tests := []struct {
		name string
		args args

		wantCnt int
		peerIDs []types.PeerID
	}{
		//
		{"TMeyCoinNoWezen", args{false, nil, &common.ONEMainNet}, 0, []types.PeerID{}},
		{"TMeyCoinMainDefault", args{true, nil, &common.ONEMainNet}, 1, []types.PeerID{wezenIDMain}},
		{"TMeyCoinMainPlusCfg", args{true, []string{polar2, polar3}, &common.ONEMainNet}, 3, []types.PeerID{wezenIDMain, dummyPeerID2, dummyPeerID3}},
		{"TMeyCoinTestDefault", args{true, nil, &common.ONETestNet}, 1, []types.PeerID{wezenIDTest}},
		{"TMeyCoinTestPlusCfg", args{true, []string{polar2, polar3}, &common.ONETestNet}, 3, []types.PeerID{wezenIDTest, dummyPeerID2, dummyPeerID3}},
		{"TCustom", args{true, nil, &customChainID}, 0, []types.PeerID{}},
		{"TCustomPlusCfg", args{true, []string{polar2, polar3}, &customChainID}, 2, []types.PeerID{dummyPeerID2, dummyPeerID3}},
		{"TWrongWezenAddr", args{true, []string{"/ip4/256.256.1.1/tcp/8915/p2p/16Uiu2HAmU8Wc925gZ5QokM4sGDKjysdPwRCQFoYobvoVnyutccCD"}, &customChainID}, 0, []types.PeerID{}},
		{"TWrongWezenAddr2", args{true, []string{"/egwgew5/p2p/16Uiu2HAmU8Wc925gZ5QokM4sGDKjysdPwRCQFoYobvoVnyutccCD"}, &customChainID}, 0, []types.PeerID{}},
		{"TWrongWezenAddr3", args{true, []string{"/dns/nowhere1234.io/tcp/8915/p2p/16Uiu2HAmU8Wc925gZ5QokM4sGDKjysdPwRCQFoYobvoVnyutccCD"}, &customChainID}, 1, []types.PeerID{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockNT := p2pmock.NewMockNetworkTransport(ctrl)
			pmapDummyNTC.nt = mockNT
			pmapDummyNTC.chainID = tt.args.chainID

			cfg := config.NewServerContext("", "").GetDefaultP2PConfig()
			cfg.NPUseWezen = tt.args.use
			cfg.NPAddWezenes = tt.args.wezenes

			pcs := NewWezenConnectSvc(cfg, pmapDummyNTC)

			if len(pcs.mapServers) != tt.wantCnt {
				t.Errorf("NewWezenConnectSvc() = %v, want %v", len(pcs.mapServers), tt.wantCnt)
			}
			for _, wantPeerID := range tt.peerIDs {
				found := false
				for _, wezenMeta := range pcs.mapServers {
					if wantPeerID == wezenMeta.ID {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("initSvc() want exist %v but not ", wantPeerID)
				}
			}
		})
	}
}

func TestWezenConnectSvc_BeforeStop(t *testing.T) {

	type fields struct {
		BaseComponent *component.BaseComponent
	}
	tests := []struct {
		name   string
		fields fields

		calledStreamHandler bool
	}{
		{"TNot", fields{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockNT := p2pmock.NewMockNetworkTransport(ctrl)
			pmapDummyNTC.nt = mockNT
			pms := NewWezenConnectSvc(pmapDummyCfg.P2P, pmapDummyNTC)

			mockNT.EXPECT().AddStreamHandler(common.WezenPingSub, gomock.Any()).Times(1)
			mockNT.EXPECT().RemoveStreamHandler(common.WezenPingSub).Times(1)

			pms.AfterStart()

			pms.BeforeStop()

			ctrl.Finish()
		})
	}
}

func TestWezenConnectSvc_queryToWezen(t *testing.T) {
	sampleCahinID := &types.ChainID{Consensus:"dpos"}
	sm := p2pcommon.PeerMeta{ID:types.RandomPeerID()}
	ss := &types.Status{}

	sErr := errors.New("send erroor")
	rErr := errors.New("send erroor")

	succR  := &types.MapResponse{Status:types.ResultStatus_OK, Addresses:[]*types.PeerAddress{&types.PeerAddress{}}}
	oldVerR  := &types.MapResponse{Status:types.ResultStatus_FAILED_PRECONDITION, Message:common.TooOldVersionMsg}
	otherR  := &types.MapResponse{Status:types.ResultStatus_INVALID_ARGUMENT,Message:"arg is wrong"}

	type args struct {
		mapServerMeta p2pcommon.PeerMeta
		peerStatus *types.Status
	}
	tests := []struct {
		name    string
		args    args

		sendErr error
		readErr error
		mapResp *types.MapResponse

		wantCnt int
		wantErr bool
	}{
		{"TSingle", args{sm, ss}, nil, nil, succR, 1, false},
		{"TSendFail", args{sm, ss}, sErr, nil, succR,0, true},
		{"TRecvFail", args{sm, ss}, nil, rErr, succR,0, true},
		{"TOldVersion", args{sm, ss}, nil, nil, oldVerR, 0, true},
		{"TFailResp", args{sm, ss}, nil, nil, otherR, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cfg := config.NewServerContext("", "").GetDefaultP2PConfig()
			cfg.NPUseWezen = true

			mockNTC := p2pmock.NewMockNTContainer(ctrl)
			mockNTC.EXPECT().GenesisChainID().Return(sampleCahinID).AnyTimes()

			mockRW := p2pmock.NewMockMsgReadWriter(ctrl)

			payload, _ := p2putil.MarshalMessageBody(tt.mapResp)
			dummyData := common.NewWezenMessage(p2pcommon.NewMsgID(), common.MapResponse, payload)

			mockRW.EXPECT().WriteMsg(gomock.AssignableToTypeOf(&common.WezenMessage{})).Return(tt.sendErr)
			if tt.sendErr == nil {
				mockRW.EXPECT().ReadMsg().Return(dummyData, tt.readErr)
			}
			pcs := NewWezenConnectSvc(cfg, mockNTC)


			got, err := pcs.queryToWezen(tt.args.mapServerMeta, mockRW, tt.args.peerStatus)
			if (err != nil) != tt.wantErr {
				t.Errorf("connectAndQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.wantCnt {
				t.Errorf("connectAndQuery() len(got) = %v, want %v", len(got), tt.wantCnt)
			}
		})
	}
}