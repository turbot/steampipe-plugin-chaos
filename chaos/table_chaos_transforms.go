package chaos

import (
	"context"
	"fmt"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

type ListStruct struct {
	Id                int
	OptionalKeyColumn *string
	FromTagColumn     *string `fromtag:"from_tag_column,omitempty"`
	FromJSONTag       *string `json:"from_json_tag,omitempty"`
	FromFieldColumn   *string
	ColumnD           *string
}

type GetStruct struct {
	Id                int
	OptionalKeyColumn *string
	FromTagColumn     *string `fromtag:"from_tag_column,omitempty"`
	FromJSONTag       *string `json:"from_json_tag,omitempty"`
	ColumnC           *string
	ColumnD           *string
}

func (t ListStruct) TransformMethod() (string, error) {
	return "Transform method", nil
}
func (t GetStruct) TransformMethod() (string, error) {
	return "Transform method", nil
}

func transformFromFunctionsTable() *plugin.Table {
	return &plugin.Table{
		Name:             "chaos_transforms",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate:    transformFromFieldList,
			KeyColumns: plugin.OptionalColumns([]string{"optional_key_column"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    transformFromFieldGet,
		},
		GetMatrixItem: BuildMatrixItem,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{
				Name:        "optional_key_column",
				Description: "column to test optional key column quals",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "from_json_tag",
				Description: "column to test FromJSONTag transform",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "from_field_column_single",
				Type:        proto.ColumnType_STRING,
				Description: "Column to test FromField transform function from a single source",
				Transform:   transform.FromField("ColumnD"),
			},
			{
				Name:        "from_field_column_multiple",
				Type:        proto.ColumnType_STRING,
				Description: "Column to test FromField transform function from multiple sources",
				Transform:   transform.FromField("FromFieldColumn", "ColumnC"),
			},
			{
				Name:        "from_qual_column",
				Type:        proto.ColumnType_STRING,
				Description: "Column to test FromQual transform function with KeyColumn",
				Transform:   transform.FromQual("id"),
			},
			{
				Name:        "from_optional_qual_column",
				Type:        proto.ColumnType_STRING,
				Description: "Column to test FromQual transform function with OptionalKeyColumn",
				Transform:   transform.FromQual("optional_key_column"),
			},
			{
				Name:        "transform_method_column",
				Type:        proto.ColumnType_STRING,
				Description: "column to test the FromMethod transform by invoking a function on the hydrate item",
				Transform:   transform.FromMethod("TransformMethod"),
			},
			{
				Name:        "from_value_column",
				Type:        proto.ColumnType_STRING,
				Description: "column to test the FromValue transform function",
				Hydrate:     transformHydrate,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "from_tag_column",
				Type:        proto.ColumnType_STRING,
				Description: "column to test the FromTag transform function",
				Transform:   transform.FromTag("fromtag"),
			},
			{
				Name:        "from_constant_column",
				Type:        proto.ColumnType_STRING,
				Description: "column to test the FromConstant transform function",
				Transform:   transform.FromConstant("from constant"),
			},
			{
				Name:        "from_transform_column",
				Type:        proto.ColumnType_STRING,
				Description: "column to test the From transform",
				Transform:   transform.From(transformFunction),
			},
			{
				Name:        "from_matrix_item_column",
				Type:        proto.ColumnType_STRING,
				Description: "column to test the FromMatrixItem transform",
				Transform:   transform.FromMatrixItem("location"),
			},
		},
	}
}

func BuildMatrixItem(ctx context.Context, connection *plugin.Connection) []map[string]interface{} {
	matrix := make([]map[string]interface{}, 1)
	matrix[0] = map[string]interface{}{"location": "location1"}
	return matrix
}

func transformFromFieldList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	log.Printf("[INFO] transformFromFieldList")
	for i := 0; i < 2; i++ {
		columnA := fmt.Sprintf("column-%d", i)
		columnOptional := fmt.Sprintf("optional-column-%d", i)
		fromTagColumn := "from tag"
		column := "THIS IS COMING FROM LIST CALL"
		columnD := "LIST CALL COLUMN D"
		item := ListStruct{Id: i, OptionalKeyColumn: &columnOptional, FromTagColumn: &fromTagColumn, FromJSONTag: &columnA, FromFieldColumn: &column, ColumnD: &columnD}
		log.Printf("[INFO] transformFromFieldList STREAM")
		d.StreamLeafListItem(ctx, item)
	}
	return nil, nil
}

func transformFromFieldGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	i := d.KeyColumnQuals["id"].GetInt64Value()
	columnA := fmt.Sprintf("column-%d", i)
	columnOptional := fmt.Sprintf("optional-column-%d", i)
	fromTagColumn := "from tag"
	column := "THIS IS COMING FROM GET CALL"
	columnD := "LIST CALL COLUMN D"
	item := GetStruct{Id: int(i), OptionalKeyColumn: &columnOptional, FromTagColumn: &fromTagColumn, FromJSONTag: &columnA, ColumnC: &column, ColumnD: &columnD}
	return item, nil
}

func transformHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	columnVal := "From Value"

	return columnVal, nil
}

func transformFunction(_ context.Context, d *transform.TransformData) (interface{}, error) {
	value := "from transform function"

	return value, nil
}
