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
		Description: "Chaos table to test the retry error handling in List function where list function throws an error after streaming object partially",
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
	for i := 0; i < successfulItemCount; i++ {
		columnValue := fmt.Sprintf("%s-%v", "columnValue", i)
		item := map[string]interface{}{"retry_column": columnValue}
		d.StreamListItem(ctx, item)

	}
	return nil, errors.New(hydrateErrorString)
}
