package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type chaosConfig struct {
	Regions []string `hcl:"regions,optional"`
}

func ConfigInstance() interface{} {
	return &chaosConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) chaosConfig {
	if connection == nil {
		return chaosConfig{}
	}
	raw := connection.GetConfig()
	if raw == nil {
		return chaosConfig{}
	}
	config, _ := raw.(chaosConfig)
	return config
}
