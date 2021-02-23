package chaos

import (
	"context"
	"errors"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func getShouldIgnoreConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_get_should_ignore_config",
		Description: "Chaos table to test the Get function with Should Ignore Error defined in the Get config",
		List: &plugin.ListConfig{
			Hydrate: shouldIgnoreListConfig,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           shouldIgnoreGetConfig,
			ShouldIgnoreError: shouldIgnoreError,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "should_ignore_error", Type: proto.ColumnType_STRING},
		},
	}
}

func shouldIgnoreListConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 5; i++ {
		item := map[string]interface{}{"id": i, "should_ignore_error": "SUCCESS"}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func shouldIgnoreGetConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE GET CALL")

	return nil, errors.New("ResourceNotFound")
}
