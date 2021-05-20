package chaos

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func listStreamingTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_streaming_table",
		Description: "Chaos table which streams blocks of 1000 rows with random data per second for 30 seconds",

		List: &plugin.ListConfig{
			Hydrate: streamList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "data", Type: proto.ColumnType_INT, Hydrate: rowData, Transform: transform.FromValue()},
		},
	}
}

func streamList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	startTime := time.Now()
	ticker := time.NewTicker(1 * time.Second)
	blocksSent := 0
	// wait before sending out the first block
	time.Sleep(5 * time.Second)

	for t := range ticker.C {
		if t.Sub(startTime) > (10 * time.Second) {
			ticker.Stop()
			break
		}
		for i := 0; i < 1000; i++ {
			item := map[string]interface{}{"id": (blocksSent * 1000) + i}
			d.StreamListItem(ctx, item)
		}
		blocksSent++
	}
	return nil, nil
}

func rowData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := id * 3
	return columnVal, nil
}
