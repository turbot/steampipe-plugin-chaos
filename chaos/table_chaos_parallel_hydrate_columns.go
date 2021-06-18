package chaos

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type s1 struct {
	Val string
}
type s2 struct {
	Val string
}
type s3 struct {
	Val string
}
type s4 struct {
	Val string
}
type s5 struct {
	Val string
}
type s6 struct {
	Val string
}
type s7 struct {
	Val string
}
type s8 struct {
	Val string
}
type s9 struct {
	Val string
}
type s10 struct {
	Val string
}
type s11 struct {
	Val string
}
type s12 struct {
	Val string
}
type s13 struct {
	Val string
}
type s14 struct {
	Val string
}
type s15 struct {
	Val string
}
type s16 struct {
	Val string
}
type s17 struct {
	Val string
}
type s18 struct {
	Val string
}
type s19 struct {
	Val string
}
type s20 struct {
	Val string
}

func parallelHydrateColumnsTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_parallel_hydrate_columns",
		Description: "Chaos table to test the execution of multiple hydrate functions and transform functions asynchronously\n The main intention o fthis table is to verify the correct transform data is passed to each transform function",
		List: &plugin.ListConfig{
			Hydrate: hydList,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:      "column_1",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate1,
				Transform: transform.From(parallelTransform1),
			},
			{
				Name:      "column_2",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate2,
				Transform: transform.From(parallelTransform2),
			},
			{
				Name:      "column_3",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate3,
				Transform: transform.From(parallelTransform3),
			},
			{
				Name:      "column_4",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate4,
				Transform: transform.From(parallelTransform4),
			},
			{
				Name:      "column_5",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate5,
				Transform: transform.From(parallelTransform5),
			},
			{
				Name:      "column_6",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate6,
				Transform: transform.From(parallelTransform6),
			},
			{
				Name:      "column_7",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate7,
				Transform: transform.From(parallelTransform7),
			},
			{
				Name:      "column_8",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate8,
				Transform: transform.From(parallelTransform8),
			},
			{
				Name:      "column_9",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate9,
				Transform: transform.From(parallelTransform9),
			},
			{
				Name:      "column_10",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate10,
				Transform: transform.From(parallelTransform10),
			},
			{
				Name:      "column_11",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate11,
				Transform: transform.From(parallelTransform11),
			},
			{
				Name:      "column_12",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate12,
				Transform: transform.From(parallelTransform12),
			},
			{
				Name:      "column_13",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate13,
				Transform: transform.From(parallelTransform13),
			},
			{
				Name:      "column_14",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate14,
				Transform: transform.From(parallelTransform14),
			},
			{
				Name:      "column_15",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate15,
				Transform: transform.From(parallelTransform15),
			},
			{
				Name:      "column_16",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate16,
				Transform: transform.From(parallelTransform16),
			},
			{
				Name:      "column_17",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate17,
				Transform: transform.From(parallelTransform17),
			},
			{
				Name:      "column_18",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate18,
				Transform: transform.From(parallelTransform18),
			},
			{
				Name:      "column_19",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate19,
				Transform: transform.From(parallelTransform19),
			},
			{
				Name:      "column_20",
				Type:      proto.ColumnType_STRING,
				Hydrate:   parallelHydrate20,
				Transform: transform.From(parallelTransform20),
			},
		},
	}
}

func hydInputKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	item := quals["id"].GetInt64Value()
	return item, nil
}

func hydList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	for i := 0; i < 500; i++ {
		item := map[string]interface{}{"id": i}
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func hydGet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)

	item := map[string]interface{}{"id": id}
	return item, nil

}

func parallelHydrate1(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(100, 300))

	return &s1{Val: "parallelHydrate1"}, nil
}

func parallelHydrate2(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(250, 400))

	return &s2{Val: "parallelHydrate2"}, nil
}

func parallelHydrate3(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(350, 500))

	return &s3{Val: "parallelHydrate3"}, nil
}

func parallelHydrate4(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(450, 600))

	return &s4{Val: "parallelHydrate4"}, nil
}

func parallelHydrate5(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(550, 700))

	return &s5{Val: "parallelHydrate5"}, nil
}

func parallelHydrate6(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(650, 800))

	return &s6{Val: "parallelHydrate6"}, nil
}
func parallelHydrate7(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(750, 900))

	return &s7{Val: "parallelHydrate7"}, nil
}
func parallelHydrate8(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(850, 1000))

	return &s8{Val: "parallelHydrate8"}, nil
}
func parallelHydrate9(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(500, 650))

	return &s9{Val: "parallelHydrate9"}, nil
}
func parallelHydrate10(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(100, 250))

	return &s10{Val: "parallelHydrate10"}, nil
}
func parallelHydrate11(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(800, 950))

	return &s11{Val: "parallelHydrate11"}, nil
}
func parallelHydrate12(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(250, 750))

	return &s12{Val: "parallelHydrate12"}, nil
}
func parallelHydrate13(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(300, 650))

	return &s13{Val: "parallelHydrate13"}, nil
}
func parallelHydrate14(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(50, 300))

	return &s14{Val: "parallelHydrate14"}, nil
}
func parallelHydrate15(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(100, 450))

	return &s15{Val: "parallelHydrate15"}, nil
}
func parallelHydrate16(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(200, 650))

	return &s16{Val: "parallelHydrate16"}, nil
}
func parallelHydrate17(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(250, 650))

	return &s17{Val: "parallelHydrate17"}, nil
}
func parallelHydrate18(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(350, 850))

	return &s18{Val: "parallelHydrate18"}, nil
}
func parallelHydrate19(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(450, 950))

	return &s19{Val: "parallelHydrate19"}, nil
}
func parallelHydrate20(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(randomTimeDelay(550, 850))

	return &s20{Val: "parallelHydrate20"}, nil
}

func parallelTransform1(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s1)

	return item.Val, nil
}

func parallelTransform2(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s2)

	return item.Val, nil
}

func parallelTransform3(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s3)

	return item.Val, nil
}

func parallelTransform4(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s4)

	return item.Val, nil
}

func parallelTransform5(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s5)

	return item.Val, nil
}

func parallelTransform6(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s6)

	return item.Val, nil
}

func parallelTransform7(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s7)

	return item.Val, nil
}

func parallelTransform8(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s8)

	return item.Val, nil
}

func parallelTransform9(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s9)

	return item.Val, nil
}

func parallelTransform10(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s10)

	return item.Val, nil
}

func parallelTransform11(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s11)

	return item.Val, nil
}

func parallelTransform12(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s12)

	return item.Val, nil
}

func parallelTransform13(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s13)

	return item.Val, nil
}

func parallelTransform14(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s14)

	return item.Val, nil
}

func parallelTransform15(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s15)

	return item.Val, nil
}

func parallelTransform16(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s16)

	return item.Val, nil
}

func parallelTransform17(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s17)

	return item.Val, nil
}

func parallelTransform18(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s18)

	return item.Val, nil
}

func parallelTransform19(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s19)

	return item.Val, nil
}

func parallelTransform20(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(*s20)

	return item.Val, nil
}
