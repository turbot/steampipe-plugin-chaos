package chaos

import (
	"context"
	"errors"
	log "log"
	"sync"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

var retryHydrateError = map[string]int{}
var hydrateErrorString = "retriableError"
var hydrateMutex = &sync.Mutex{}

func hydrateRetryConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_retry_config",
		Description: "Chaos table to test the Hydrate function with Retry config in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: hydrateRetryConfigList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: retryHydrateConfig,
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
				Hydrate:   retryHydrateConfig,
				Transform: transform.FromValue(),
			},
			{
				Name:      "fatal_error",
				Type:      proto.ColumnType_STRING,
				Hydrate:   buildRetryHydrate("fatalError", 4),
				Transform: transform.FromValue(),
			},
		},
	}
}

func hydrateRetryConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func retryHydrateConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE HYDRATE CALL")
	var failureCount = 2

	hydrateMutex.Lock()
	retryHydrateError[hydrateErrorString]++
	errorCount := retryHydrateError[hydrateErrorString]
	hydrateMutex.Unlock()

	if errorCount == failureCount {
		hydrateMutex.Lock()
		retryHydrateError[hydrateErrorString] = 0
		hydrateMutex.Unlock()
		return "SUCCESS", nil
	}

	return nil, errors.New(hydrateErrorString)
}

func buildRetryHydrate(errorName string, failureCount int) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

		hydrateMutex.Lock()
		retryHydrateError[errorName]++
		hydrateMutex.Unlock()

		errorCount := retryHydrateError[errorName]
		if errorCount == failureCount {
			hydrateMutex.Lock()
			retryHydrateError[hydrateErrorString] = 0
			hydrateMutex.Unlock()
			return "SUCCESS", nil
		}

		return nil, errors.New(errorName)
	}

}
