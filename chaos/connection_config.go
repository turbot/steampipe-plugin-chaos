package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type chaosConfig struct {
	Regions []string `hcl:"regions,optional"`
}

func ConfigInstance() interface{} {
	return &chaosConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) chaosConfig {
	if connection == nil || connection.Config == nil {
		return chaosConfig{}
	}
	config, _ := connection.Config.(chaosConfig)
	return config
}
