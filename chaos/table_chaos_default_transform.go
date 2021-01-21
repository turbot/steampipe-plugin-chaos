package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type testStruct struct {
	Id      int    `json:"id,omitempty"`
	Column1 string `json:"column_1,omitempty"`
	Column2 string `json:"column_2,omitempty"`
	Column3 string `json:"column_3,omitempty"`
}

func defaultTransformTable() *plugin.Table {
	return &plugin.Table{
		Name:             "chaos_default_transform",
		Description:      "Chaos table to test the default transform functionality from specified json tags",
		DefaultTransform: transform.FromJSONTag(),
		List: &plugin.ListConfig{
			Hydrate: transformList,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_1", Type: proto.ColumnType_STRING},
			{Name: "column_2", Type: proto.ColumnType_STRING},
			{Name: "column_3", Type: proto.ColumnType_STRING},
		},
	}
}

func transformList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		id := i
		column1 := fmt.Sprintf("column_1-%d", id)
		column2 := fmt.Sprintf("column_2-%d", id)
		column3 := fmt.Sprintf("column_3-%d", id)

		item := testStruct{Id: id, Column1: column1, Column2: column2, Column3: column3}
		d.StreamLeafListItem(ctx, item)
	}
	return nil, nil
}
