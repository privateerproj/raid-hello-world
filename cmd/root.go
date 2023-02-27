package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/privateerproj/privateer-pack-wireframe/pack"
	"github.com/privateerproj/privateer-sdk/command"
	"github.com/privateerproj/privateer-sdk/plugin"
	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	buildVersion       string
	buildGitCommitHash string
	buildTime          string

	// RaidName is the name of this raid
	RaidName = "Wireframe"

	// runCmd represents the base command when called without any subcommands
	runCmd = &cobra.Command{
		Use:   RaidName,
		Short: "A brief description of your application",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			command.InitializeConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) > 1 && os.Args[1] == "debug" {
				raidengine.Run("Wireframe", pack.Policies)
			} else {
				// Serve plugin
				raid := &Raid{}
				serveOpts := &plugin.ServeOpts{
					Plugin: raid,
				}

				plugin.Serve(serveOpts)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the runCmd.
func Execute(version, commitHash, builtAt string) {
	buildVersion = version
	buildGitCommitHash = commitHash
	buildTime = builtAt

	err := runCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	command.SetBase(runCmd)
}

// Raid meets the Privateer Service Pack interface
type Raid struct {
	RaidName string
}

// Start is called from Privateer after the plugin is served
func (sp *Raid) Start() error {
	// raidengine.SetupCloseHandler(sigtermProtection)
	return raidengine.Run(RaidName, pack.Policies) // TODO handle these errors
}
