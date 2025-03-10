package chaos

import (
	"context"
	"log"
	"slices"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func getErrorsDefaultConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_get_errors_default_config",
		Description: "Chaos table to test the Get function using default Retry and Ignore config defined at plugin level",
		List: &plugin.ListConfig{
			Hydrate: listGetErrorsDefaultConfig,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    chaosGetDefaultConfigHydrate,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "retryable_error_default_config", Type: proto.ColumnType_BOOL, Description: "Column to test the Get function with retry config in case of non fatal error"},
			{Name: "ignorable_error_default_config", Type: proto.ColumnType_BOOL, Description: "Column to test the  Get function with Ignorable errors"},
		},
	}
}

func listGetErrorsDefaultConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < getErrorsRowCount; i++ {
		item := populateItem(i, d.Table)
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func chaosGetDefaultConfigHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Printf("[INFO] chaosGetDefaultConfigHydrate")
	if slices.Contains(d.QueryContext.Columns, "retryable_error_default_config") {
		buildConfig := &getBuildConfig{getError: RetryableError, failureCount: 5}
		hy, err := buildGetHydrate(buildConfig)(ctx, d, h)
		return hy, err
	}
	if slices.Contains(d.QueryContext.Columns, "ignorable_error_default_config") {
		buildConfig := &getBuildConfig{getError: IgnorableError}
		return buildGetHydrate(buildConfig)(ctx, d, h)
	}
	return nil, nil
}
