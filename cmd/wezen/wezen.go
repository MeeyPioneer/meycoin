/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */
package main

import (
	"fmt"
	"github.com/meeypioneer/mey-actor/actor"
	"github.com/meeypioneer/mey-library/log"
	"github.com/meeypioneer/meycoin/config"
	"github.com/meeypioneer/meycoin/internal/common"
	"github.com/meeypioneer/meycoin/p2p/p2pkey"
	"github.com/meeypioneer/meycoin/pkg/component"
	common2 "github.com/meeypioneer/meycoin/wezen/common"
	"github.com/meeypioneer/meycoin/wezen/server"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "wezen",
		Short: "Wezen node discovery service",
		Long:  "Wezen node discovery service for providing peer addresses to connect.",
		Run:   rootRun,
	}
	homePath       string
	configFilePath string
	enableTestmode bool
	svrlog         *log.Logger

	cfg *config.Config
)

func init() {
	cobra.OnInitialize(initConfig)
	fs := rootCmd.PersistentFlags()
	fs.StringVar(&homePath, "home", "", "path of meycoin home")
	fs.StringVar(&configFilePath, "config", "", "path of configuration file")
	fs.BoolVar(&enableTestmode, "testmode", false, "enable unsafe test mode (skips certain validations)")
}

func initConfig() {
	serverCtx := config.NewServerContext(homePath, configFilePath)
	cfg = serverCtx.GetDefaultConfig().(*config.Config)
	// change some different default props for wezen
	arrangeDefaultCfgForWezen(cfg)
	err := serverCtx.LoadOrCreateConfig(cfg)
	if err != nil {
		fmt.Printf("Fail to load configuration file %v: %v", serverCtx.Vc.ConfigFileUsed(), err.Error())
		os.Exit(1)
	}
	if enableTestmode {
		cfg.EnableTestmode = true
	}
	cfg.Consensus.EnableBp = false
}

func arrangeDefaultCfgForWezen(cfg *config.Config) {
	cfg.RPC.NetServicePort = common2.DefaultRPCPort
	cfg.P2P.NetProtocolPort = common2.DefaultSrvPort
}

func rootRun(cmd *cobra.Command, args []string) {

	svrlog = log.NewLogger("wezen")
	svrlog.Info().Msg("WEZEN STARTED")

	p2pkey.InitNodeInfo(&cfg.BaseConfig, cfg.P2P, githash, svrlog)

	if cfg.EnableProfile {
		svrlog.Info().Msgf("Enable Profiling on localhost: %d", cfg.ProfilePort)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.ProfilePort), nil)
			svrlog.Info().Err(err).Msg("Run Profile Server")
		}()
	}

	if cfg.EnableTestmode {
		svrlog.Warn().Msgf("Running with unsafe test mode. Turn off test mode for production use!")
	}

	p2pkey.InitNodeInfo(&cfg.BaseConfig, cfg.P2P, "TODO", svrlog)

	compMng := component.NewComponentHub()

	lNTC := server.NewNTContainer(cfg)
	wezenSvc := server.NewWezenService(cfg, lNTC)
	rpcSvc := server.NewWezenRPC(cfg)

	// Register services to Hub. Don't need to do nil-check since Register
	// function skips nil parameters.
	compMng.Register(lNTC, wezenSvc, rpcSvc)

	//consensusSvc, err := impl.New(cfg.Consensus, compMng, chainSvc)
	//if err != nil {
	//	svrlog.Error().Err(err).Msg("Failed to start consensus service.")
	//	os.Exit(1)
	//}

	// All the services objects including Consensus must be created before the
	// actors are started.
	compMng.Start()

	var interrupt = common.HandleKillSig(func() {
		//consensus.Stop(consensusSvc)
		compMng.Stop()
	}, svrlog)

	// Wait main routine to stop
	<-interrupt.C
}

type RedirectService struct {
	*component.BaseComponent
}

func NewRedirectService(cfg *config.Config, svcPid string) *RedirectService {
	logger := log.NewLogger(svcPid)
	rs := &RedirectService{}
	rs.BaseComponent = component.NewBaseComponent(svcPid, rs, logger)

	return rs
}

func (rs *RedirectService) Receive(context actor.Context) {
	// ignore for now
}

func (rs *RedirectService) BeforeStart() {}
func (rs *RedirectService) AfterStart()  {}
func (rs *RedirectService) BeforeStop()  {}

func (rs *RedirectService) Statistics() *map[string]interface{} {
	dummy := make(map[string]interface{})
	return &dummy
}
