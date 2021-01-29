package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

var regions = []string{
	"us-east-2",
	"us-east-1",
	"us-west-1",
	"us-west-2",
	"af-south-1",
	"ap-east-1",
	"ap-south-1",
	"ap-northeast-3",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"ca-central-1",
	"cn-north-1",
	"cn-northwest-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-south-1",
	"eu-west-3",
	"eu-north-1",
	"me-south-1",
	"sa-east-1",
	"us-gov-east-1",
	"us-gov-west-1"}

func ListRegions(listFunc plugin.HydrateFunc) plugin.HydrateFunc {
	return func(ctx context.Context, queryData *plugin.QueryData, hydrateData *plugin.HydrateData) (interface{}, error) {
		// build a list of hydrate params objects - one per region
		paramsList := make([]map[string]string, len(regions))
		for i, region := range regions {
			paramsList[i] = map[string]string{"region": region}
		}
		return plugin.ListForPartitions(ctx, queryData, hydrateData, listFunc, paramsList)
	}
}

func GetRegions(getFunc plugin.HydrateFunc) plugin.HydrateFunc {
	return func(ctx context.Context, queryData *plugin.QueryData, hydrateData *plugin.HydrateData) (interface{}, error) {
		// build a list of hydrate params objects - one per region
		paramsList := make([]map[string]string, len(regions))
		for i, region := range regions {
			paramsList[i] = map[string]string{"region": region}
		}
		return plugin.GetForPartitions(ctx, queryData, hydrateData, getFunc, paramsList)
	}
}

func multiRegionTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_multi_region",
		Description: "Chaos table to test the multi region features",
		List: &plugin.ListConfig{
			Hydrate: ListRegions(multiRegionList),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    multiRegionGet,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "region", Type: proto.ColumnType_STRING},
		},
	}
}

func multiRegionList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Warn("multiRegionList", "params", h.Params)
	// get region from hydrate params
	region := h.Params["region"]
	for i := 0; i < 5; i++ {
		id := fmt.Sprintf("%d-%s", i, region)
		item := map[string]interface{}{"id": id, "region": region}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func multiRegionGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetInt64Value()
	column1 := fmt.Sprintf("column_1-%d", id)
	column2 := fmt.Sprintf("column_2-%d", id)
	column3 := fmt.Sprintf("column_3-%d", id)

	item := map[string]interface{}{"id": id, "column_1": column1, "column_2": column2, "column_3": column3}
	return item, nil
}
