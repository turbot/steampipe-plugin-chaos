package chaos

import (
	"context"
	"errors"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listRetryPartialTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_retry_partial",
		Description: "Chaos table to test the List function with Retry config in case of non fatal error",
		List: &plugin.ListConfig{
			Hydrate: listRetryPartialList,
		},
		Columns: []*plugin.Column{
			{Name: "retry_column", Type: proto.ColumnType_STRING},
		},
	}
}

func listRetryPartialList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	var successfulItemCount = 2
	listMutex.Lock()
	errorCount := retryListError[listErrorString]
	retryListError[listErrorString] = errorCount + 1
	listMutex.Unlock()

	listMutex.Lock()
	retryListError[listErrorString] = 0
	listMutex.Unlock()
	for i := 0; i < successfulItemCount; i++ {
		columnValue := fmt.Sprintf("%s-%v", "columnValue", i)
		item := map[string]interface{}{"retry_column": columnValue}
		d.StreamListItem(ctx, item)

	}
	return nil, errors.New(hydrateErrorString)
}
