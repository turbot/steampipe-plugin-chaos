package chaos

import (
	"context"

	"github.com/turbot/go-kit/helpers"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func buildChaosListTable(tableDef *chaosTable) *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list",
		Description: "Chaos table to test the List calls with all the possible scenarios like errors, panics and delays",
		List: &plugin.ListConfig{
			Hydrate:           chaosListList,
			ShouldIgnoreError: shouldIgnoreError,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "fatal_error", Type: proto.ColumnType_BOOL, Description: "Column to test the table with fatal error"},
			{Name: "fatal_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the table with fatal error after streaming some rows"},
			{Name: "retryable_error", Type: proto.ColumnType_BOOL, Description: "Column to test the test the List function with retry config in case of non fatal error"},
			{Name: "retryable_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the test the List function with retry config in case of non fatal errors occured after streaming a few rows"},
			{Name: "should_ignore_error", Type: proto.ColumnType_BOOL, Description: "Column to test the test the List function with Ignorable errors"},
			{Name: "should_ignore_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the test the List function with Ignorable errors occuring after already streaming some rows"},
			{Name: "delay", Type: proto.ColumnType_BOOL, Description: "Column to test delay in List function"},
			{Name: "panic", Type: proto.ColumnType_BOOL, Description: "Column to test panicking List function"},
			{Name: "panic_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test panicking List function, where function panics after streaming a few rows"},
		},
	}
}

func chaosListList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if helpers.StringSliceContains(d.QueryContext.Columns, "fatal_error") {
		listConfig := &listConfig{listError: FailError}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "fatal_error_after_streaming") {
		listConfig := &listConfig{listError: FailError, rowCount: 15, listErrorRows: 5}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error") {
		listConfig := &listConfig{listError: RetryableError, rowCount: 5, retryCount: 5}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error_after_streaming") {
		listConfig := &listConfig{listError: RetryableError, rowCount: 5, listErrorRows: 5, retryCount: 200}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "should_ignore_error") {
		listConfig := &listConfig{listError: IgnorableError, rowCount: 15}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "should_ignore_error_after_streaming") {
		listConfig := &listConfig{listError: IgnorableError, rowCount: 15, listErrorRows: 5}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "delay") {
		listConfig := &listConfig{listDelay: true}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "panic") {
		listConfig := &listConfig{listError: FailPanic}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "panic_after_streaming") {
		listConfig := &listConfig{listError: FailPanic, rowCount: 15, listErrorRows: 5}
		return getList(listConfig)(ctx, d, h)
	}
	return nil, nil
}
