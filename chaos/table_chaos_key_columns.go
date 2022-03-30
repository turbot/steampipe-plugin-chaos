package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type keyColumnType string

const (
	singleColumn keyColumnType = "single"
	anyColumn    keyColumnType = "any"
	allColumns   keyColumnType = "all"
)

type keyColumnTableDefinitions struct {
	listKeyColumnSetType keyColumnType
	getKeyColumnSetType  keyColumnType
	name                 string
	description          string
}

func KeyColumnTableBuilder(def *keyColumnTableDefinitions) *plugin.Table {
	table := &plugin.Table{
		Name:        def.name,
		Description: def.description,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_a", Type: proto.ColumnType_STRING},
			{Name: "combined_columns", Type: proto.ColumnType_STRING},
		},
	}
	if def.listKeyColumnSetType != "" {
		table.List = &plugin.ListConfig{
			Hydrate:    buildListUsingKeyColumns(),
			KeyColumns: calculateListKeyColumns(def.listKeyColumnSetType),
		}
	}
	if def.getKeyColumnSetType != "" {
		table.Get = &plugin.GetConfig{
			Hydrate:    buildGetUsingKeyColumns(),
			KeyColumns: calculateGetKeyColumns(def.getKeyColumnSetType),
		}
	}

	return table
}

func buildListUsingKeyColumns() plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		item := getItemFromKeyColumns(d)
		d.StreamListItem(ctx, item)

		return nil, nil
	}
}

func buildGetUsingKeyColumns() plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		item := getItemFromKeyColumns(d)
		return item, nil
	}
}
func getItemFromKeyColumns(d *plugin.QueryData) map[string]interface{} {
	item := map[string]interface{}{}
	id, gotId := d.KeyColumnQuals["id"]
	if gotId {
		item = map[string]interface{}{"id": id.GetInt64Value()}
	}
	columnA, gotColumnA := d.KeyColumnQuals["column_a"]
	if gotColumnA {
		item["column_a"] = columnA.GetStringValue()
	}
	if gotId && gotColumnA {
		item["combine_column"] = fmt.Sprintf("%s-%s", id, columnA)
	}
	return item
}

func calculateListKeyColumns(keyColumnSetType keyColumnType) plugin.KeyColumnSlice {
	if keyColumnSetType == singleColumn {
		return plugin.SingleColumn("id")
	}
	if keyColumnSetType == anyColumn {
		return plugin.AnyColumn([]string{"id", "column_a"})
	}
	if keyColumnSetType == allColumns {
		return plugin.AllColumns([]string{"id", "column_a"})
	}
	return nil
}

func calculateGetKeyColumns(keyColumnSetType keyColumnType) plugin.KeyColumnSlice {
	if keyColumnSetType == anyColumn {
		return plugin.AnyColumn([]string{"id", "column_a"})
	}
	if keyColumnSetType == allColumns {
		return plugin.AllColumns([]string{"id", "column_a"})
	}
	return plugin.SingleColumn("id")
}
