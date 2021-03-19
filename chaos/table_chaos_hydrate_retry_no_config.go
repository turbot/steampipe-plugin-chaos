package chaos

import (
	"context"
	"errors"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func hydrateRetryNoConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_retry_no_config",
		Description: "Chaos table to test the Hydrate function with Default Retry config defined at plugin level in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: hydrateRetryNoConfigErrorList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "retriable_errors",
				Type:      proto.ColumnType_STRING,
				Hydrate:   retryHydrateNoConfig,
				Transform: transform.FromValue(),
			},
		},
	}
}

func hydrateRetryNoConfigErrorList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func retryHydrateNoConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var failureCount = 2

	hydrateMutex.Lock()
	retryHydrateError[hydrateErrorString]++
	hydrateMutex.Unlock()

	errorCount := retryHydrateError[hydrateErrorString]
	if errorCount == failureCount {
		hydrateMutex.Lock()
		retryHydrateError[hydrateErrorString] = 0
		hydrateMutex.Unlock()
		return "SUCCESS", nil
	}

	return nil, errors.New(hydrateErrorString)
}
