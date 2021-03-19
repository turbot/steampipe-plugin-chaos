package chaos

import (
	"context"
	"errors"
	"sync"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

var retryGetError = map[string]int{}
var getErrorString = "retriableError"
var getMutex = &sync.Mutex{}

func getRetryConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_get_retry_config",
		Description: "Chaos table to test the Get function with Retry config in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: getRetryList,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    retryGetFunc,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "retriable_error", Type: proto.ColumnType_STRING},
		},
	}
}

func getRetryList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"id": i, "retriable_error": "SUCCESS"}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func retryGetFunc(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetInt64Value()
	var failureCount = 200

	getMutex.Lock()
	retryGetError[getErrorString]++
	getMutex.Unlock()

	errorCount := retryGetError[getErrorString]
	if errorCount == failureCount {
		getMutex.Lock()
		retryGetError[getErrorString] = 0
		getMutex.Unlock()
		item := map[string]interface{}{"id": id, "retriable_error": "SUCCESS"}
		return item, nil
	}

	return nil, errors.New(getErrorString)
}
