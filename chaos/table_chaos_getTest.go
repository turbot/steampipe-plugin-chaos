package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func getTestTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_get_test",
		Description: "Chaos table to test the get function",
		List: &plugin.ListConfig{
			Hydrate: getTableList,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTableGet,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_1", Type: proto.ColumnType_STRING},
			{Name: "column_2", Type: proto.ColumnType_STRING},
			{Name: "column_3", Type: proto.ColumnType_STRING},
		},
	}
}

func getTableList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		id := i
		column1 := fmt.Sprintf("column_1-%d", id)
		column2 := fmt.Sprintf("column_2-%d", id)
		column3 := fmt.Sprintf("column_3-%d", id)

		item := map[string]interface{}{"id": id, "column_1": column1, "column_2": column2, "column_3": column3}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func getTableGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetInt64Value()
	column1 := fmt.Sprintf("column_1-%d", id)
	column2 := fmt.Sprintf("column_2-%d", id)
	column3 := fmt.Sprintf("column_3-%d", id)

	item := map[string]interface{}{"id": id, "column_1": column1, "column_2": column2, "column_3": column3}
	return item, nil
}
