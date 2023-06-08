package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func listKeyColumnsSingleEqualsTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_list_key_column_single_equal",
		List: &plugin.ListConfig{
			Hydrate: listKeyColumnsList,
			KeyColumns: []*plugin.KeyColumn{{
				Name:      "id",
				Operators: []string{"="},
			}},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "col_1", Type: proto.ColumnType_INT},
		},
	}

}
func listKeyColumnsList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	for i := 0; i < 10; i++ {
		item := map[string]interface{}{"id": i, "col_1": 100 - i}
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
