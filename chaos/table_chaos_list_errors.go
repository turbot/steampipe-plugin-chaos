package chaos

import (
	"context"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type listBuildConfig struct {
	listError FailType
	// the number of rows returned before a list error/hydrate error is raised
	listErrorRows int
	listDelay     bool
	rowCount      int
	columnCount   int
	failureCount  int
}

var listTableErrorCount = 0

func chaosListTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_errors",
		Description: "Chaos table to test the List calls with all the possible scenarios like errors, panics and delays",
		List: &plugin.ListConfig{
			Hydrate:           chaosListHydrate,
			ShouldIgnoreError: shouldIgnoreError,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "fatal_error", Type: proto.ColumnType_BOOL, Description: "Column to test the table with fatal error"},
			{Name: "fatal_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the table with fatal error after streaming some rows"},
			{Name: "retryable_error", Type: proto.ColumnType_BOOL, Description: "Column to test the List function with retry config in case of non fatal error"},
			{Name: "retryable_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the List function with retry config in case of non fatal errors occured after streaming a few rows"},
			{Name: "should_ignore_error", Type: proto.ColumnType_BOOL, Description: "Column to test the List function with Ignorable errors"},
			{Name: "should_ignore_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the List function with Ignorable errors occuring after already streaming some rows"},
			{Name: "delay", Type: proto.ColumnType_BOOL, Description: "Column to test delay in List function"},
			{Name: "panic", Type: proto.ColumnType_BOOL, Description: "Column to test panicking List function"},
			{Name: "panic_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test panicking List function, where function panics after streaming a few rows"},
		},
	}
}

func chaosListHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if helpers.StringSliceContains(d.QueryContext.Columns, "fatal_error") {
		listBuildConfig := &listBuildConfig{listError: FailError, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "fatal_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: FailError, rowCount: 15, listErrorRows: 5}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error") {
		listBuildConfig := &listBuildConfig{listError: RetryableError, rowCount: 10, failureCount: 5}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: RetryableError, rowCount: 10, listErrorRows: 5, failureCount: 200}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "should_ignore_error") {
		listBuildConfig := &listBuildConfig{listError: IgnorableError, rowCount: 15}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "should_ignore_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: IgnorableError, rowCount: 15, listErrorRows: 5}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "delay") {
		listBuildConfig := &listBuildConfig{listDelay: true, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "panic") {
		listBuildConfig := &listBuildConfig{listError: FailPanic, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "panic_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: FailPanic, rowCount: 15, listErrorRows: 5}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	return nil, nil
}
