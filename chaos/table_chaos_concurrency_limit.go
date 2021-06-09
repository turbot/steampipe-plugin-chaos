package chaos

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

var hydrateCount = map[string]int{}
var mutex = &sync.Mutex{}

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

// increment hydrate count for this name, do some work(sleep), decrement hydrate count
// return number of instances of this hydrate function running, and total number of hydrate calls running
func doHydrateCall(name string) (int, int) {
	mutex.Lock()
	hydrateCount[name] = hydrateCount[name] + 1
	callsForThisHydrate := hydrateCount[name]
	var totalCalls = 0
	for _, calls := range hydrateCount {
		totalCalls += calls
	}
	mutex.Unlock()

	time.Sleep(1 * time.Second)

	mutex.Lock()
	hydrateCount[name] = hydrateCount[name] - 1
	mutex.Unlock()

	return callsForThisHydrate, totalCalls
}
