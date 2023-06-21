package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const pluginName = "steampipe-provider-chaos"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		DefaultGetConfig: &plugin.GetConfig{
			RetryConfig:  &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError},
			IgnoreConfig: &plugin.IgnoreConfig{ShouldIgnoreErrorFunc: shouldIgnoreError},
		},

		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency:   500,
			DefaultMaxConcurrency: 150,
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryError,
		},
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreError,
		},
		TableMap: map[string]*plugin.Table{
			"chaos_very_high_row_count":                                    buildTable(&chaosTable{name: "chaos_very_high_row_count", description: "Chaos table to test steampipe with very high row count", listBuildConfig: &listBuildConfig{rowCount: 1000000}, cache: &plugin.TableCacheOptions{Enabled: true}}),
			"chaos_high_row_count":                                         buildTable(&chaosTable{name: "chaos_high_row_count", description: "Chaos table to test steampipe with high row count", listBuildConfig: &listBuildConfig{rowCount: 5000}, cache: &plugin.TableCacheOptions{Enabled: true}}),
			"chaos_high_column_count":                                      buildTable(&chaosTable{name: "chaos_high_column_count", description: "Chaos table to test steampipe with high column count", listBuildConfig: &listBuildConfig{columnCount: 100}}),
			"chaos_list_errors":                                            chaosListTable(),                                       // test the List calls with all the possible scenarios like errors, panics and delays
			"chaos_list_parent_child":                                      chaosListParentChildTable(),                            // test the List calls having parent-child dependencies with all the possible scenarios like errors, panics and delays at both parent and child levels
			"chaos_get_errors":                                             chaosGetTable(),                                        // test the Get call with all the possible scenarios like errors, panics and delays
			"chaos_get_errors_default_config":                              getErrorsDefaultConfigTable(),                          // test the Get function with Default Retry config defined at plugin level in case of non fatal error
			"chaos_multi_region":                                           multiRegionTable(),                                     // test the multi region features
			"chaos_all_column_types":                                       allColumnsTable(),                                      // test all columns of different types
			"chaos_hydrate_columns_dependency":                             hydrateColumnsDependencyTable(),                        // test dependencies between hydrate functions
			"chaos_parallel_hydrate_columns":                               parallelHydrateColumnsTable(),                          // test the execution of multiple hydrate functions and transform functions asynchronously
			"chaos_all_numeric_column":                                     numericColumnsTable(),                                  // test all flavours of integer and float data types
			"chaos_cache_check":                                            checkCacheTable(),                                      // test the cache functionality
			"chaos_cache_with_delay_quals":                                 cacheDelayQualsTable(),                                 // test cache
			"chaos_concurrency_limit":                                      getConcurrencyLimitTable(),                             // test the concurrency limit of hydrate calls
			"chaos_concurrency_no_limit":                                   getConcurrencyNoLimitTable(),                           // test high concurrency and no limit (apart from the plugin level limit)
			"chaos_hydrate_retry_config":                                   hydrateRetryConfigTable(),                              // test the Hydrate function with Retry config in case of non fatal error
			"chaos_hydrate_retry_config_legacy":                            hydrateRetryConfigLegacyTable(),                        // test the Hydrate function with Retry config in case of non fatal error
			"chaos_hydrate_retry_config_with_table_default":                hydrateRetryConfigWithTableDefaultTable(),              // test the Hydrate function with Retry config in case of non fatal error
			"chaos_hydrate_retry_config_with_table_default_legacy":         hydrateRetryConfigWithTableDefaultLegacyTable(),        // test the Hydrate function with Retry config in case of non fatal error
			"chaos_hydrate_should_ignore_config_with_table_default_legacy": hydrateShouldIgnoreConfigTableWithTableDefaultLegacy(), // test the Hydrate function with Should Ignore Error defined in the Hydrate config
			"chaos_hydrate_should_ignore_config_legacy":                    hydrateShouldIgnoreConfigTableLegacy(),                 // test the Hydrate function with Should Ignore Error defined in the Hydrate config
			"chaos_hydrate_should_ignore_config_with_table_default":        hydrateShouldIgnoreConfigTableWithTableDefault(),       // test the Hydrate function with Should Ignore Error defined in the Hydrate config
			"chaos_hydrate_should_ignore_config":                           hydrateShouldIgnoreConfigTable(),                       // test the Hydrate function with Should Ignore Error defined in the Hydrate config
			"chaos_limit":                                                  chaosLimitTable(),                                      // test the limit functionality
			"chaos_list_paging":                                            listPagingTable(),                                      // test the list function supporting pagination fails to send results after some pages with a non fatal error
			"chaos_list_single_key_columns":                                KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_list_single_key_columns", description: "Chaos table to test the list call with single key column", listKeyColumnSetType: "single"}),
			"chaos_list_all_key_columns":                                   KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_list_all_key_columns", description: "Chaos table to test the list call with all key columns", listKeyColumnSetType: "all"}),
			"chaos_list_any_key_columns":                                   KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_list_any_key_columns", description: "Chaos table to test the list call with any of the specified key columns", listKeyColumnSetType: "any"}),
			"chaos_get_single_key_columns":                                 KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_get_single_key_columns", description: "Chaos table to test the get call with single key column", getKeyColumnSetType: "single"}),
			"chaos_get_all_key_columns":                                    KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_get_all_key_columns", description: "Chaos table to test the get call with all key columns", getKeyColumnSetType: "all"}),
			"chaos_get_any_key_columns":                                    KeyColumnTableBuilder(&keyColumnTableDefinitions{name: "chaos_get_any_key_columns", description: "Chaos table to test the get call with any of the specified key columns", getKeyColumnSetType: "any"}),
			"chaos_hydrate_errors":                                         chaosHydrateTable(),                      // test the Hydrate call with all the possible scenarios like errors, panics and delays
			"chaos_plugin_crash":                                           pluginCrashTable(),                       // test the plugin crashing issue
			"chaos_regions":                                                regionsTable(),                           // test the connection config parsing
			"chaos_transform_errors":                                       chaosTransformTable(),                    // test the Transform functions with all the possible scenarios like errors, panics and delays
			"chaos_transforms":                                             transformFromFunctionsTable(),            // test all the From Transforms functions
			"chaos_list_key_column_single_equal":                           listKeyColumnsSingleEqualsTable(),        // test the Hydrate call with all the possible scenarios like errors, panics and delays
			"chaos_list_key_columns_any_multiple_operator":                 listKeyColumnsAnyMultipleOperatorTable(), // test the Hydrate call with all the possible scenarios like errors, panics and delays
			"chaos_list_key_columns_all_multiple_operator":                 listKeyColumnsAllMultipleOperatorTable(), // test the Hydrate call with all the possible scenarios like errors, panics and delays
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
	}

	return p
}
