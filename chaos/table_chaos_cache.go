package chaos

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type listTimeWithID struct {
	Id        int
	UniqueCol int
	TimeCol   string
}

func checkCacheTable() *plugin.Table {
	return &plugin.Table{
		DefaultTransform: transform.FromCamel(),
		Name:             "chaos_cache_check",
		Description:      "Chaos table to print the current time and check the cache functionality.",
		List: &plugin.ListConfig{
			Hydrate: listIdsWithTimeFunction,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "time_col",
					Require:   plugin.Optional,
					Operators: []string{"=", "<", ">", "<=", ">=", "!="},
				},
				{
					Name:      "int_col",
					Require:   plugin.Optional,
					Operators: []string{"=", "<", ">", "<=", ">=", "!="},
				},
				{
					Name:      "float_col",
					Require:   plugin.Optional,
					Operators: []string{"=", "<", ">", "<=", ">=", "!="},
				},
			},
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Hydrate: listIdsWithTimeFunction},
			{Name: "unique_col", Type: proto.ColumnType_INT, Hydrate: listIdsWithTimeFunction},
			{Name: "a", Type: proto.ColumnType_STRING, Hydrate: colAHydrate},
			{Name: "b", Type: proto.ColumnType_STRING, Hydrate: colBHydrate},
			{Name: "c", Type: proto.ColumnType_STRING, Hydrate: colCHydrate},
			{Name: "d", Type: proto.ColumnType_STRING, Hydrate: colDHydrate},
			{Name: "int_col", Type: proto.ColumnType_INT, Hydrate: intColHydrate, Transform: transform.FromValue()},
			{Name: "float_col", Type: proto.ColumnType_DOUBLE, Hydrate: floatColHydrate, Transform: transform.FromValue()},
			{Name: "time_col", Type: proto.ColumnType_STRING, Hydrate: listIdsWithTimeFunction},
			{Name: "delay", Type: proto.ColumnType_STRING, Hydrate: delayHydrate},
			{Name: "long_delay", Type: proto.ColumnType_STRING, Hydrate: longDelayHydrate},
			//{Name: "error_after_delay", Type: proto.ColumnType_STRING, Hydrate: errorAfterDelayHydrate},
		},
	}
}

func listIdsWithTimeFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	time_now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		time_now = time_now.AddDate(0, 1, 1)
		d.StreamListItem(ctx, listTimeWithID{i, rand.Intn(500), time_now.String()})
	}
	return nil, nil
}

func colAHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"A": "a"}
	return item, nil
}

func colBHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"B": "b"}
	return item, nil
}

func colCHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"C": "c"}
	return item, nil
}

func colDHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"D": "d"}
	return item, nil
}

func intColHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(listTimeWithID).Id
	return data, nil
}

func floatColHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(listTimeWithID).Id
	f := float64(data) / 10
	return f, nil
}

func delayHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	delay := 10 * time.Second
	item := map[string]interface{}{"Delay": delay.String()}
	time.Sleep(delay)
	return item, nil
}

func longDelayHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	delay := 10 * time.Hour
	item := map[string]interface{}{"LongDelay": delay.String()}
	time.Sleep(delay)
	return item, nil
}

func errorAfterDelayHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	delay := 10 * time.Second
	time.Sleep(delay)
	return nil, fmt.Errorf("errorAfterDelayHydrate")
}
