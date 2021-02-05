package chaos

import (
	"context"
	"fmt"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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

const fetchMetdataKeyRegion = "region"

func BuildFetchMetadataList() []map[string]interface{} {
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
		FetchMetadata: BuildFetchMetadataList(),

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromFetchMetadata(fetchMetdataKeyRegion)},
		},
	}
}

func regionAwareList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetFetchMetadata(ctx)[fetchMetdataKeyRegion].(string)
	plugin.Logger(ctx).Warn("regionAwareList", "region", region)
	// get region from hydrate params

	for i := 0; i < 5; i++ {
		id := buildId(i, region)
		item := map[string]interface{}{"id": id}
		plugin.Logger(ctx).Warn("regionAwareList", "item", item)
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func regionAwareGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["id"].GetStringValue()
	idRegion := regionFromId(id)
	// get region from context
	region := plugin.GetFetchMetadata(ctx)[fetchMetdataKeyRegion].(string)
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
