package main

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-chaos/chaos"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos.Plugin})
}
