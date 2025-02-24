package main

import (
	"fmt"

	"os"

	"github.com/privateerproj/privateer-plugin-example-plugin/armory"

	"github.com/privateerproj/privateer-sdk/command"
	"github.com/privateerproj/privateer-sdk/config"
)

var (
	// Version is to be replaced at build time by the associated tag
	Version = "0.0.0"
	// VersionPostfix is a marker for the version such as "dev", "beta", "rc", etc.
	VersionPostfix = "dev"
	// GitCommitHash is the commit at build time
	GitCommitHash = ""
	// BuiltAt is the actual build datetime
	BuiltAt = ""

	PluginName   = "example-plugin"
	RequiredVars = []string{
		"your",
		"required",
		"variables",
	}

	runCmd = command.NewPluginCommands(
		PluginName,
		Version,
		VersionPostfix,
		GitCommitHash,
		&armory.Armory,
		initializer,
		RequiredVars,
	)
)

// initializer is a custom function to run after the config has been read
// this could be omitted or replaced by something like armory.SetupArmory(c)
func initializer(c *config.Config) (err error) {
	return
}

func main() {
	if VersionPostfix != "" {
		Version = fmt.Sprintf("%s-%s", Version, VersionPostfix)
	}

	err := runCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
