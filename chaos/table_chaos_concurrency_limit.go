package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func getConcurrencyLimitTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_concurrency_limit",
		Description: "Chaos table to test the concurrency limit of hydrate calls",
		List: &plugin.ListConfig{
			Hydrate: getConcurrencyList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:           hydrateCallColumn1,
				MaxConcurrency: 20,
			},
			{
				Func:           hydrateCallColumn2,
				MaxConcurrency: 10,
			},
			{
				Func:           totalHydrateCallsColumn,
				MaxConcurrency: 5,
			},
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
