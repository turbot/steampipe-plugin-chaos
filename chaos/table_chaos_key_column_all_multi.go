package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listKeyColumnsAllMultipleOperatorTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_list_key_columns_all_multiple_operator",
		List: &plugin.ListConfig{
			Hydrate: listKeyColumnsList,
			KeyColumns: plugin.NewKeyColumnSet([]*plugin.KeyColumn{
				{
					Column:    "id",
					Operators: []string{"=", "<", "<=", ">", ">="},
				},
				{
					Column:    "col_1",
					Operators: []string{"<", ">"},
				},
			}),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "col_1", Type: proto.ColumnType_INT},
		},
	}
}
