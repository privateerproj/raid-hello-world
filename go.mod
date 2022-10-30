module github.com/privateerproj/privateer-pack-wireframe

go 1.14

require (
	github.com/cucumber/godog v0.11.0
	github.com/hashicorp/go-hclog v0.15.0 // indirect
	github.com/markbates/pkger v0.17.1
	github.com/privateerproj/privateer-sdk v0.0.0
)

// For Development Only
replace github.com/privateerproj/privateer-sdk => ../privateer-sdk
