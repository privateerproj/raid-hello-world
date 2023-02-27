module github.com/privateerproj/privateer-pack-wireframe

go 1.14

require (
	github.com/cucumber/godog v0.12.6
	github.com/hashicorp/go-hclog v1.2.0
	github.com/privateerproj/privateer-sdk v0.0.0
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.15.0
)

// For Development Only
replace github.com/privateerproj/privateer-sdk => ../privateer-sdk
