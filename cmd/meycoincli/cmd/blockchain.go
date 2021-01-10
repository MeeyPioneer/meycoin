/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import (
	"context"

	"github.com/meeypioneer/meycoin/cmd/meycoincli/util"
	meycoinrpc "github.com/meeypioneer/meycoin/types"
	"github.com/spf13/cobra"
)

var printHex bool

func init() {
	rootCmd.AddCommand(blockchainCmd)
	blockchainCmd.Flags().BoolVar(&printHex, "hex", false, "Print bytes to hex format")
}

var blockchainCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "Print current blockchain status",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := client.Blockchain(context.Background(), &meycoinrpc.Empty{})
		if err != nil {
			cmd.Printf("Failed: %s\n", err.Error())
			return
		}
		if printHex {
			cmd.Println(util.ConvHexBlockchainStatus(msg))
		} else {
			cmd.Println(util.ConvBlockchainStatus(msg))
		}
	},
}
