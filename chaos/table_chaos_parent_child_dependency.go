package chaos

import (
	"context"
	"fmt"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func parentChildTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_parent_child_dependency",
		Description: "Chaos table to test the parent-child dependencies in the list function",
		List: &plugin.ListConfig{
			ParentHydrate: parentList,
			Hydrate:       childList,
		},

		Columns: []*plugin.Column{
			{Name: "parent_id", Type: proto.ColumnType_INT},
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_1", Type: proto.ColumnType_STRING},
		},
	}
}

func parentList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	log.Println("[TRACE] parentList")

	for i := 0; i < 10; i++ {
		id := i
		item := map[string]interface{}{"parent_id": id}
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func childList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// NOTE: case the parent item to whatever type the parent hydrate function returns
	parentItem := h.Item.(map[string]interface{})

	for i := 0; i < 10; i++ {
		id := i
		parentId := parentItem["parent_id"].(int)
		column := fmt.Sprintf("child-%d-%d", id, parentId)

		item := map[string]interface{}{"parent_id": parentId, "id": id, "column_1": column}

		d.StreamLeafListItem(ctx, item)
	}

	return nil, nil
}
