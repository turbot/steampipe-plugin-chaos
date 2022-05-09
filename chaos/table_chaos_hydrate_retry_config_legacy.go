package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func hydrateRetryConfigLegacyTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_retry_config_legacy",
		Description: "Chaos table to test the Hydrate function with Retry config",
		List: &plugin.ListConfig{
			Hydrate: hydrateRetryConfigList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: hydrateWithRetriableErrors1,
				RetryConfig: &plugin.RetryConfig{
					ShouldRetryError: shouldRetryErrorLegacy,
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "retriable_errors_with_retry_config",
				Type:      proto.ColumnType_STRING,
				Hydrate:   hydrateWithRetriableErrors1,
				Transform: transform.FromValue(),
			},
			{
				Name:      "retriable_errors_with_no_retry_config",
				Type:      proto.ColumnType_STRING,
				Hydrate:   hydrateWithRetriableErrors2,
				Transform: transform.FromValue(),
			},
		},
	}
}
