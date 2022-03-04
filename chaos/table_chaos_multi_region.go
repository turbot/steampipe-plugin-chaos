package chaos

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

const matrixKeyRegion = "region"

var listCalls int
var hydrateCalls int
var mut sync.Mutex

// build a list of matrix items, one per region
func GetRegions(ctx context.Context, connection *plugin.Connection) []map[string]interface{} {
	// reset counters
	listCalls = 0
	hydrateCalls = 0

	// retrieve regions from connection config
	regions := GetConfig(connection).Regions

	matrix := make([]map[string]interface{}, len(regions))
	for i, region := range regions {
		matrix[i] = map[string]interface{}{matrixKeyRegion: region}
	}

	return matrix
}

func multiRegionTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_multi_region",
		Description: "Chaos table to test the multi region features",
		List: &plugin.ListConfig{
			Hydrate: regionAwareList,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    regionAwareGet,
		},
		GetMatrixItem: GetRegions,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromMatrixItem(matrixKeyRegion)},
			{Name: "list_calls", Type: proto.ColumnType_INT, Hydrate: getListCalls, Transform: transform.FromValue()},
			{Name: "hydrate_calls", Type: proto.ColumnType_INT, Hydrate: getHydrateCalls, Transform: transform.FromValue()},
			{Name: "c1", Type: proto.ColumnType_STRING, Hydrate: doHydrate1},
			{Name: "c2", Type: proto.ColumnType_STRING, Hydrate: doHydrate2},
			{Name: "c3", Type: proto.ColumnType_STRING, Hydrate: doHydrate3},
			{Name: "c4", Type: proto.ColumnType_STRING, Hydrate: doHydrate4},
			{Name: "c5", Type: proto.ColumnType_STRING, Hydrate: doHydrate5},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:    getHydrateCalls,
				Depends: []plugin.HydrateFunc{doHydrate1, doHydrate2, doHydrate3, doHydrate4, doHydrate5},
			},
		},
	}
}

func regionAwareList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	mut.Lock()
	defer mut.Unlock()

	region := regionFromMatrixItem(ctx)
	plugin.Logger(ctx).Warn("regionAwareList", "region", region)
	if region == "" {
		return nil, nil
	}

	plugin.Logger(ctx).Warn("regionAwareList", "region", region)

	const itemsPerRegion = 5
	for i := 0; i < itemsPerRegion; i++ {
		id := buildId(i, region)
		item := map[string]interface{}{"id": id}
		plugin.Logger(ctx).Warn("regionAwareList", "item", item)
		d.StreamListItem(ctx, item)
	}
	// update counter
	listCalls++
	return nil, nil
}

func regionFromMatrixItem(ctx context.Context) string {
	matrixItem := plugin.GetMatrixItem(ctx)
	if matrixItem == nil {
		return ""
	}
	return matrixItem[matrixKeyRegion].(string)
}

func regionAwareGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetStringValue()
	idRegion := regionFromId(id)
	region := regionFromMatrixItem(ctx)
	if region == "" {
		return nil, nil
	}

	if region == idRegion {
		return map[string]interface{}{"id": id, "matching_region": region}, nil
	}
	return nil, nil
}

// build an id from an index and a region
func buildId(i int, region string) string {
	return fmt.Sprintf("%d_%s", i, region)
}

// extract the region from the id
func regionFromId(id string) string {
	parts := strings.Split(id, "_")
	return parts[1]
}

func getListCalls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	mut.Lock()
	defer mut.Unlock()
	hydrateCalls++
	return listCalls, nil
}
func getHydrateCalls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	mut.Lock()
	defer mut.Unlock()
	hydrateCalls++
	return hydrateCalls, nil
}

func doHydrate1(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return doHydrate("c1", h)
}
func doHydrate2(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return doHydrate("c2", h)
}
func doHydrate3(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return doHydrate("c3", h)
}
func doHydrate4(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return doHydrate("c4", h)
}
func doHydrate5(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return doHydrate("c5", h)
}

// hydrate call which does not doe anything interesting
func doHydrate(column string, h *plugin.HydrateData) (interface{}, error) {
	mut.Lock()
	defer mut.Unlock()
	hydrateCalls++
	return map[string]string{column: fmt.Sprintf("%s_val_%d", column, hydrateCalls)}, nil
}
