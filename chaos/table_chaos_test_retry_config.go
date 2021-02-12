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

var retryError = map[string]int{}
var errorString = "retriableError"
var mut = &sync.Mutex{}

func getRetryConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_retry_config",
		Description: "Chaos table to test the Hydrate function with Retry config in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: getRetryErrorList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: retryHydrate,
				RetryConfig: &plugin.RetryConfig{
					ShouldRetryError: shouldRetryError([]string{errorString}),
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "retriable_errors",
				Type:      proto.ColumnType_STRING,
				Hydrate:   retryHydrate,
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

func getRetryErrorList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func retryHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE HYDRATE CALL")
	var failureCount = 4

	mut.Lock()
	retryError[errorString]++
	mut.Unlock()

	errorCount := retryError[errorString]
	if errorCount == failureCount {
		mut.Lock()
		retryError[errorString] = 0
		mut.Unlock()
		return "SUCCESS", nil
	}

	return nil, errors.New(errorString)
}

func buildRetryHydrate(errorName string, failureCount int) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

		mut.Lock()
		retryError[errorName]++
		mut.Unlock()

		errorCount := retryError[errorName]
		if errorCount == failureCount {
			mut.Lock()
			retryError[errorString] = 0
			mut.Unlock()
			return "SUCCESS", nil
		}

		return nil, errors.New(errorName)
	}

}
