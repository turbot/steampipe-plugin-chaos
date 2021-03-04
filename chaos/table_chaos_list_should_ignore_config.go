package chaos

import (
	"context"
	"errors"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func listShouldIgnoreConfigTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_should_ignore_config",
		Description: "Chaos table to test the List function with Should Ignore Error defined in the Hydrate config",
		List: &plugin.ListConfig{
			Hydrate: listShouldIgnoreConfigList,
			ShouldIgnoreError: shouldIgnoreError,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
		},
	}
}

func listShouldIgnoreConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[INFO] INSIDE LIST CALL")

	return nil, errors.New("ResourceNotFound")
}
