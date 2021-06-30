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
			Hydrate:    buildFetchHydrateUsingKeyColumns(def.listKeyColumnSetType),
			KeyColumns: calculateListKeyColumns(def.listKeyColumnSetType),
		}
	}
	if def.getKeyColumnSetType != "" {
		table.Get = &plugin.GetConfig{
			Hydrate:    buildFetchHydrateUsingKeyColumns(def.getKeyColumnSetType),
			KeyColumns: calculateGetKeyColumns(def.getKeyColumnSetType),
		}
	}

	return table
}

func buildFetchHydrateUsingKeyColumns(keyColumnSetType keyColumnType) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		item := map[string]interface{}{}
		if keyColumnSetType == singleColumn {
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
