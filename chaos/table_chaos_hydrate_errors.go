package chaos

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
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
				Func: chaosHydrateErrorsRetryHydrate,
				RetryConfig: &plugin.RetryConfig{
					ShouldRetryError: shouldRetryErrorLegacy,
				},
			},
			{
				Func:              chaosHydrateErrorsIgnorableHydrate,
				ShouldIgnoreError: shouldIgnoreErrorLegacy,
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{
				Name:        "fatal_error",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test the table with fatal error",
				Hydrate:     chaosHydrateErrorsFatalHydrate,
			},
			{
				Name:        "retryable_error",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test the Hydrate function with retry config in case of non fatal error",
				Hydrate:     chaosHydrateErrorsRetryHydrate,
			},
			{
				Name:        "ignorable_error",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test the  Hydrate function with Ignorable errors",
				Hydrate:     chaosHydrateErrorsIgnorableHydrate,
			},
			{
				Name:        "delay",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test delay in Hydrate function",
				Hydrate:     chaosHydrateErrorsDelayHydrate,
			},
			{
				Name:        "panic",
				Type:        proto.ColumnType_BOOL,
				Description: "Column to test panicking Hydrate function",
				Hydrate:     chaosHydrateErrorsPanicHydrate,
			},
		},
	}
}

func listHydrateErrors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	item := map[string]interface{}{"id": 0}
	d.StreamListItem(ctx, item)
	return nil, nil

}

func chaosHydrateErrorsRetryHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	buildConfig := &hydrateBuildConfig{hydrateError: RetryableError, failureCount: 2}
	return buildHydrate(buildConfig)(ctx, d, h)

}

func chaosHydrateErrorsFatalHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	buildConfig := &hydrateBuildConfig{hydrateError: FailError}
	return buildHydrate(buildConfig)(ctx, d, h)
}

func chaosHydrateErrorsIgnorableHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	buildConfig := &hydrateBuildConfig{hydrateError: IgnorableError}
	return buildHydrate(buildConfig)(ctx, d, h)
}

func chaosHydrateErrorsDelayHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	buildConfig := &hydrateBuildConfig{hydrateDelay: true}
	return buildHydrate(buildConfig)(ctx, d, h)
}

func chaosHydrateErrorsPanicHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	buildConfig := &hydrateBuildConfig{hydrateError: FailPanic}
	return buildHydrate(buildConfig)(ctx, d, h)
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
			log.Printf("[DEBUG] RetryableError error count %d, configured failure count %d", hydrateTableErrorCount, buildConfig.failureCount)
			// failureCount is the number of times the error occurs before we succeed
			if hydrateTableErrorCount < buildConfig.failureCount {
				log.Printf("[DEBUG] return retryable error")
				hydrateTableErrorCount++
				return nil, errors.New(RetryableError)
			}
			log.Printf("[DEBUG] we have failed 'failureCount' times, reset hydrateTableErrorCount and fall through to return item")
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
