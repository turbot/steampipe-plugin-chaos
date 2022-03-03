package chaos

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func chaosLimitTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_limit",
		Description: "Chaos table to check the limit functionality.",
		List: &plugin.ListConfig{
			Hydrate: listLimits,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "c1",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "c2",
					Require:   plugin.Optional,
					Operators: []string{"<="},
				},
				{
					Name:      "c3",
					Require:   plugin.Optional,
					Operators: []string{">="},
				},
			},
		},

		Columns: []*plugin.Column{
			{Name: "c1", Type: proto.ColumnType_INT},
			{Name: "c2", Type: proto.ColumnType_INT},
			{Name: "c3", Type: proto.ColumnType_INT},
			{Name: "c4", Type: proto.ColumnType_TIMESTAMP},
			{Name: "c5", Type: proto.ColumnType_STRING},
			{Name: "c6", Type: proto.ColumnType_INT},
			{Name: "limit_value", Type: proto.ColumnType_INT},
		},
	}
}

func listLimits(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"c1": i, "c2": i + 1, "c3": i + 2, "c4": time.Now(), "c5": "num", "c6": i + 5, "limit_value": d.QueryContext.Limit}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
