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
			"chaos_high_row_count":               buildTable(&chaosTable{name: "chaos_high_row_count", description: "Chaos table to test steampipe with high row count", listBuildConfig: &listBuildConfig{rowCount: 1000}}),
			"chaos_high_column_count":            buildTable(&chaosTable{name: "chaos_high_column_count", description: "Chaos table to test steampipe with high column count", listBuildConfig: &listBuildConfig{columnCount: 100}}),
			"chaos_list_errors":                  chaosListTable(),            // test the List calls with all the possible scenarios like errors, panics and delays
			"chaos_list_parent_child":            chaosListParentChildTable(), // test the List calls having parent-child dependencies with all the possible scenarios like errors, panics and delays at both parent and child levels
			"chaos_transform_error":              buildTable(&chaosTable{name: "chaos_transform_error", description: "Chaos table to test error handling in Transform function", transformError: FailError}),
			"chaos_transform_panic":              buildTable(&chaosTable{name: "chaos_transform_panic", description: "Chaos table to test panicking Transform function", transformError: FailPanic}),
			"chaos_transform_delay":              buildTable(&chaosTable{name: "chaos_transform_delay", description: "Chaos table to test delay in in Transform function", transformDelay: true}),
			"chaos_get_errors":                   chaosGetTable(),                  // test the Get call with all the possible scenarios like errors, panics and delays
			"chaos_multi_region":                 multiRegionTable(),               // test the multi region features
			"chaos_all_column_types":             allColumnsTable(),                // test all columns of different types
			"chaos_hydrate_columns_dependency":   hydrateColumnsDependencyTable(),  // test dependencies between hydrate functions
			"chaos_json_tag_transform":           jsonTagTransformTable(),          // test the default transform functionality from specified json tags
			"chaos_parallel_hydrate_columns":     parallelHydrateColumnsTable(),    // test the execution of multiple hydrate functions and transform functions asynchronously
			"chaos_all_numeric_column":           numericColumnsTable(),            // test all flavours of integer and float data types
			"chaos_transform_from_method":        transformFromMethodTable(),       // test the transform FromMethod invoking a function on the hydrate item
			"chaos_concurrency_limit":            getConcurrencyLimitTable(),       // test the concurrency limit of hydrate calls
			"chaos_concurrency_no_limit":         getConcurrencyNoLimitTable(),     // test high concurrency and no limit (apart from the plugin level limit)
			"chaos_hydrate_retry_config":         hydrateRetryConfigTable(),        // test the Hydrate function with Retry config in case of non fatal error
			"chaos_hydrate_retry_no_config":      hydrateRetryNoConfigTable(),      // test the Hydrate function with Default Retry config defined at plugin level in case of non fatal error
			"chaos_hydrate_should_ignore_config": hydrateShouldIgnoreConfigTable(), // test the Hydrate function with Should Ignore Error defined in the Hydrate config
			"chaos_get_errors_default_config":    getErrorsDefaultConfigTable(),    // test the Get function with Default Retry config defined at plugin level in case of non fatal error
			"chaos_list_paging":                  listPagingTable(),                // test the list function supporting pagination fails to send results after some pages with a non fatal error
			"chaos_transform_from_fields":        transformFromFieldsTable(),
			"chaos_list_single_key_columns":      KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_list_single_key_columns", description: "Chaos table to test the list call with single key column", call: listCall, keyColumnSetType: "single"}),
			"chaos_list_all_key_columns":         KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_list_all_key_columns", description: "Chaos table to test the list call with all key columns", call: listCall, keyColumnSetType: "all"}),
			"chaos_list_any_key_columns":         KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_list_any_key_columns", description: "Chaos table to test the list call with any of the specified key columns", call: listCall, keyColumnSetType: "any"}),
			"chaos_get_single_key_columns":       KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_get_single_key_columns", description: "Chaos table to test the get call with single key column", call: getCall, keyColumnSetType: "single"}),
			"chaos_get_all_key_columns":          KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_get_all_key_columns", description: "Chaos table to test the get call with all key columns", call: getCall, keyColumnSetType: "all"}),
			"chaos_get_any_key_columns":          KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_get_any_key_columns", description: "Chaos table to test the get call with any of the specified key columns", call: getCall, keyColumnSetType: "any"}),
			"chaos_hydrate_errors":               chaosHydrateTable(), // test the Hydrate call with all the possible scenarios like errors, panics and delays
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
	}

	return p
}
