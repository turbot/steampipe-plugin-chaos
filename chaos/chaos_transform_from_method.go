package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type structWithMethod struct {
	Id      int
	Column1 string
	Column2 string
	Column3 string
	Column4 string
}

func (t structWithMethod) TransformMethod() (string, error) {
	return "column_4", nil
}

func transformFromMethodTable() *plugin.Table {
	return &plugin.Table{
		Name:             "chaos_transform_from_method",
		Description:      "Chaos table to test the transform FromMethod invoking a function on the hydrate item",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: transformMethodList,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_1", Type: proto.ColumnType_STRING},
			{Name: "column_2", Type: proto.ColumnType_STRING},
			{Name: "column_3", Type: proto.ColumnType_STRING},
			{
				Name:      "column_4",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromMethod("TransformMethod"),
			},
		},
	}
}

func transformMethodList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		id := i
		column1 := fmt.Sprintf("column_1-%d", id)
		column2 := fmt.Sprintf("column_2-%d", id)
		column3 := fmt.Sprintf("column_3-%d", id)

		item := structWithMethod{Id: id, Column1: column1, Column2: column2, Column3: column3}
		d.StreamLeafListItem(ctx, item)
	}
	return nil, nil
}
