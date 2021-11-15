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

type hydrateBuildConfig struct {
	hydrateError FailType
	hydrateDelay bool
	errorType    FailType
	failureCount int
}

var hydrateTableErrorCount = 0

func chaosHydrateTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_errors",
		Description: "Chaos table to test the Hydrate call with all the possible scenarios like errors, panics and delays",
		List: &plugin.ListConfig{
			Hydrate: listHydrateErrors,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: chaosHydrateFunction,
				RetryConfig: &plugin.RetryConfig{
					ShouldRetryError: shouldRetryError,
				},
				ShouldIgnoreError: shouldIgnoreError,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{
				Name:        "fatal_error",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test the table with fatal error",
				Hydrate:     chaosHydrateFunction,
			},
			{
				Name:        "retryable_error",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test the Hydrate function with retry config in case of non fatal error",
				Hydrate:     chaosHydrateFunction,
			},
			{
				Name:        "ignorable_error",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test the  Hydrate function with Ignorable errors",
				Hydrate:     chaosHydrateFunction,
			},
			{
				Name:        "delay",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test delay in Hydrate function",
				Hydrate:     chaosHydrateFunction,
			},
			{
				Name:        "panic",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test panicking Hydrate function",
				Hydrate:     chaosHydrateFunction,
			},
		},
	}
}

func listHydrateErrors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 10; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

func chaosHydrateFunction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if helpers.StringSliceContains(d.QueryContext.Columns, "fatal_error") {
		buildConfig := &hydrateBuildConfig{hydrateError: FailError}
		return buildHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "retryable_error") {
		buildConfig := &hydrateBuildConfig{hydrateError: RetryableError, failureCount: 5}
		return buildHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "ignorable_error") {
		buildConfig := &hydrateBuildConfig{hydrateError: IgnorableError}
		return buildHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "delay") {
		buildConfig := &hydrateBuildConfig{hydrateDelay: true}
		return buildHydrate(buildConfig)(ctx, d, h)
	}
	if helpers.StringSliceContains(d.QueryContext.Columns, "panic") {
		buildConfig := &hydrateBuildConfig{hydrateError: FailPanic}
		return buildHydrate(buildConfig)(ctx, d, h)
	}
	return nil, nil
}

func buildHydrate(buildConfig *hydrateBuildConfig) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		time.Sleep(delayValue)
		key := h.Item.(map[string]interface{})
		id := key["id"].(int)
		if buildConfig.hydrateDelay {
			time.Sleep(delayValue)
		}
		if buildConfig.hydrateError == RetryableError {
			log.Printf("[DEBUG] RetryableError")
			// failureCount is the number of times the error occurs before we succeed
			if hydrateTableErrorCount <= buildConfig.failureCount {
				hydrateTableErrorCount++
				return nil, errors.New(RetryableError)
			}
			// if we have failed 'failureCount' times, reset hydrateTableErrorCount and fall through to return item
			hydrateTableErrorCount = 0

		}
		if buildConfig.hydrateError == IgnorableError {
			log.Printf("[DEBUG] IgnorableError")
			return nil, errors.New(IgnorableError)
		}
		if buildConfig.hydrateError == FailError {
			log.Printf("[DEBUG] FatalError")
			return nil, errors.New(FatalError)
		}
		if buildConfig.hydrateError == FailPanic {
			log.Printf("[DEBUG] FailPanic")
			panic(FailPanic)
		}

		item := populateItem(id, d.Table)
		log.Printf("[DEBUG] RETURN ITEM")
		return item, nil
	}
}
