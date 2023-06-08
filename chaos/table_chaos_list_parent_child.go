package chaos

import (
	"context"
	"errors"
	"time"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

var parentChildListTableErrorCount = 0

func chaosListParentChildTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_parent_child",
		Description: "Chaos table to test the List calls having parent-child dependencies with all the possible scenarios like errors, panics and delays at both parent and child levels",
		List: &plugin.ListConfig{
			ParentHydrate:     listParentRetryTable,
			Hydrate:           getChildList(&listBuildConfig{}),
			ShouldIgnoreError: shouldIgnoreErrorLegacy,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "parent_fatal_error", Type: proto.ColumnType_BOOL, Description: "Column to test the parent list function with fatal error"},
			{Name: "parent_fatal_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the the parent list function with fatal error after streaming some rows"},
			{Name: "parent_retryable_error", Type: proto.ColumnType_BOOL, Description: "Column to test the the parent list function with retry in case of non fatal error"},
			{Name: "parent_retryable_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the the parent list function with retry in case of non fatal errors occured after streaming a few rows"},
			{Name: "parent_should_ignore_error", Type: proto.ColumnType_BOOL, Description: "Column to test the the parent list function with Ignorable errors"},
			{Name: "parent_should_ignore_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the the parent list function with Ignorable errors occuring after already streaming some rows"},
			{Name: "parent_delay", Type: proto.ColumnType_BOOL, Description: "Column to test delay in parent list function"},
			{Name: "parent_panic", Type: proto.ColumnType_BOOL, Description: "Column to test panicking the parent list function"},
			{Name: "child_fatal_error", Type: proto.ColumnType_BOOL, Description: "Column to test the child list function with fatal error"},
			{Name: "child_fatal_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the the child list function with fatal error after streaming some rows"},
			{Name: "child_retryable_error", Type: proto.ColumnType_BOOL, Description: "Column to test the the child list function with retry in case of non fatal error"},
			{Name: "child_retryable_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the the child list function with retry in case of non fatal errors occured after streaming a few rows"},
			{Name: "child_should_ignore_error", Type: proto.ColumnType_BOOL, Description: "Column to test the the child list function with Ignorable errors"},
			{Name: "child_should_ignore_error_after_streaming", Type: proto.ColumnType_BOOL, Description: "Column to test the the child list function with Ignorable errors occuring after already streaming some rows"},
			{Name: "child_delay", Type: proto.ColumnType_BOOL, Description: "Column to test delay in child list function"},
			{Name: "child_panic", Type: proto.ColumnType_BOOL, Description: "Column to test panicking the child list function"},
		},
	}
}

func listParentRetryTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_fatal_error") {
		listBuildConfig := &listBuildConfig{listError: FailError, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_fatal_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: FailError, rowCount: 10, listRowsBeforeError: 2}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_retryable_error") {
		listBuildConfig := &listBuildConfig{listError: RetryableError, rowCount: 10, failureCount: 2}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	// TODO this currently does not actually work - the test only passes as the failure count (15)
	//  is greater than the sdk error retry count
	// https://github.com/turbot/steampipe-plugin-sdk/issues/324
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_retryable_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: RetryableError, rowCount: 10, listRowsBeforeError: 5, failureCount: 15}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_should_ignore_error") {
		listBuildConfig := &listBuildConfig{listError: IgnorableError, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_should_ignore_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: IgnorableError, rowCount: 10, listRowsBeforeError: 2}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_delay") {
		listBuildConfig := &listBuildConfig{listDelay: true, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "parent_panic") {
		listBuildConfig := &listBuildConfig{listError: FailPanic, rowCount: 10}
		return buildListHydrate(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_fatal_error") {
		listBuildConfig := &listBuildConfig{listError: FailError, rowCount: 10}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_fatal_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: FailError, rowCount: 10, listRowsBeforeError: 2}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	// TODO this currently does not actually work - the test only passes as the failure count (15)
	//  is greater than the sdk error retry count
	// https://github.com/turbot/steampipe-plugin-sdk/issues/324
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_retryable_error") {
		listBuildConfig := &listBuildConfig{listError: RetryableError, rowCount: 10, failureCount: 15}
		return getChildList(listBuildConfig)(ctx, d, h)
	}

	if helpers.StringSliceContains(d.QueryContext.Columns, "child_retryable_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: RetryableError, rowCount: 10, failureCount: 5}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_should_ignore_error") {
		listBuildConfig := &listBuildConfig{listError: IgnorableError, rowCount: 10}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_should_ignore_error_after_streaming") {
		listBuildConfig := &listBuildConfig{listError: IgnorableError, rowCount: 10, listRowsBeforeError: 2}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_delay") {
		listBuildConfig := &listBuildConfig{listDelay: true, rowCount: 2}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "child_panic") {
		listBuildConfig := &listBuildConfig{listError: FailPanic, rowCount: 10}
		return getChildList(listBuildConfig)(ctx, d, h)
	}
	return nil, nil
}

func getChildList(listBuildConfig *listBuildConfig) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		// if listDelay is specified, sleep
		if listBuildConfig.listDelay {
			time.Sleep(delayValue)
		}

		// rowCount is the number of rows to return
		rowCount := listBuildConfig.rowCount
		if rowCount == 0 {
			rowCount = defaultRowCount
		}

		for i := 0; i < rowCount; i++ {

			//log.Printf("[DEBUG] ROW LOOP streamed %d error limit %d", i, listBuildConfig.listRowsBeforeError)
			// listErrorRows is the number of rows to return successfully before raising an error
			// if we stream that many rows, let's raise an error
			if i == listBuildConfig.listRowsBeforeError {
				switch listBuildConfig.listError {
				case RetryableError:
					// failureCount is the number of times the error occurs before we succeed
					if parentChildListTableErrorCount <= listBuildConfig.failureCount {
						parentChildListTableErrorCount++
						return nil, errors.New(RetryableError)
					}
					// if we have failed 'failureCount' times, reset listTableErrorCount and fall through to return item
					listTableErrorCount = 0
				case IgnorableError:
					return nil, errors.New(IgnorableError)
				case FailError:
					return nil, errors.New(FatalError)
				case FailPanic:
					panic(FailPanic)

				}

			}
			item := populateItem(i, d.Table)
			d.StreamLeafListItem(ctx, item)

		}
		return nil, nil
	}
}
