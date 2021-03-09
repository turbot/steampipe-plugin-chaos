package chaos

import (
	"context"
	"errors"
	log "log"
	"time"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listRetryTableBuild(tableDef *chaosTable) *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_parent_child",
		Description: "Chaos table to test the List calls with all the possible output",
		List: &plugin.ListConfig{
			ParentHydrate: listParentRetryTable,
			Hydrate:       getChildList(&listConfig{}),
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "parent_fatal_error", Type: proto.ColumnType_BOOL},
			{Name: "parent_fatal_error_after_streaming", Type: proto.ColumnType_BOOL},
			{Name: "parent_retryable_error", Type: proto.ColumnType_BOOL},
			{Name: "parent_retryable_error_after_streaming", Type: proto.ColumnType_BOOL},
			{Name: "child_fatal_error", Type: proto.ColumnType_BOOL},
			{Name: "child_fatal_error_after_streaming", Type: proto.ColumnType_BOOL},
			{Name: "child_retryable_error", Type: proto.ColumnType_BOOL},
			{Name: "child_retryable_error_after_streaming", Type: proto.ColumnType_BOOL},
		},
	}

}

func listParentRetryTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_fatal_error") {
		listConfig := &listConfig{listError: FailError}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_fatal_error_after_streaming") {
		listConfig := &listConfig{listError: FailError, rowCount: 15, listErrorRows: 5}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_retryable_error") {
		listConfig := &listConfig{listError: RetryableError}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_retryable_error_after_streaming") {
		listConfig := &listConfig{listError: RetryableError, rowCount: 15, listErrorRows: 5}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_fatal_error") {
		listConfig := &listConfig{listError: FailError}
		return getChildList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_fatal_error_after_streaming") {
		listConfig := &listConfig{listError: FailError, rowCount: 15, listErrorRows: 5}
		return getChildList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_retryable_error") {
		listConfig := &listConfig{listError: RetryableError}
		return getChildList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_retryable_error_after_streaming") {
		listConfig := &listConfig{listError: RetryableError, rowCount: 15, listErrorRows: 5}
		return getChildList(listConfig)(ctx, d, h)
	}
	return nil, nil
}

func getChildList(listConfig *listConfig) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		log.Printf("[DEBUG] INSIDE LIST CALL")
		if listConfig.listDelay {
			time.Sleep(delayValue)
		}

		rowCount := listConfig.rowCount
		if rowCount == 0 {
			rowCount = defaultRowCount
		}
		rowsStreamed := 0

		for i := 0; i < rowCount; i++ {

			log.Printf("[DEBUG] ROW LOOP streamed %d error limit %d", rowsStreamed, listConfig.listErrorRows)
			if rowsStreamed >= listConfig.listErrorRows {
				if listConfig.listError == RetryableError {
					return nil, errors.New(RetryableError)
				}
				if listConfig.listError == IgnorableError {
					return nil, errors.New(IgnorableError)
				}
				if listConfig.listError == FailError {
					log.Printf("[DEBUG] LIST ERROR ")
					return nil, errors.New(FatalError)
				}
				if listConfig.listError == FailPanic {
					panic(FailPanic)
				}
			}

			item := populateItem(i, d.Table)
			d.StreamLeafListItem(ctx, item)

		}
		return nil, nil
	}
}
