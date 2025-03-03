package main

import (
	"fmt"

	"os"

	"plugin-example-plugin/evaluations"

	"github.com/privateerproj/privateer-sdk/command"
	"github.com/privateerproj/privateer-sdk/pluginkit"
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

)

func main() {
	if VersionPostfix != "" {
		Version = fmt.Sprintf("%s-%s", Version, VersionPostfix)
	}

	// NewVessel may take a payload for all suites to reference
	pvtrVessel := pluginkit.NewVessel(PluginName, evaluations.LoadData(), RequiredVars)

	// Evaluation Suite may optionally take a payload to selectively override the data specified in NewVessel
	pvtrVessel.AddEvaluationSuite("FINOS_CCC", nil, evaluations.FINOS_CCC)

	runCmd := command.NewPluginCommands(
		PluginName,
		Version,
		VersionPostfix,
		GitCommitHash,
		pvtrVessel,
	)

	err := runCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}