package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type ListStruct struct {
	Id              int
	ColumnA         *string
	FromFieldColumn *string
	ColumnD         *string
}

type GetStruct struct {
	Id      int
	ColumnA *string
	ColumnC *string
	ColumnD *string
}

func transformFromFieldsTable() *plugin.Table {
	return &plugin.Table{
		Name:             "chaos_transform_from_fields",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: transformFromFieldList,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    transformFromFieldGet,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_a", Type: proto.ColumnType_STRING},
			{
				Name:      "from_field_column_single",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("ColumnD"),
			},
			{
				Name:      "from_field_column_multiple",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("FromFieldColumn", "ColumnC"),
			},
			// {
			// 	Name:      "from_qual_column",
			// 	Type:      proto.ColumnType_STRING,
			// 	Transform: transform.FromQual("id"),
			// },
		},
	}
}

func transformFromFieldList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 5; i++ {
		columnA := fmt.Sprintf("column-%d", i)
		column := "THIS IS COMING FROM LIST CALL"
		columnD := "LIST CALL COLUMN D"
		item := ListStruct{Id: i, ColumnA: &columnA, FromFieldColumn: &column, ColumnD: &columnD}
		d.StreamLeafListItem(ctx, item)
	}
	return nil, nil
}

func transformFromFieldGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	i := d.KeyColumnQuals["id"].GetInt64Value()
	columnA := fmt.Sprintf("column-%d", i)
	column := "THIS IS COMING FROM GET CALL"
	columnD := "LIST CALL COLUMN D"
	item := GetStruct{Id: int(i), ColumnA: &columnA, ColumnC: &column, ColumnD: &columnD}
	return item, nil
}
