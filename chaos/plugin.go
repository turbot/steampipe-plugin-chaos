package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const pluginName = "steampipe-provider-chaos"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		DefaultGetConfig: &plugin.GetConfig{
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryError: shouldRetryError,
			},
			ShouldIgnoreError: shouldIgnoreError,
		},

		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency:   500,
			DefaultMaxConcurrency: 150,
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryError: shouldRetryError,
		},
		TableMap: map[string]*plugin.Table{
			"chaos_high_row_count":               buildTable(&chaosTable{name: "chaos_high_row_count", description: "Chaos table to test steampipe with high row count", listConfig: &listConfig{rowCount: 10}}),
			"chaos_high_column_count":            buildTable(&chaosTable{name: "chaos_high_column_count", description: "Chaos table to test steampipe with high column count", columnCount: 100}),
			"chaos_list":                         buildChaosListTable(&chaosTable{listConfig: &listConfig{}}),
			"chaos_list_parent_child":            listRetryTableBuild(&chaosTable{listConfig: &listConfig{}}),
			"chaos_get_error":                    buildTable(&chaosTable{name: "chaos_get_error", description: "Chaos table to test error handling in Get function", getConfig: &getConfig{getError: FailError}}),
			"chaos_get_panic":                    buildTable(&chaosTable{name: "chaos_get_panic", description: "Chaos table to test panicking Get function", getConfig: &getConfig{getError: FailPanic}}),
			"chaos_get_delay":                    buildTable(&chaosTable{name: "chaos_get_delay", description: "Chaos table to test delay in in Get function", getConfig: &getConfig{getDelay: true}}),
			"chaos_hydrate_error":                buildTable(&chaosTable{name: "chaos_hydrate_error", description: "Chaos table to test error handling in Hydrate function", hydrateError: FailError, listConfig: &listConfig{rowCount: 15, listErrorRows: 2}}),
			"chaos_hydrate_panic":                buildTable(&chaosTable{name: "chaos_hydrate_panic", description: "Chaos table to test panicking Hydrate function", hydrateError: FailPanic}),
			"chaos_hydrate_delay":                buildTable(&chaosTable{name: "chaos_hydrate_delay", description: "Chaos table to test delay in in Hydrate function", hydrateDelay: true}),
			"chaos_transform_error":              buildTable(&chaosTable{name: "chaos_transform_error", description: "Chaos table to test error handling in Transform function", transformError: FailError}),
			"chaos_transform_panic":              buildTable(&chaosTable{name: "chaos_transform_panic", description: "Chaos table to test panicking Transform function", transformError: FailPanic}),
			"chaos_transform_delay":              buildTable(&chaosTable{name: "chaos_transform_delay", description: "Chaos table to test delay in in Transform function", transformDelay: true}),
			"chaos_get_test":                     getTestTable(),
			"chaos_multi_region":                 multiRegionTable(),
			"chaos_all_column_types":             allColumnsTable(),
			"chaos_hydrate_columns_dependency":   hydrateColumnsTable(),
			"chaos_parent_child_dependency":      parentChildTable(),
			"chaos_default_transform":            defaultTransformTable(),
			"chaos_parallel_hydrate_columns":     parallelHydrateColumnsTable(),
			"chaos_all_numeric_column":           numericColumnsTable(),
			"chaos_transform_method_test":        transformMethodTable(),
			"chaos_parallel_hydrate_test":        getTestParallelismTable(),
			"chaos_concurrency_limit":            getConcurrencyLimitTable(),
			"chaos_concurrency_no_limit":         getConcurrencyNoLimitTable(),
			"chaos_hydrate_retry_config":         hydrateRetryConfigTable(),
			"chaos_hydrate_retry_no_config":      hydrateRetryNoConfigTable(),
			"chaos_hydrate_should_ignore_config": hydrateShouldIgnoreConfigTable(),
			"chaos_get_retry_config":             getRetryConfigTable(),
			"chaos_get_retry_no_config":          getRetryNoConfigTable(),
			"chaos_get_should_ignore_config":     getShouldIgnoreConfigTable(),
			"chaos_get_should_ignore_no_config":  getShouldIgnoreNoConfigTable(),
			"chaos_list_paging":                  listPagingTable(),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
	}

	return p
}
