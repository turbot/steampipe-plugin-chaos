package chaos

import (
	"context"
	"errors"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type transformBuildConfig struct {
	transformError FailType
	transformDelay bool
}

func chaosTransformTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_transform_errors",
		Description: "Chaos table to test the Transform call with all the possible scenarios like errors, panics and delays",
		List: &plugin.ListConfig{
			Hydrate: listTransformErrors,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "error", Type: proto.ColumnType_BOOL, Description: "Column to test the Transform function with fatal error", Transform: transform.From(buildTransform(&transformBuildConfig{transformError: FailError}))},
			{Name: "delay", Type: proto.ColumnType_INT, Description: "Column to test delay in Transform function", Transform: transform.From(buildTransform(&transformBuildConfig{transformDelay: true}))},
			{Name: "panic", Type: proto.ColumnType_BOOL, Description: "Column to test panicking Transform function", Transform: transform.From(buildTransform(&transformBuildConfig{transformError: FailPanic}))},
		},
	}
}

func listTransformErrors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		item := populateItem(i, d.Table)
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}

// // Transform functions ////
func buildTransform(tableDef *transformBuildConfig) transform.TransformFunc {
	return func(_ context.Context, d *transform.TransformData) (interface{}, error) {
		if tableDef.transformError == FailError {
			return nil, errors.New("TRANSFORM ERROR")
		}
		if tableDef.transformError == FailPanic {
			panic("TRANSFORM PANIC")
		}
		if tableDef.transformDelay {
			time.Sleep(delayValue)
		}
		item := d.HydrateItem.(map[string]interface{})
		id := item["id"].(int)
		return id, nil
	}
}
