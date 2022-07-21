package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func listKeyColumnsAllMultipleOperatorTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_list_key_columns_all_multiple_operator",
		List: &plugin.ListConfig{
			Hydrate: listKeyColumnsList,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Operators: []string{"=", "<", "<=", ">", ">="},
				},
				{
					Name:      "col_1",
					Operators: []string{"<", ">"},
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "col_1", Type: proto.ColumnType_INT},
		},
	}
}
