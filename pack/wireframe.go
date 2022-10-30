package pack

import (
	"github.com/markbates/pkger"
	"github.com/privateerproj/privateer-pack-wireframe/internal/welcome"
	"github.com/privateerproj/privateer-sdk/probeengine"
)

// GetProbes returns a list of probe objects
func GetProbes() []probeengine.Probe {
	return []probeengine.Probe{
		welcome.Probe,
	}
}

func init() {
	// pkger.Include is a no-op that directs the pkger tool to include the desired file or folder.
	pkger.Include("/internal/welcome/welcome.feature")
}
