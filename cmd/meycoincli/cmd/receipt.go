/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import (
	"context"
	"log"

	"github.com/meeypioneer/meycoin/cmd/meycoincli/util"
	meycoinrpc "github.com/meeypioneer/meycoin/types"
	"github.com/mr-tron/base58/base58"
	"github.com/spf13/cobra"
)

func init() {
	receiptCmd := &cobra.Command{
		Use:   "receipt [flags] subcommand",
		Short: "Receipt command",
	}
	rootCmd.AddCommand(receiptCmd)

	receiptCmd.AddCommand(
		&cobra.Command{
			Use:   "get [flags] tx_hash",
			Short: "Get a receipt",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				txHash, err := base58.Decode(args[0])
				if err != nil {
					log.Fatal(err)
				}
				msg, err := client.GetReceipt(context.Background(), &meycoinrpc.SingleBytes{Value: txHash})
				if err != nil {
					log.Fatal(err)
				}
				cmd.Println(util.JSON(msg))
			},
		},
	)
}
