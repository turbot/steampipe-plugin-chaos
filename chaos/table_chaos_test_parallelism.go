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

func getTestParallelismTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_parallel_hydrate_test",
		Description: "Chaos table to test the parallelism in the hydrate calls",
		List: &plugin.ListConfig{
			Hydrate: getParallelList,
		},
		HydrateDependencies: []plugin.HydrateDependencies{
			{
				Func:    hydrateCallC,
				Depends: []plugin.HydrateFunc{hydrateCallA, hydrateCallB},
			},
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "hydrate_a",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallA,
				Transform: transform.FromValue(),
			},
			{
				Name:      "hydrate_b",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallB,
				Transform: transform.FromValue(),
			},
			{
				Name:      "hydrate_c",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallC,
				Transform: transform.FromValue(),
			},
			{
				Name:      "hydrate_d",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallD,
				Transform: transform.FromValue(),
			},
			{
				Name:      "total_hydrate_calls",
				Type:      proto.ColumnType_INT,
				Hydrate:   hydrateCallTotal,
				Transform: transform.FromValue(),
			},
		},
	}
}

func getParallelList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 100; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func hydrateCallA(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE CALL A")
	count, _ := doHydrateCall("a")

	return count, nil
}

func hydrateCallB(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE CALL B")
	count, _ := doHydrateCall("b")

	return count, nil
}

func hydrateCallC(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	count, _ := doHydrateCall("c")

	return count, nil
}

func hydrateCallD(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	count, _ := doHydrateCall("d")

	return count, nil
}

func hydrateCallTotal(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE CALL C")
	_, totalCount := doHydrateCall("hydrate calls total")

	return totalCount, nil
}

// increment hydrate count for this name, do some work(sleep), decrement hydrate count
// return number of instances of this hydrate function running, and total number of hydrate calls running
func doHydrateCall(name string) (int, int) {
	log.Println("[WARN] INSIDE DO HYDRATE")
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
