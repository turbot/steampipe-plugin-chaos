package chaos

import (
	"context"
	"errors"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

var hydrateWithRetriableErrors1ErrorCount = 0
var hydrateWithRetriableErrors2ErrorCount = 0

func hydrateRetryConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_retry_config",
		Description: "Chaos table to test the Hydrate function with Retry config",
		List: &plugin.ListConfig{
			Hydrate: hydrateRetryConfigList,
		},
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

func hydrateRetryConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func hydrateWithRetriableErrors1(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var failureCount = 2

	hydrateWithRetriableErrors1ErrorCount++

	if hydrateWithRetriableErrors1ErrorCount == failureCount {
		log.Printf("[INFO] chaos_hydrate_retry_config hydrateWithRetriableErrors1 error count %d, returning success", hydrateWithRetriableErrors1ErrorCount)
		hydrateWithRetriableErrors1ErrorCount = 0
		return "SUCCESS", nil
	}

	log.Printf("[INFO] chaos_hydrate_retry_config hydrateWithRetriableErrors error count %d, returning error", hydrateWithRetriableErrors1ErrorCount)

	return nil, errors.New(retriableErrorString)
}

func hydrateWithRetriableErrors2(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var failureCount = 2

	hydrateWithRetriableErrors2ErrorCount++

	if hydrateWithRetriableErrors2ErrorCount == failureCount {
		log.Printf("[INFO] chaos_hydrate_retry_config hydrateWithRetriableErrors1 error count %d, returning success", hydrateWithRetriableErrors2ErrorCount)
		hydrateWithRetriableErrors2ErrorCount = 0
		return "SUCCESS", nil
	}

	log.Printf("[INFO] chaos_hydrate_retry_config hydrateWithRetriableErrors error count %d, returning error", hydrateWithRetriableErrors2ErrorCount)

	return nil, errors.New(retriableErrorString)
}
