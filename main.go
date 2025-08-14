package main

import (
	"github.com/turbot/steampipe-plugin-chaos/v2/chaos"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos.Plugin})
}
