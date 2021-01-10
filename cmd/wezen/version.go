/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package main

import "github.com/spf13/cobra"

var githash = "No git hash provided"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Wezen",
	Long:  `The version of Wezen is following the one of Arego`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Wezen %s\n", githash)
	},
}
