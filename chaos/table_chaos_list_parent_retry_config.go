package chaos

import (
	"context"
	"errors"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listParentRetryTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_parent_retry_config",
		Description: "Chaos table to test the retry error handling in parent-child function where the parent list function fails after streaming some objects",
		List: &plugin.ListConfig{
			ParentHydrate: parentRetryErrorList,
			Hydrate:       childRetryList,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
		},

		Columns: []*plugin.Column{
			{Name: "parent_id", Type: proto.ColumnType_INT},
			{Name: "child_column", Type: proto.ColumnType_STRING},
		},
	}
}

func parentRetryErrorList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"parent_id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, errors.New(listErrorString)
}

func childRetryList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	parentItem := h.Item.(map[string]interface{})

	for i := 0; i < successfulChildItemCount; i++ {
		parentId := parentItem["parent_id"].(int)
		column := fmt.Sprintf("child-%d", i)
		item := map[string]interface{}{"parent_id": parentId, "child_column": column}
		d.StreamLeafListItem(ctx, item)
	}
	return nil, nil
}
