/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import (
	"context"
	"github.com/meeypioneer/meycoin/cmd/meycoincli/util"
	"time"

	"github.com/mr-tron/base58/base58"

	"github.com/meeypioneer/meycoin/types"
	"github.com/spf13/cobra"
)

var CurrentPeersCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current peers list",
	Run:   execCurrentPeers,
}
var cpKey string
var cpSize int

func init() {
	rootCmd.AddCommand(CurrentPeersCmd)

	CurrentPeersCmd.Flags().StringVar(&cpKey, "ref", "", "Reference Key")
	CurrentPeersCmd.Flags().IntVar(&cpSize, "size", 20, "Max list size")

}

func execCurrentPeers(cmd *cobra.Command, args []string) {
	var blockHash []byte
	var err error

	if cmd.Flags().Changed("ref") == true {
		blockHash, err = base58.Decode(cpKey)
		if err != nil {
			cmd.Printf("Failed: %s", err.Error())
			return
		}
	}
	if cpSize <= 0 {
		cpSize = 20
	}

	uparams := &types.Paginations{
		Ref:  blockHash,
		Size: uint32(cpSize),
	}

	msg, err := client.CurrentList(context.Background(), uparams)
	if err != nil {
		cmd.Printf("Failed: %s", err.Error())
		return
	}
	// TODO decorate other props also;e.g. uint64 timestamp to human readable time format!
	ppList := make([]JSONWezenPeer, len(msg.Peers))
	for i, p := range msg.Peers {
		ppList[i] = NewJSONWezenPeer(p)
	}
	cmd.Println(util.B58JSON(ppList))
}

type WezenPeerAlias types.WezenPeer

// JSONWezenPeer is simple wrapper to print human readable json output.
type JSONWezenPeer struct {
	*WezenPeerAlias
	Connected time.Time `json:"connected"`
	LastCheck time.Time `json:"lastCheck"`
}

func NewJSONWezenPeer(pp *types.WezenPeer) JSONWezenPeer {
	return JSONWezenPeer{
		(*WezenPeerAlias)(pp),
		time.Unix(0, pp.Connected),
		time.Unix(0, pp.LastCheck),
	}
}