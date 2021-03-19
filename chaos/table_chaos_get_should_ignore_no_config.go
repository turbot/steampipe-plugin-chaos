package chaos

import (
	"context"
	"errors"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func getShouldIgnoreNoConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_get_should_ignore_no_config",
		Description: "Chaos table to test the Get function with Default Should Ignore Error defined at the plugin level",
		List: &plugin.ListConfig{
			Hydrate: shouldIgnoreListNoConfig,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    shouldIgnoreGetNoConfig,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "should_ignore_error", Type: proto.ColumnType_STRING},
		},
	}
}

func shouldIgnoreListNoConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"id": i, "should_ignore_error": "SUCCESS"}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func shouldIgnoreGetNoConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return nil, errors.New(notFoundError)
}
