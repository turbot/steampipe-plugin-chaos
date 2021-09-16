package chaos

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const allColumnsRowCount = 100

type JSONPolicy struct {
	Name      string
	Id        int
	Statement Statement
}

type Statement struct {
	Action string
	Effect string
}

type TransformColumnAttributes struct {
	Key   string
	Value string
}

func allColumnsTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_all_column_types",
		Description: "Chaos table to test all columns of different types",
		List: &plugin.ListConfig{
			Hydrate: allColumnsList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{
				Name:        "empty_hydrate",
				Type:        proto.ColumnType_STRING,
				Default:     "I AM A DEFAULT",
				Hydrate:     emptyHydrate,
				Description: "This column tests both a hydrate function returning no results, and the column default mechanism",
			},

			{Name: "string_column", Type: proto.ColumnType_STRING},
			{
				Name:      "boolean_column",
				Type:      proto.ColumnType_BOOL,
				Hydrate:   booleanColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:      "date_time_column",
				Type:      proto.ColumnType_DATETIME,
				Hydrate:   dateTimeColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:      "double_column",
				Type:      proto.ColumnType_DOUBLE,
				Hydrate:   doubleColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:      "ipaddress_column",
				Type:      proto.ColumnType_IPADDR,
				Hydrate:   ipAddressColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:      "json_column",
				Type:      proto.ColumnType_JSON,
				Hydrate:   jsonColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:      "cidr_column",
				Type:      proto.ColumnType_CIDR,
				Hydrate:   cidrColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:      "long_string_column",
				Type:      proto.ColumnType_STRING,
				Hydrate:   longStringColumnValue,
				Transform: transform.FromValue(),
			},
			{
				Name:        "array_element",
				Description: "This column test the functionality of array lookup in Transform function",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("TransformColumn[0]"),
			},
			{
				Name:      "epoch_column_seconds",
				Type:      proto.ColumnType_DATETIME,
				Hydrate:   epochColumnSecValue,
				Transform: transform.FromValue().Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "epoch_column_milliseconds",
				Type:      proto.ColumnType_DATETIME,
				Hydrate:   epochColumnMsValue,
				Transform: transform.FromValue().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:      "string_to_array_column",
				Type:      proto.ColumnType_JSON,
				Hydrate:   stringToArrayColumn,
				Transform: transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:      "array_to_maps_column",
				Type:      proto.ColumnType_JSON,
				Hydrate:   stringArrayToMapColumn,
				Transform: transform.FromValue().Transform(transform.StringArrayToMap),
			},
		},
	}

}

func allColumnsList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < allColumnsRowCount; i++ {
		id := i
		columnVal := "\u0000"
		if i%2 == 0 {
			columnVal = d.SteampipeMetadata.SteampipeVersion
		}
		columnAttributes := TransformColumnAttributes{Key: columnVal, Value: "value"}
		var transformColumn []TransformColumnAttributes
		transformColumn = append(transformColumn, columnAttributes)
		item := map[string]interface{}{"id": id, "string_column": columnVal, "TransformColumn": transformColumn}
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func booleanColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := id%2 == 0
	return columnVal, nil

}

func dateTimeColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	dates := []string{"2001-02-20T01:28:00Z", "2002-05-05T09:23:00Z", "2001-08-27T05:00:00Z", "2001-12-12T06:19:00Z", "2001-07-14T18:00:34Z", "2001-09-11T07:49:00Z", "2001-06-04T22:00:32Z"}
	date, _ := time.Parse(time.RFC3339, dates[id%len(dates)])
	return date, nil
}

func doubleColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := float64(id) / 17
	return columnVal, nil
}

func ipAddressColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	ipAddress := []string{"10.0.1.4", "10.0.0.1", "10.0.2.2", "10.0.1.1", "10.0.0.4", "10.0.0.5", "10.0.0.10", "10.0.1.3", "10.0.0.7", "10.0.0.6", "10.0.0.8", "10.0.2.9", "10.0.1.10"}
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := ipAddress[id%len(ipAddress)]
	return columnVal, nil
}

func jsonColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	name := "\u0000"
	actions := []string{"iam:GetContextKeysForCustomPolicy", "iam:GetContextKeysForPrincipalPolicy", "iam:SimulateCustomPolicy", "iam:SimulatePrincipalPolicy"}
	effects := []string{"Allow", "Deny"}
	statement := Statement{Action: actions[id%len(actions)], Effect: effects[id%len(effects)]}
	JSONPolicy := JSONPolicy{Name: name, Id: id, Statement: statement}
	return JSONPolicy, nil

}

func cidrColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cidr := []string{"10.0.0.0/24", "10.0.0.1/32", "10.84.0.0/24", "172.31.0.0/16", "192.168.0.0/22", "172.16.0.0/16", "10.1.0.0/16", "175.0.0.0/16"}
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	columnVal := cidr[id%len(cidr)]
	return columnVal, nil
}

func emptyHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return nil, nil

}

func epochColumnSecValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	dates := []string{"1611063579", "1611064454", "1585990763", "1612145454", "1562145454", "1555145454", "1612145454"}
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	item := dates[id%len(dates)]

	return item, nil
}

func epochColumnMsValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	dates := []string{"1611057198070", "1598057778070", "1432117198070", "1699851193076", "1611599593076", "1611335198070", "166657198070"}
	key := h.Item.(map[string]interface{})
	id := key["id"].(int)
	item := dates[id%len(dates)]

	return item, nil
}

func stringToArrayColumn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	item := key["string_column"]

	return item, nil

}

func stringArrayToMapColumn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := h.Item.(map[string]interface{})
	var item []string
	item = append(item, key["string_column"].(string))

	return item, nil

}

func longStringColumnValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	stringLength := 100
	key := h.Item.(map[string]interface{})
	item := fmt.Sprintf(strings.Repeat(key["string_column"].(string), stringLength))

	return item, nil

}
