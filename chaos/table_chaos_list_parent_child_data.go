package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

type parentData struct {
	id int
}

type childData struct {
	id           int
	child_column string
}

func chaosListParentChildHydrateTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_parent_child_hydrate_data",
		Description: "Chaos table to test the List calls having parent-child dependencies with all the possible scenarios like errors, panics and delays at both parent and child levels",
		List: &plugin.ListConfig{
			Hydrate: listChildHydrateTable,
			// ParentHydrate: listParentHydrateTable,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTable,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "child_column", Type: proto.ColumnType_STRING, Description: "Column to test the the parent list function with fatal error after streaming some rows"},
		},
	}
}

func listChildHydrateTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 2; i++ {
		item := map[string]interface{}{"id": i, "child_column": fmt.Sprintf("child_column-%d", i)}
		d.StreamListItem(ctx, item)
	}

	return nil, nil

}

func getTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetInt64Value()

	item := map[string]interface{}{"id": id, "child_column": fmt.Sprintf("child_column-get-%v", id)}
	d.StreamListItem(ctx, item)
	return nil, nil
}
