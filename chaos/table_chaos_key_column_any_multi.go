package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/option"
)

func listKeyColumnsAnyMultipleOperatorTable() *plugin.Table {
	return &plugin.Table{
		Name: "chaos_list_key_columns_any_multiple_operator",
		List: &plugin.ListConfig{
			Hydrate: listKeyColumnsList,
			KeyColumns: plugin.NewKeyColumnSet([]*plugin.KeyColumn{
				{
					Column:    "id",
					Operators: []string{"=", "<", "<=", ">", ">="},
					Optional:  true,
				},
				{
					Column:    "col_1",
					Operators: []string{"<", ">"},
					Optional:  true,
				},
			}, option.WithAtLeast(1)),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "col_1", Type: proto.ColumnType_INT},
		},
	}
}
