package chaos

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func pluginCrashTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_plugin_crash",
		Description: "Chaos table to print 50 rows and do an os.Exit(-1) to simulate a plugin crash",
		List: &plugin.ListConfig{
			Hydrate: listIdsFunction,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
		},
	}
}

func listIdsFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 50; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
		os.Exit(-1)
	}

	return nil, nil
}
