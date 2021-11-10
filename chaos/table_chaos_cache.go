package chaos

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func checkCacheTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_cache_check",
		Description: "Chaos table to print the current time and check the cache functionality.",
		List: &plugin.ListConfig{
			Hydrate: listIdsWithTimeFunction,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Hydrate: listIdsWithTimeFunction},
			{Name: "a", Type: proto.ColumnType_STRING, Hydrate: listColAFunction},
			{Name: "b", Type: proto.ColumnType_STRING, Hydrate: listColBFunction},
			{Name: "c", Type: proto.ColumnType_STRING, Hydrate: listColCFunction},
			{Name: "d", Type: proto.ColumnType_STRING, Hydrate: listColDFunction},
			{Name: "time_now", Type: proto.ColumnType_TIMESTAMP, Hydrate: listIdsWithTimeFunction},
		},
	}
}

func listIdsWithTimeFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"id": i, "time_now": time.Now()}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func listColAFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"a": "a"}
	return item, nil
}

func listColBFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"b": "b"}
	return item, nil
}

func listColCFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"c": "c"}
	return item, nil
}

func listColDFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"d": "d"}
	return item, nil
}
