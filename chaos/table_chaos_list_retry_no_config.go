package chaos

import (
	"context"
	"errors"
	"fmt"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listRetryNoConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_retry_no_config",
		Description: "Chaos table to test the List function with the default retry config at plugin level in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: listRetryNoConfigList,
		},
		Columns: []*plugin.Column{
			{Name: "retry_column", Type: proto.ColumnType_STRING},
		},
	}
}

func listRetryNoConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")

	var failureCount = 200
	listMutex.Lock()
	errorCount := retryListError[listErrorString]
	retryListError[listErrorString] = errorCount + 1
	listMutex.Unlock()

	if errorCount < failureCount {
		return nil, errors.New(listErrorString)
	}
	listMutex.Lock()
	retryListError[listErrorString] = 0
	listMutex.Unlock()
	for i := 0; i < 5; i++ {
		columnValue := fmt.Sprintf("%s-%v", "columnValue", i)
		item := map[string]interface{}{"retry_column": columnValue}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
