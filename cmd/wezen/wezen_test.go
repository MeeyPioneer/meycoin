package main

import (
	"github.com/meeypioneer/meycoin/wezen/common"
	_ "net/http/pprof"
	"testing"

	"github.com/meeypioneer/meycoin/config"
)

func Test_arrangeDefaultCfgForWezen(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TNormal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg = config.NewServerContext("/","/").GetDefaultConfig().(*config.Config)
			if cfg.RPC.NetServicePort != 7845 {
				t.Errorf("Assumption failure: default cfg.RPC.NetServicePort = %d, want %d",cfg.RPC.NetServicePort, 7845)
			}
			if cfg.P2P.NetProtocolPort != 7846 {
				t.Errorf("Assumption failure: default cfg.P2P.NetProtocolPort = %d, want %d",cfg.P2P.NetProtocolPort, 7846)
			}
			arrangeDefaultCfgForWezen(cfg)
			if cfg.RPC.NetServicePort != common.DefaultRPCPort {
				t.Errorf("cfg.RPC.NetServicePort = %d, want %d",cfg.RPC.NetServicePort, common.DefaultRPCPort)
			}
			if cfg.P2P.NetProtocolPort != common.DefaultSrvPort {
				t.Errorf("cfg.P2P.NetProtocolPort = %d, want %d",cfg.P2P.NetProtocolPort, common.DefaultSrvPort)
			}
		})
	}
}
