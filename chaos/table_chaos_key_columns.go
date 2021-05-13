package chaos

import (
	"context"

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
			Hydrate:    buildListKeyColumns(def),
			KeyColumns: calculateListKeyColumns(def),
		},
		// Get: &plugin.GetConfig{
		// 	KeyColumns: calculateGetKeyColumns(def),
		// 	Hydrate:    buildGetKeyColumns(def),
		// },
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "column_a", Type: proto.ColumnType_STRING},
			{Name: "column_b", Type: proto.ColumnType_STRING},
			{Name: "column_c", Type: proto.ColumnType_STRING},
		},
	}

}

func buildListKeyColumns(def *keyColumnTableDefinitions) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		item := map[string]interface{}{}
		if def.keyColumnSetType == singleColumn || def.keyColumnSetType == anyColumn {
			id := d.KeyColumnQuals["id"].GetInt64Value()
			item = map[string]interface{}{"id": id, "column_a": "columnA", "column_b": "columnB", "column_c": "columnC"}
		}
		if def.keyColumnSetType == allColumns {
			id := d.KeyColumnQuals["id"].GetInt64Value()
			columnA := d.KeyColumnQuals["column_a"].GetStringValue()
			item = map[string]interface{}{"id": id, "column_a": columnA, "column_b": "columnB", "column_c": "columnC"}
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

func calculateListKeyColumns(def *keyColumnTableDefinitions) *plugin.KeyColumnSet {
	if def.call == listCall {
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
