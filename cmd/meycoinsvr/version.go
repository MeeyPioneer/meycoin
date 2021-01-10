package main

import "github.com/spf13/cobra"

var githash = "No git hash provided"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of MeyCoinsvr",
	Long:  `All software has versions. This is MeyCoin's`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("MeyCoinsvr %s\n", githash)
	},
}
