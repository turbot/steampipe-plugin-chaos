package main

import (
	"log"

	"github.com/turbot/steampipe-plugin-chaos/chaos"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	log.Printf("[WARN] PLUGIN MAIN")
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos.Plugin})
}
