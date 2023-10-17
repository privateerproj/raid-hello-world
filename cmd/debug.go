package cmd

import (
	"github.com/spf13/cobra"

	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	// debugCmd represents the base command when called without any subcommands
	debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "Run the Raid in debug mode",
		Run: func(cmd *cobra.Command, args []string) {
			raidengine.Run(RaidName, getStrikes())
		},
	}
)

func init() {
	runCmd.AddCommand(debugCmd) // This enables the debug command for use while working on your Raid
}
