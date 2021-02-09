package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func getConcurrencyNoLimitTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_concurrency_no_limit",
		Description: "Chaos table with high concurrency and no limit (apart from the plugin level limit)",
		List: &plugin.ListConfig{
			Hydrate: getConcurrencyList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "hydrate_call_1",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallColumn1,
				Transform: transform.FromValue(),
			},
			{
				Name:      "hydrate_call_2",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallColumn2,
				Transform: transform.FromValue(),
			},
			{
				Name:      "total_calls",
				Type:      proto.ColumnType_INT,
				Hydrate:   totalHydrateCallsColumn,
				Transform: transform.FromValue(),
			},
		},
	}
}
