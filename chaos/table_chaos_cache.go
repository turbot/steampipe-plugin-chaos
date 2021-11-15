package chaos

import (
	"context"
	"fmt"
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
			{Name: "a", Type: proto.ColumnType_STRING, Hydrate: colAHydrate},
			{Name: "b", Type: proto.ColumnType_STRING, Hydrate: colBHydrate},
			{Name: "c", Type: proto.ColumnType_STRING, Hydrate: colCHydrate},
			{Name: "d", Type: proto.ColumnType_STRING, Hydrate: colDHydrate},
			{Name: "time_now", Type: proto.ColumnType_TIMESTAMP, Hydrate: listIdsWithTimeFunction},
			{Name: "delay", Type: proto.ColumnType_STRING, Hydrate: delayHydrate},
			{Name: "long_delay", Type: proto.ColumnType_STRING, Hydrate: longDelayHydrate},
			//{Name: "error_after_delay", Type: proto.ColumnType_STRING, Hydrate: errorAfterDelayHydrate},
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

func colAHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"a": "a"}
	return item, nil
}

func colBHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"b": "b"}
	return item, nil
}

func colCHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"c": "c"}
	return item, nil
}

func colDHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"d": "d"}
	return item, nil
}

func delayHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	delay := 3 * time.Second
	item := map[string]interface{}{"delay": delay.String()}
	time.Sleep(delay)
	return item, nil
}

func longDelayHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	delay := 10 * time.Hour
	item := map[string]interface{}{"delay": delay.String()}
	time.Sleep(delay)
	return item, nil
}

func errorAfterDelayHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	delay := 10 * time.Second
	time.Sleep(delay)
	return nil, fmt.Errorf("errorAfterDelayHydrate")
}
