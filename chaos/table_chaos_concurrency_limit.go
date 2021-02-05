package chaos

import (
	"context"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func getConcurrencyLimitTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_concurrency_limit_test",
		Description: "Chaos table to test the concurrency limit of hydrate calls",
		List: &plugin.ListConfig{
			Hydrate: getConcurrencyList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:           hydrateCallColumn1,
				MaxConcurrency: 40,
			},
			{
				Func:           hydrateCallColumn2,
				MaxConcurrency: 5,
			},
			{
				Func: totalHydrateCallsColumn,
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

func getConcurrencyList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 100; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func hydrateCallColumn1(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE CALL hydrateCallColumn1")
	count, _ := doHydrateCall("hydrateCallColumn1")

	return count, nil
}

func hydrateCallColumn2(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE CALL hydrateCallColumn2")
	count, _ := doHydrateCall("hydrateCallColumn2")

	return count, nil
}

func totalHydrateCallsColumn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE CALL totalHydrateCallsColumn")
	_, totalCount := doHydrateCall("totalCalls")

	return totalCount, nil
}
