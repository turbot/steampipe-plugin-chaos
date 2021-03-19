package chaos

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

type listConfig struct {
	listError FailType
	// the number of rows returned before a list error/hydrate error is raised
	listErrorRows int
	listDelay     bool
	rowCount      int
	failureCount  int
}

var listTableErrorCount = 0

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
		listConfig := &listConfig{listError: RetryableError, rowCount: 10, failureCount: 5}
		return getList(listConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error_after_streaming") {
		listConfig := &listConfig{listError: RetryableError, rowCount: 10, listErrorRows: 5, failureCount: 200}
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

// factory function which returns a list call with the behaviour determined by the list config
func getList(listConfig *listConfig) plugin.HydrateFunc {

	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		// if listDelay is specified, sleep
		if listConfig.listDelay {
			time.Sleep(delayValue)
		}

		// rowCount is the number of rows to return
		rowCount := listConfig.rowCount
		if rowCount == 0 {
			rowCount = defaultRowCount
		}

		for i := 0; i < rowCount; i++ {

			log.Printf("[DEBUG] ROW LOOP streamed %d error limit %d", i, listConfig.listErrorRows)
			// listErrorRows is the number of rows to return successfully before raising an error
			// if we stream that many rows, let's raise an error
			if i == listConfig.listErrorRows {
				switch listConfig.listError {
				case RetryableError:
					// failureCount is the number of times the error occurs before we succeed
					if listTableErrorCount <= listConfig.failureCount {
						listTableErrorCount++
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
			d.StreamListItem(ctx, item)

		}
		return nil, nil
	}
}
