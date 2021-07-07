package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func getKeyColumnsSingleEqualsTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_get_key_column_single_equal",
		Get: &plugin.GetConfig{
			Hydrate: listSingleKeyColumnsGet,
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
func listSingleKeyColumnsGet(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	var id int64
	if d.Quals["id"] != nil {
		for _, q := range d.Quals["id"].Quals {
			id = q.Value.GetInt64Value()
		}
	}

	item := map[string]interface{}{"id": id, "col_1": 100 - id}

	return item, nil
}
