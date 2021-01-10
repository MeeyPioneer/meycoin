/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package common

import "github.com/meeypioneer/meycoin/types"

var (
	// 89.16 is ceiling of declination of Wezen
	MainnetMapServer = []string{
		"/dns4/mainnet-wezen.meycoin.io/tcp/8916/p2p/16Uiu2HAkuxyDkMTQTGFpmnex2SdfTVzYfPztTyK339rqUdsv3ZUa",
	}

	// 89.16 is ceiling of declination of Wezen
	TestnetMapServer = []string{
		"/dns4/wezen.meycoin.io/tcp/8916/p2p/16Uiu2HAkvJTHFuJXxr15rFEHsJWnyn1QvGatW2E9ED9Mvy4HWjVF",
	}

	// Hardcoded chainID of ONE MAINNET and ONE TESTNET
	ONEMainNet types.ChainID
	ONETestNet types.ChainID
)

func init() {
	mnGen := types.GetMainNetGenesis()
	if mnGen == nil {
		panic("Failed to get MainNet GenesisInfo")
	}
	ONEMainNet = mnGen.ID

	tnGen := types.GetTestNetGenesis()
	if tnGen == nil {
		panic("Failed to get TestNet GenesisInfo")
	}
	ONETestNet = tnGen.ID
}

