package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func hydrateShouldIgnoreConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_hydrate_should_ignore_config",
		Description: "Chaos table to test the Hydrate function with Should Ignore Error defined in the Hydrate config",
		List: &plugin.ListConfig{
			Hydrate: hydrateShouldIgnoreConfigList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:         ignorableErrorWithShouldIgnore,
				IgnoreConfig: &plugin.IgnoreConfig{ShouldIgnoreErrorFunc: shouldIgnoreError},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:    "ignorable_error_with_ignore_config",
				Type:    proto.ColumnType_STRING,
				Hydrate: ignorableErrorWithShouldIgnore,
				// verify that null hydrate items resulting from ignored errors do not get passed to transform functions
				Transform: transform.From(checkNilTransform),
			},
			{
				Name:    "ignorable_error_with_no_ignore_config",
				Type:    proto.ColumnType_STRING,
				Hydrate: ignorableErrorWithoutShouldIgnore,
				// verify that null hydrate items resulting from ignored errors do not get passed to transform functions
				Transform: transform.From(checkNilTransform),
			},
		},
	}
}
