package chaos

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type parentData struct {
	id int
}

type childData struct {
	id           int
	child_column string
}

func chaosListParentChildHydrateTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_parent_child_hydrate_data",
		Description: "Chaos table to test the List calls having parent-child dependencies with all the possible scenarios like errors, panics and delays at both parent and child levels",
		List: &plugin.ListConfig{
			Hydrate:       listChildHydrateTable,
			ParentHydrate: listParentHydrateTable,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("child_id"),
			Hydrate:    getTable,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "child_id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "child_column", Type: proto.ColumnType_STRING, Description: "Column to test the the parent list function with fatal error after streaming some rows"},
			{Name: "transform_column", Type: proto.ColumnType_STRING, Description: "Column to test the the parent list function with retry in case of non fatal error", Transform: transform.From(transformFunction)},
		},
	}
}

func listParentHydrateTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 2; i++ {
		item := parentData{id: i}
		log.Printf("[ERROR] PARENT CALL EXECUTED, HAS =======> %v", item)
		d.StreamListItem(ctx, item)
	}

	return nil, nil

}

func listChildHydrateTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(parentData)
	id := key.id

	item := map[string]interface{}{"id": id, "child_id": id, "child_column": fmt.Sprintf("child_column-%d", id)}
	d.StreamLeafListItem(ctx, item)

	return nil, nil

}

func getTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.KeyColumnQuals["child_id"].GetInt64Value()

	item := map[string]interface{}{"id": id, "child_id": id, "child_column": fmt.Sprintf("child_column-get-%v", id)}
	d.StreamListItem(ctx, item)
	return nil, nil
}

func transformFunction(_ context.Context, d *transform.TransformData) (interface{}, error) {
	log.Printf("[WARN] The whole Transform data====> %v", reflect.TypeOf(d.HydrateItem))
	key := d.HydrateItem.(map[string]interface{})
	log.Printf("[WARN] The data in Transform data====> %v", key)
	childData := key["child_column"]
	id := key["id"]

	item := fmt.Sprintf("%s-%d", childData, id)

	return item, nil

}
