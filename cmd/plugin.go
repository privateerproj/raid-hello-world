package cmd

import (
	"github.com/privateerproj/privateer-plugin-example/armory"

	"github.com/privateerproj/privateer-sdk/pluginkit"
)

var (
	Vessel = pluginkit.Vessel{
		PluginName: "example", // Double check that this is what you want the plugin to be named
    Armory:     &armory.Armory,
	} // Used by the plugin or debug function to run the Plugin
)

type Plugin struct {}

// Plugin.Start() is called by plugin.Serve
func (p *Plugin) Start() (err error) {
	err = Vessel.Mobilize()
	return
}
