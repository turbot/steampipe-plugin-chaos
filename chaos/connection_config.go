package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/schema"
)

type chaosConfig struct {
	Regions []string `cty:"regions"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"regions": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
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
