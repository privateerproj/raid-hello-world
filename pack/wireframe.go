package pack

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
)

// GetProbes returns a list of probe objects
func GetProbes() []raidengine.Strike {
	return Policies["PCI"]
}

// TODO: is this the best way to handle this?
// Policies contains a list of different functions to run for each policy type
var Policies = map[string][]raidengine.Strike{
	"PCI": {
		KnockKnock,
	},
	"CIS": {
		KnockKnock,
		ChickenCrossedRoad,
	},
}
