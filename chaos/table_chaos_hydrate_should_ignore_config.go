package chaos

import (
	"context"
	"errors"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func hydrateShouldIgnoreConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_should_ignore_config",
		Description: "Chaos table to test the Hydrate function with Should Ignore Error defined in the Hydrate config",
		List: &plugin.ListConfig{
			Hydrate: hydrateShouldIgnoreConfigList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:              shouldIgnoreConfigHydrateA,
				ShouldIgnoreError: shouldIgnoreError,
			},
			{
				Func:              shouldIgnoreConfigHydrateB,
				ShouldIgnoreError: shouldIgnoreError,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:    "should_ignore_error_a",
				Type:    proto.ColumnType_STRING,
				Hydrate: shouldIgnoreConfigHydrateA,
			},
			{
				Name:    "should_ignore_error_b",
				Type:    proto.ColumnType_STRING,
				Hydrate: shouldIgnoreConfigHydrateB,
			},
		},
	}
}

func hydrateShouldIgnoreConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func shouldIgnoreConfigHydrateA(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE HYDRATE CALL")

	return nil, errors.New(notFoundError)
}

func shouldIgnoreConfigHydrateB(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE HYDRATE CALL")

	return nil, errors.New(notFoundError)
}
