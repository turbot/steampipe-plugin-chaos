package chaos

import (
	"context"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func chaosLimitVerifyRowsRemainingTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_limit_verify_rows_remaining",
		Description: "Chaos table to check the SDK rows remaining functionality when a limit is passed to a query.",
		List: &plugin.ListConfig{
			Hydrate:    listRowsRemaining,
			KeyColumns: commonQuals(),
		},

		Columns: commonColumns(
			&plugin.Column{Name: "rows_streamed", Type: proto.ColumnType_INT, Description: "Column that returns the number of row streamed."},
			&plugin.Column{Name: "sdk_rows_remaining", Type: proto.ColumnType_INT, Description: "Column that returns the number of rows remaining to be streamed."},
			&plugin.Column{Name: "limit_value", Type: proto.ColumnType_INT, Description: "Column that returns the limit value which is passed in the query."},
		),
	}
}

func listRowsRemaining(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	limit := d.QueryContext.Limit
	rows := DefaultTotalRowCount
	if int(d.EqualsQuals["row_count"].GetInt64Value()) != 0 {
		rows = int(d.EqualsQuals["row_count"].GetInt64Value())
	}
	log.Printf("[INFO] row_count=%d", rows)
	log.Printf("[INFO] limit=%d", limit)

	for i := 0; i < rows; i++ {
		item := map[string]interface{}{
			"rows_streamed":      i,
			"sdk_rows_remaining": d.RowsRemaining(ctx),
			"limit_value":        limit,
		}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
