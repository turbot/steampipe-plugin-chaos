package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func getKeyColumnsAnyMultipleOperatorTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_get_key_columns_any_multiple_operator",
		Get: &plugin.GetConfig{
			Hydrate: listAnyKeyColumnsGet,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Operators: []string{"="},
					Require:   plugin.Optional,
				},
				{
					Name:      "col_1",
					Operators: []string{"="},
					Require:   plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "col_1", Type: proto.ColumnType_INT},
		},
	}
}

func listAnyKeyColumnsGet(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	var id int64
	
	if d.Quals["id"] != nil {
		for _, q := range d.Quals["id"].Quals {
			id = q.Value.GetInt64Value()
		}
	}

	item := map[string]interface{}{"id": id, "col_1": 100 - id}

	return item, nil
}
