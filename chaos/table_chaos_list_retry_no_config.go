package chaos

import (
	"context"
	"errors"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listRetryNoConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_retry_no_config",
		Description: "Chaos table to test the List function with Retry config in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: listRetryNoConfigList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
		},
	}
}

func listRetryNoConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")

	var failureCount = 200
	for i := 0; i < 5; i++ {
		listMutex.Lock()
		retryListError[listErrorString]++
		listMutex.Unlock()

		errorCount := retryListError[listErrorString]
		if errorCount == failureCount {
			listMutex.Lock()
			retryListError[listErrorString] = 0
			listMutex.Unlock()
			item := map[string]interface{}{"id": i}
			d.StreamListItem(ctx, item)
		}
		return nil, errors.New(listErrorString)

	}
	return nil, nil
}
