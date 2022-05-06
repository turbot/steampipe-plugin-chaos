package chaos

import (
	"context"
	"errors"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

var hydrateRetryConfigErrorCount = 0

func hydrateRetryConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_retry_config",
		Description: "Chaos table to test the Hydrate function with Retry config",
		List: &plugin.ListConfig{
			Hydrate: hydrateRetryConfigList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: hydrateWithRetriableErrors,
				RetryConfig: &plugin.RetryConfig{
					ShouldRetryError: shouldRetryError,
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "retriable_errors",
				Type:      proto.ColumnType_STRING,
				Hydrate:   hydrateWithRetriableErrors,
				Transform: transform.FromValue(),
			},
		},
	}
}

func hydrateRetryConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func hydrateWithRetriableErrors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var failureCount = 2

	hydrateRetryConfigErrorCount++

	if hydrateRetryConfigErrorCount == failureCount {
		log.Printf("[WARN] chaos_hydrate_retry_config hydrateWithRetriableErrors error count %d, returning success", hydrateRetryConfigErrorCount)
		hydrateRetryConfigErrorCount = 0
		return "SUCCESS", nil
	}

	log.Printf("[WARN] chaos_hydrate_retry_config hydrateWithRetriableErrors error count %d, returning error", hydrateRetryConfigErrorCount)

	return nil, errors.New(retriableErrorString)
}
