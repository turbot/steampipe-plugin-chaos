package main

import (
	"github.com/turbot/steampipe-plugin-chaos/chaos"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos.Plugin})
}
