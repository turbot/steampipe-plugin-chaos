package chaos

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

type callType string
type keyColumnType string

const (
	listCall     callType      = "list"
	getCall      callType      = "get"
	singleColumn keyColumnType = "single"
	anyColumn    keyColumnType = "any"
	allColumns   keyColumnType = "all"
)

type keyColumnTableDefinitions struct {
	name             string
	description      string
	call             callType
	keyColumnSetType keyColumnType
}

func KeyColumnTableBuilder(def *keyColumnTableDefinitions) *plugin.Table {
	return &plugin.Table{
		Name:        def.name,
		Description: def.description,
		List: &plugin.ListConfig{
			Hydrate:    buildKeyColumns(def),
			KeyColumns: calculateKeyColumns(def),
		},
		Get: &plugin.GetConfig{
			KeyColumns: calculateKeyColumns(def),
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
		if def.keyColumnSetType == singleColumn {
			id := d.KeyColumnQuals["id"].GetInt64Value()
			item = map[string]interface{}{"id": id}
		}
		if def.keyColumnSetType == anyColumn || def.keyColumnSetType == allColumns {
			id := d.KeyColumnQuals["id"].GetInt64Value()
			columnA := d.KeyColumnQuals["column_a"].GetStringValue()
			combinedColumn := fmt.Sprintf("%d-%s", id, columnA)
			item = map[string]interface{}{"id": id, "column_a": columnA, "combined_columns": combinedColumn}
		}
		d.StreamListItem(ctx, item)
		return nil, nil
	}
}

// func buildGetKeyColumns(def *keyColumnTableDefinitions) plugin.HydrateFunc {
// 	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 		i := d.KeyColumnQuals["id"].GetInt64Value()
// 		columnA := d.KeyColumnQuals["column_a"].GetStringValue()
// 		columnB := "RANDOM STRING"
// 		columnC := "RANDOM STRING"
// 		item := map[string]interface{}{"id": i, "column_a": columnA, "column_b": columnB, "column_c": columnC}
// 		return item, nil

// 	}
// }

func calculateKeyColumns(def *keyColumnTableDefinitions) *plugin.KeyColumnSet {
	if def.call == listCall || def.call == getCall {
		if def.keyColumnSetType == singleColumn {
			return plugin.SingleColumn("id")
		}
		if def.keyColumnSetType == anyColumn {
			return plugin.AnyColumn([]string{"id", "column_a"})
		}

		if def.keyColumnSetType == allColumns {
			return plugin.AllColumns([]string{"id", "column_a"})
		}
	}
	if def.call != getCall {
		return plugin.SingleColumn("id")
	}
	return nil
}

// func calculateGetKeyColumns(def *keyColumnTableDefinitions) *plugin.KeyColumnSet {
// 	if def.call == getCall {
// 		if def.keyColumnSetType == singleColumn {
// 			return plugin.SingleColumn("id")
// 		}
// 		if def.keyColumnSetType == anyColumn {
// 			return plugin.AnyColumn([]string{"id", "column_a"})
// 		}

// 		if def.keyColumnSetType == allColumns {
// 			return plugin.AllColumns([]string{"id", "column_a"})
// 		}
// 	}
// 	return plugin.SingleColumn("id")
// }
