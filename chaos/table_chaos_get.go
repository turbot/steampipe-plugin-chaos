package chaos

// import (
// 	"context"

// 	"github.com/turbot/go-kit/helpers"

// 	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
// 	"github.com/turbot/steampipe-plugin-sdk/plugin"
// )

// func chaosGetTable() *plugin.Table {
// 	return &plugin.Table{
// 		Name:        "chaos_get",
// 		Description: "Chaos table to test the List Error handling",
// 		List: &plugin.ListConfig{
// 			Hydrate: chaosGetList,
// 		},
// 		Get: &plugin.GetConfig{
// 			KeyColumns: plugin.SingleColumn("id"),
// 			Hydrate:    chaosGetGet,
// 		},

// 		Columns: []*plugin.Column{
// 			{Name: "id", Type: proto.ColumnType_INT},
// 			{Name: "get_fatal_error", Type: proto.ColumnType_BOOL},
// 			{Name: "get_retryable_error", Type: proto.ColumnType_BOOL},
// 			{Name: "get_delay", Type: proto.ColumnType_BOOL},
// 			{Name: "get_panic", Type: proto.ColumnType_BOOL},
// 			{Name: "get_ignore_error", Type: proto.ColumnType_BOOL},
// 		},
// 	}
// }

// func chaosGetGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	if helpers.StringSliceContains(d.QueryContext.Columns, "get_fatal_error") {
// 		getConfig := &getConfig{getError: FatalError}
// 		return getGet(getConfig)(ctx, d, h)
// 	}
// 	if helpers.StringSliceContains(d.QueryContext.Columns, "get_retryable_error") {
// 		getConfig := &getConfig{getError: RetryableError}
// 		return getGet(getConfig)(ctx, d, h)
// 	}
// 	if helpers.StringSliceContains(d.QueryContext.Columns, "get_delay") {
// 		getConfig := &getConfig{getError: RetryableError}
// 		return getGet(getConfig)(ctx, d, h)
// 	}
// 	if helpers.StringSliceContains(d.QueryContext.Columns, "get_panic") {
// 		getConfig := &getConfig{getError: RetryableError}
// 		return getGet(getConfig)(ctx, d, h)
// 	}
// 	return nil, nil
// }
