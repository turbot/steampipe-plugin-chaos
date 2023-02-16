package chaos

import (
	"context"
	"math/rand"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func cacheDelayQualsTable() *plugin.Table {
	return &plugin.Table{
		DefaultTransform: transform.FromCamel(),
		Name:             "chaos_cache_with_delay_quals",
		Description:      "Chaos table to check the cache functionality.",
		List: &plugin.ListConfig{
			Hydrate: listUnique,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "unique_col",
					Require:   plugin.Optional,
					Operators: []string{"=", "<", ">", "<=", ">=", "!="},
				},
				{
					Name:      "delay",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
		},

		Columns: []*plugin.Column{
			{Name: "unique_col", Type: proto.ColumnType_INT, Hydrate: listUnique},
			{Name: "a", Type: proto.ColumnType_STRING, Hydrate: colAHydrate},
			{Name: "b", Type: proto.ColumnType_STRING, Hydrate: colBHydrate},
			{Name: "c", Type: proto.ColumnType_STRING, Hydrate: colCHydrate},
			{Name: "d", Type: proto.ColumnType_STRING, Hydrate: colDHydrate},
			{Name: "delay", Type: proto.ColumnType_INT, Transform: transform.FromQual("delay"), Default: 0},
		},
	}
}

func listUnique(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	val := rand.Intn(5000)
	item := map[string]interface{}{"UniqueCol": val}
	if d.KeyColumnQuals["delay"] != nil {
		time.Sleep(time.Duration(d.KeyColumnQuals["delay"].GetInt64Value() * int64(time.Second)))
	}
	d.StreamListItem(ctx, item)
	return nil, nil
}

func aHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"A": "a"}
	return item, nil
}

func bHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"B": "b"}
	return item, nil
}

func cHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"C": "c"}
	return item, nil
}

func dHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"D": "d"}
	return item, nil
}
