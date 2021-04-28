package chaos

import (
	"context"
	"fmt"
	log "log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type ListStruct struct {
	Id              int
	ColumnA         *string
	FromFieldColumn *string
}

type GetStruct struct {
	Id      int
	ColumnA *string
	ColumnC *string
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
				Name:      "from_field_column",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromFields([]string{"FromFieldColumn", "ColumnC"}),
			},
		},
	}
}

func transformFromFieldList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[ERROR] INSIDE LIST CALL")
	for i := 0; i < 5; i++ {
		columnA := fmt.Sprintf("column-%d", i)
		column := "THIS IS COMING FROM LIST CALL"
		item := ListStruct{Id: i, ColumnA: &columnA, FromFieldColumn: &column}
		d.StreamLeafListItem(ctx, item)
	}
	return nil, nil
}

func transformFromFieldGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log.Println("[ERROR] INSIDE GET CALL")
	i := d.KeyColumnQuals["id"].GetInt64Value()
	columnA := fmt.Sprintf("column-%d", i)
	column := "THIS IS COMING FROM GET CALL"
	item := GetStruct{Id: int(i), ColumnA: &columnA, ColumnC: &column}
	return item, nil
}
