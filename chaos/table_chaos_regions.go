package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func regionsTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_regions",
		Description: "Chaos table to print the regions which are present in the connection config, to check parsing functionality.",
		List: &plugin.ListConfig{
			Hydrate: listRegionsFunction,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "region_name", Type: proto.ColumnType_STRING},
		},
	}
}

func listRegionsFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	var config_regions []string
	chaosConfig := GetConfig(d.Connection)

	if chaosConfig.Regions != nil {
		config_regions = chaosConfig.Regions
	}
	all_regions := [5]string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "us-central-1"}

	for a := 0; a < len(config_regions); a++ {
		for b := 0; b < 5; b++ {
			if config_regions[a] == all_regions[b] {
				item := map[string]interface{}{"id": b, "region_name": config_regions[a]}
				d.StreamListItem(ctx, item)
			}
		}
	}
	return nil, nil
}
