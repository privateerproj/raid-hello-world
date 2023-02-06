package pack

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
)

// GetProbes returns a list of probe objects
func GetProbes() []raidengine.Strike {
	return Policies["PCI"]
}

// Policies contains a list of different functions to run for each policy type
var Policies = map[string][]raidengine.Strike{
	"PCI": []raidengine.Strike{KnockKnock},
	"CIS": []raidengine.Strike{KnockKnock, ChickenCrossedRoad},
}
