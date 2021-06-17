package chaos

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func hydrateColumnsDependencyTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_columns_dependency",
		Description: "Chaos table to test dependencies between hydrate functions",

		List: &plugin.ListConfig{
			Hydrate: hydrateList,
		},

		HydrateDependencies: []plugin.HydrateDependencies{
			{
				Func:    hydrate2,
				Depends: []plugin.HydrateFunc{hydrate1},
			},
			{
				Func:    hydrate3,
				Depends: []plugin.HydrateFunc{hydrate2},
			},
			{
				Func:    hydrate5,
				Depends: []plugin.HydrateFunc{hydrate4, hydrate1},
			},
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "hydrate_column_1", Type: proto.ColumnType_STRING, Hydrate: hydrate1},
			{Name: "hydrate_column_2", Type: proto.ColumnType_STRING, Hydrate: hydrate2},
			{Name: "hydrate_column_3", Type: proto.ColumnType_STRING, Hydrate: hydrate3},
			{Name: "hydrate_column_4", Type: proto.ColumnType_STRING, Hydrate: hydrate4},
			{Name: "hydrate_column_5", Type: proto.ColumnType_STRING, Hydrate: hydrate5},
		},
	}
}

func hydrateInputKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	item := quals["id"].GetInt64Value()
	return item, nil
}

func hydrateList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	for i := 0; i < 2; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)

	}

	return nil, nil
}

func hydrateGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)

	item := map[string]interface{}{"id": id}
	return item, nil

}

func hydrate1(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(2 * time.Second)

	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := fmt.Sprintf("hydrate1-%d", id)

	item := map[string]interface{}{"hydrate_column_1": columnVal}

	return item, nil
}

func hydrate2(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	hydrate1Results := h.HydrateResults["hydrate1"].(map[string]interface{})

	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := fmt.Sprintf("hydrate2-%d-%s", id, hydrate1Results["hydrate_column_1"])

	item := map[string]interface{}{"hydrate_column_2": columnVal}

	return item, nil
}

func hydrate3(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	hydrate2Results := h.HydrateResults["hydrate2"].(map[string]interface{})

	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := fmt.Sprintf("hydrate3-%d-%s", id, hydrate2Results["hydrate_column_2"])

	item := map[string]interface{}{"hydrate_column_3": columnVal}

	return item, nil
}

func hydrate4(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := fmt.Sprintf("hydrate4-%d", id)

	item := map[string]interface{}{"hydrate_column_4": columnVal}

	return item, nil
}

func hydrate5(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	hydrate1Results := h.HydrateResults["hydrate1"].(map[string]interface{})
	hydrate4Results := h.HydrateResults["hydrate4"].(map[string]interface{})

	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := fmt.Sprintf("hydrate5-%d-%s-%s", id, hydrate4Results["hydrate_column_4"], hydrate1Results["hydrate_column_1"])

	item := map[string]interface{}{"hydrate_column_5": columnVal}

	return item, nil
}
