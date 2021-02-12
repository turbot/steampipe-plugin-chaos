package chaos

import (
	"context"
	"errors"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func getRetryNoConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_retry_no_config",
		Description: "Chaos table to test the Hydrate function with Default Retry config in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: getRetryErrorList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "retriable_errors",
				Type:      proto.ColumnType_STRING,
				Hydrate:   retryHydrate,
				Transform: transform.FromValue(),
			},
		},
	}
}

func getRetryErrorNoConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func retryHydrateNoConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
