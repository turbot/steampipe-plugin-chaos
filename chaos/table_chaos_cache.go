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
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "a", Type: proto.ColumnType_INT},
			{Name: "b", Type: proto.ColumnType_INT},
			{Name: "c", Type: proto.ColumnType_INT},
			{Name: "d", Type: proto.ColumnType_INT},
			{Name: "time_now", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

func listIdsWithTimeFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"id": i, "a": i, "b": i, "c": i, "d": i, "time_now": time.Now()}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
