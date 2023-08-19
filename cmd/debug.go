package cmd

import (
	"github.com/spf13/cobra"

	"github.com/privateerproj/privateer-pack-wireframe/pack"
	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	// debugCmd represents the base command when called without any subcommands
	debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "A brief description of your application",
		Run: func(cmd *cobra.Command, args []string) {
			raidengine.Run(pack.Strikes)
		},
	}
)

func init() {
	runCmd.AddCommand(debugCmd)
}
