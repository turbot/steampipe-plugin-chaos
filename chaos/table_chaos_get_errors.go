package chaos

import (
	"context"
	"errors"
	"time"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const getErrorsRowCount = 100

type getBuildConfig struct {
	getError     FailType
	getDelay     bool
	errorType    FailType
	failureCount int
}

var getTableErrorCount = 0

func chaosGetTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_get_errors",
		Description: "Chaos table to test the Get call with all the possible scenarios like errors, panics and delays",
		List: &plugin.ListConfig{
			Hydrate: listGetErrors,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           chaosGetHydrate,
			ShouldIgnoreError: shouldIgnoreError,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "fatal_error", Type: proto.ColumnType_BOOL, Description: "Column to test the table with fatal error"},
			{Name: "retryable_error", Type: proto.ColumnType_BOOL, Description: "Column to test the Get function with retry config in case of non fatal error"},
			{Name: "ignorable_error", Type: proto.ColumnType_BOOL, Description: "Column to test the  Get function with Ignorable errors"},
			{Name: "delay", Type: proto.ColumnType_BOOL, Description: "Column to test delay in Get function"},
			{Name: "panic", Type: proto.ColumnType_BOOL, Description: "Column to test panicking Get function"},
		},
	}
}

func listGetErrors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < getErrorsRowCount; i++ {
		item := populateItem(i, d.Table)
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func chaosGetHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if helpers.StringSliceContains(d.QueryContext.Columns, "fatal_error") {
		buildConfig := &getBuildConfig{getError: FailError}
		return buildGetHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error") {
		buildConfig := &getBuildConfig{getError: RetryableError, failureCount: 200}
		return buildGetHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "ignorable_error") {
		buildConfig := &getBuildConfig{getError: IgnorableError}
		return buildGetHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "delay") {
		buildConfig := &getBuildConfig{getDelay: true}
		return buildGetHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "panic") {
		buildConfig := &getBuildConfig{getError: FailPanic}
		return buildGetHydrate(buildConfig)(ctx, d, h)
	}
	return nil, nil
}

/// get the 'Get' hydrate function ///
func buildGetHydrate(buildConfig *getBuildConfig) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		id := d.KeyColumnQuals["id"].GetInt64Value()
		if buildConfig.getDelay {
			time.Sleep(delayValue)
		}
		switch buildConfig.getError {
		case RetryableError:
			// failureCount is the number of times the error occurs before we succeed
			if getTableErrorCount <= buildConfig.failureCount {
				getTableErrorCount++
				return nil, errors.New(RetryableError)
			}
			// if we have failed 'failureCount' times, reset getTableErrorCount and fall through to return item
			getTableErrorCount = 0
		case IgnorableError:
			return nil, errors.New(IgnorableError)
		case FailError:
			return nil, errors.New(FatalError)
		case FailPanic:
			panic(FailPanic)
		}
		column := populateItem(int(id), d.Table)
		return column, nil
	}
}
