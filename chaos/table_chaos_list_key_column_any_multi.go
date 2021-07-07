package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listKeyColumnsAnyMultipleOperatorTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_list_key_columns_any_multiple_operator",
		List: &plugin.ListConfig{
			Hydrate: listKeyColumnsList,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Operators: []string{"=", "<", "<=", ">", ">="},
					Require:   plugin.Optional,
				},
				{
					Name:      "col_1",
					Operators: []string{"<", ">"},
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
