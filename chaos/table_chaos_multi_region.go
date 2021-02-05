package chaos

import (
	"context"
	"fmt"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const fetchMetdataKeyRegion = "region"

func GetRegions(_ context.Context, config interface{}) []map[string]interface{} {
	chaosConfig := config.(chaosConfig)
	regions := chaosConfig.Regions
	// build a list of fetchMetadata - one per region
	fetchMetadataList := make([]map[string]interface{}, len(regions))
	for i, region := range regions {
		fetchMetadataList[i] = map[string]interface{}{fetchMetdataKeyRegion: region}
	}
	return fetchMetadataList
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
		FetchMetadata: GetRegions,

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromFetchMetadata(fetchMetdataKeyRegion)},
		},
	}
}

func regionAwareList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := regionFromFetchMetadata(ctx)
	if region == "" {
		return nil, nil
	}

	plugin.Logger(ctx).Warn("regionAwareList", "region", region)

	for i := 0; i < 5; i++ {
		id := buildId(i, region)
		item := map[string]interface{}{"id": id}
		plugin.Logger(ctx).Warn("regionAwareList", "item", item)
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func regionFromFetchMetadata(ctx context.Context) string {
	fetchMetadata := plugin.GetFetchMetadata(ctx)
	if fetchMetadata == nil {
		return ""
	}
	return fetchMetadata[fetchMetdataKeyRegion].(string)
}

func regionAwareGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetStringValue()
	idRegion := regionFromId(id)
	region := regionFromFetchMetadata(ctx)
	if region == "" {
		return nil, nil
	}

	//plugin.Logger(ctx).Warn("regionAwareGet", "region", region)

	if region == idRegion {
		plugin.Logger(ctx).Warn("****************** regionAwareGet - match!", "id", id, "region", region, "idRegion", idRegion)
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
