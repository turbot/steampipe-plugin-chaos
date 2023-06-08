package chaos

import (
	"context"
	"errors"
	"log"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func hydrateShouldIgnoreConfigTableWithTableDefaultLegacy() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_should_ignore_config_with_table_default_legacy",
		Description: "Chaos table to test the Hydrate function with Should Ignore Error defined in the Hydrate config, and a table default ignore config",
		List: &plugin.ListConfig{
			Hydrate: hydrateShouldIgnoreConfigList,
		},
		DefaultIgnoreConfig: &plugin.IgnoreConfig{ShouldIgnoreError: shouldIgnoreErrorLegacy},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:              ignorableErrorWithShouldIgnore,
				ShouldIgnoreError: shouldIgnoreErrorLegacy,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:    "ignorable_error_with_ignore_config",
				Type:    proto.ColumnType_STRING,
				Hydrate: ignorableErrorWithShouldIgnore,
				// verify that null hydrate items resulting from ignored errors do not get passed to transform functions
				Transform: transform.From(checkNilTransform),
			},
			{
				Name:    "ignorable_error_with_no_ignore_config",
				Type:    proto.ColumnType_STRING,
				Hydrate: ignorableErrorWithoutShouldIgnore,
				// verify that null hydrate items resulting from ignored errors do not get passed to transform functions
				Transform: transform.From(checkNilTransform),
			},
		},
	}
}

func hydrateShouldIgnoreConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 1; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func ignorableErrorWithShouldIgnore(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Printf("[INFO] ignorableErrorWithShouldIgnore return error")
	return nil, errors.New(notFoundErrorString)
}

func ignorableErrorWithoutShouldIgnore(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Printf("[INFO] ignorableErrorWithShouldIgnore return error")
	return nil, errors.New(notFoundErrorString)
}

func checkNilTransform(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem
	log.Printf("[INFO] checkNilTransform hydrateItem: %v", data)
	if helpers.IsNil(data) {
		panic("NIL HYDRATE")
	}
	return data, nil
}
