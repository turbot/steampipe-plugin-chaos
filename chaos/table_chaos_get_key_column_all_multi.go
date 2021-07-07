package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func getKeyColumnsAllMultipleOperatorTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_get_key_columns_all_multiple_operator",
		Get: &plugin.GetConfig{
			Hydrate: listAllKeyColumnsGet,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Operators: []string{"="},
				},
				{
					Name:      "col_1",
					Operators: []string{"="},
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "col_1", Type: proto.ColumnType_INT},
		},
	}
}

func listAllKeyColumnsGet(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	var id int64
	var col_1 int64

	if d.Quals["id"] != nil {
		for _, q := range d.Quals["id"].Quals {
			id = q.Value.GetInt64Value()
		}
	}
	if d.Quals["col_1"] != nil {
		for _, q := range d.Quals["col_1"].Quals {
			col_1 = q.Value.GetInt64Value()
		}
	}

	item := map[string]interface{}{"id": id, "col_1": col_1}

	return item, nil
}
