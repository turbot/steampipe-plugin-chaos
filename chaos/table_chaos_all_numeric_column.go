package chaos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func numericColumnsTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_all_numeric_column",
		Description: "Chaos table to test all flavours of integer and float data types",

		List: &plugin.ListConfig{
			Hydrate: numericList,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "int_data", Type: proto.ColumnType_INT, Hydrate: intData, Transform: transform.FromValue()},
			{Name: "int8_data", Type: proto.ColumnType_INT, Hydrate: int8Data, Transform: transform.FromValue()},
			{Name: "int16_data", Type: proto.ColumnType_INT, Hydrate: int16Data, Transform: transform.FromValue()},
			{Name: "int32_data", Type: proto.ColumnType_INT, Hydrate: int32Data, Transform: transform.FromValue()},
			{Name: "int64_data", Type: proto.ColumnType_INT, Hydrate: int64Data, Transform: transform.FromValue()},
			{Name: "uint_data", Type: proto.ColumnType_INT, Hydrate: uintData, Transform: transform.FromValue()},
			{Name: "uint8_data", Type: proto.ColumnType_INT, Hydrate: uint8Data, Transform: transform.FromValue()},
			{Name: "uint16_data", Type: proto.ColumnType_INT, Hydrate: uint16Data, Transform: transform.FromValue()},
			{Name: "uint32_data", Type: proto.ColumnType_INT, Hydrate: uint32Data, Transform: transform.FromValue()},
			{Name: "uint64_data", Type: proto.ColumnType_INT, Hydrate: uint64Data, Transform: transform.FromValue()},

			{Name: "float32_data", Type: proto.ColumnType_DOUBLE, Hydrate: float32Data, Transform: transform.FromValue()},
			{Name: "float64_data", Type: proto.ColumnType_DOUBLE, Hydrate: float64Data, Transform: transform.FromValue()},
		},
	}
}

func numericList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for i := 30; i < 40; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)

	}
	return nil, nil
}

func intData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	return id, nil
}

func int8Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := int8(id) * 3
	return columnVal, nil
}

func int16Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := int16(id) * 8
	return columnVal, nil
}

func int32Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := int32(id) * 10
	return columnVal, nil
}

func int64Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := int64(id) * 15
	return columnVal, nil
}

func uintData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := uint(id) * 12
	return columnVal, nil
}

func uint8Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := uint8(id) * 16
	return columnVal, nil
}

func uint16Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := uint16(id) * 11
	return columnVal, nil
}

func uint32Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := uint32(id) * 23
	return columnVal, nil
}

func uint64Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := uint64(id) * 17
	return columnVal, nil
}

func float32Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := float32(id) / 7
	return columnVal, nil
}

func float64Data(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := float64(id) / 12
	return columnVal, nil
}
