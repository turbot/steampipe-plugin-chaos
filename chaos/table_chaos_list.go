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
		Description: "Chaos table to test the List calls with all the possible output",
		List: &plugin.ListConfig{
			Hydrate:           chaosListList,
			ShouldIgnoreError: shouldIgnoreError,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "fatal_error", Type: proto.ColumnType_BOOL},
			{Name: "fatal_error_after_streaming", Type: proto.ColumnType_BOOL},
			{Name: "retryable_error", Type: proto.ColumnType_BOOL},
			{Name: "retryable_error_after_streaming", Type: proto.ColumnType_BOOL},
			{Name: "should_ignore_error", Type: proto.ColumnType_BOOL},
			{Name: "should_ignore_error_after_streaming", Type: proto.ColumnType_BOOL},
			{Name: "delay", Type: proto.ColumnType_BOOL},
			{Name: "panic", Type: proto.ColumnType_BOOL},
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
		listConfig := &listConfig{listError: RetryableError, rowCount: 5, retryCount: 10}
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
	return nil, nil
}
