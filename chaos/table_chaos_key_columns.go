package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
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
	return &plugin.Table{
		Name:        def.name,
		Description: def.description,
		List: &plugin.ListConfig{
			Hydrate:    buildKeyColumns(def),
			KeyColumns: calculateListKeyColumns(def),
		},
		Get: &plugin.GetConfig{
			KeyColumns: calculateGetKeyColumns(def),
			Hydrate:    buildKeyColumns(def),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_a", Type: proto.ColumnType_STRING},
			{Name: "combined_columns", Type: proto.ColumnType_STRING},
		},
	}

}

func buildKeyColumns(def *keyColumnTableDefinitions) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		item := map[string]interface{}{}
		if def.listKeyColumnSetType == singleColumn || def.getKeyColumnSetType == singleColumn {
			id := d.KeyColumnQuals["id"].GetInt64Value()
			item = map[string]interface{}{"id": id}
		} else {
			id := d.KeyColumnQuals["id"].GetInt64Value()
			columnA := d.KeyColumnQuals["column_a"].GetStringValue()
			combinedColumn := fmt.Sprintf("%d-%s", id, columnA)
			item = map[string]interface{}{"id": id, "column_a": columnA, "combined_columns": combinedColumn}
		}
		d.StreamListItem(ctx, item)
		return nil, nil
	}
}

func calculateListKeyColumns(def *keyColumnTableDefinitions) *plugin.KeyColumnSet {
	if def.listKeyColumnSetType == singleColumn {
		return plugin.SingleColumn("id")
	}
	if def.listKeyColumnSetType == anyColumn {
		return plugin.AnyColumn([]string{"id", "column_a"})
	}
	if def.listKeyColumnSetType == allColumns {
		return plugin.AllColumns([]string{"id", "column_a"})
	}
	return nil
}

func calculateGetKeyColumns(def *keyColumnTableDefinitions) *plugin.KeyColumnSet {
	if def.getKeyColumnSetType == anyColumn {
		return plugin.AnyColumn([]string{"id", "column_a"})
	}
	if def.getKeyColumnSetType == allColumns {
		return plugin.AllColumns([]string{"id", "column_a"})
	}
	return plugin.SingleColumn("id")
}
