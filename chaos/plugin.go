package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const pluginName = "steampipe-provider-chaos"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency:   500,
			DefaultMaxConcurrency: 150,
		},
		TableMap: map[string]*plugin.Table{
			"chaos_high_row_count":             buildTable(&chaosTable{name: "chaos_high_row_count", description: "Chaos table to test steampipe with high row count", rowCount: 10}),
			"chaos_high_column_count":          buildTable(&chaosTable{name: "chaos_high_column_count", description: "Chaos table to test steampipe with high column count", columnCount: 100}),
			"chaos_list_error":                 buildTable(&chaosTable{name: "chaos_list_error", description: "Chaos table to test error handling in List function", listError: FailError, rowCount: 15, listErrorRows: 10}),
			"chaos_list_panic":                 buildTable(&chaosTable{name: "chaos_list_panic", description: "Chaos table to test panicking List function", listError: FailPanic}),
			"chaos_list_delay":                 buildTable(&chaosTable{name: "chaos_list_delay", description: "Chaos table to test delay in List function", listDelay: true}),
			"chaos_get_error":                  buildTable(&chaosTable{name: "chaos_get_error", description: "Chaos table to test error handling in Get function", getError: FailError}),
			"chaos_get_panic":                  buildTable(&chaosTable{name: "chaos_get_panic", description: "Chaos table to test panicking Get function", getError: FailPanic}),
			"chaos_get_delay":                  buildTable(&chaosTable{name: "chaos_get_delay", description: "Chaos table to test delay in in Get function", getDelay: true}),
			"chaos_hydrate_error":              buildTable(&chaosTable{name: "chaos_hydrate_error", description: "Chaos table to test error handling in Hydrate function", hydrateError: FailError, rowCount: 15, listErrorRows: 2}),
			"chaos_hydrate_panic":              buildTable(&chaosTable{name: "chaos_hydrate_panic", description: "Chaos table to test panicking Hydrate function", hydrateError: FailPanic}),
			"chaos_hydrate_delay":              buildTable(&chaosTable{name: "chaos_hydrate_delay", description: "Chaos table to test delay in in Hydrate function", hydrateDelay: true}),
			"chaos_transform_error":            buildTable(&chaosTable{name: "chaos_transform_error", description: "Chaos table to test error handling in Transform function", transformError: FailError}),
			"chaos_transform_panic":            buildTable(&chaosTable{name: "chaos_transform_panic", description: "Chaos table to test panicking Transform function", transformError: FailPanic}),
			"chaos_transform_delay":            buildTable(&chaosTable{name: "chaos_transform_delay", description: "Chaos table to test delay in in Transform function", transformDelay: true}),
			"chaos_get_test":                   getTestTable(),
			"chaos_multi_region":               multiRegionTable(),
			"chaos_all_column_types":           allColumnsTable(),
			"chaos_hydrate_columns_dependency": hydrateColumnsTable(),
			"chaos_parent_child_dependency":    parentChildTable(),
			"chaos_default_transform":          defaultTransformTable(),
			"chaos_parallel_hydrate_columns":   parallelHydrateColumnsTable(),
			"chaos_all_numeric_column":         numericColumnsTable(),
			"chaos_transform_method_test":      transformMethodTable(),
			"chaos_parallel_hydrate_test":      getTestParallelismTable(),
			"chaos_concurrency_limit":          getConcurrencyLimitTable(),
			"chaos_concurrency_no_limit":       getConcurrencyNoLimitTable(),
			"chaos_list_paging":                listPagingTable(),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
	}

	return p
}
