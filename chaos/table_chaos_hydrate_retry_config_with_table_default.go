package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func hydrateRetryConfigWithTableDefaultTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_retry_config_with_table_default",
		Description: "Chaos table to test the Hydrate function with Retry config, and to test table defaulting of the retry config",
		List: &plugin.ListConfig{
			Hydrate: hydrateRetryConfigList,
		},
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: hydrateWithRetriableErrors1,
				RetryConfig: &plugin.RetryConfig{
					ShouldRetryErrorFunc: shouldRetryError,
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
