package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/privateerproj/privateer-sdk/command"
	"github.com/privateerproj/privateer-sdk/plugin"
	"github.com/privateerproj/privateer-sdk/raidengine"

	"github.com/privateerproj/privateer-pack-wireframe/strikes"
)

var (
	// Build information is added by the Makefile at compile time
	buildVersion       string
	buildGitCommitHash string
	buildTime          string

	RaidName = "Wireframe" // TODO: Change this to the name of your Raid

	// runCmd represents the base command when called without any subcommands
	runCmd = &cobra.Command{
		Use:   RaidName,
		Short: "TODO: Addd a brief description of your Raid here",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			command.InitializeConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Serve plugin
			raid := &Raid{}
			serveOpts := &plugin.ServeOpts{
				Plugin: raid,
			}

			plugin.Serve(RaidName, serveOpts)
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
	command.SetBase(runCmd) // This initializes the base CLI functionality
}

// Raid meets the Privateer Service Pack interface
type Raid struct {
}

// cleanupFunc is called when the plugin is stopped
func cleanupFunc() error {
	return nil
}

// Start is called from Privateer after the plugin is served
// At minimum, this should call raidengine.Run()
// Adding raidengine.SetupCloseHandler(cleanupFunc) will allow you to append custom cleanup behavior
func (r *Raid) Start() error {
	raidengine.SetupCloseHandler(cleanupFunc)
	return raidengine.Run(RaidName, getStrikes()) // Return errors from strike executions
}

// GetStrikes returns a list of probe objects
func getStrikes() map[string][]raidengine.Strike {
	logger := raidengine.GetLogger(RaidName, false)
	a := &strikes.Antijokes{
		Log: logger,
	}
	return map[string][]raidengine.Strike{
		"PCI": {
			a.KnockKnock,
		},
		"CIS": {
			a.KnockKnock,
			a.ChickenCrossedRoad,
		},
	}
}
