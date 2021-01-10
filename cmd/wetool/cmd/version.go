/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import "github.com/spf13/cobra"

var githash = "No git hash provided"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Wetool",
	Long:  `The version of Wetool is following the one of Wezen`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Wetool %s\n", githash)
	},
}
