package chaos

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//	type commonData struct {
//		totalRowCount *int
//	}

const DefaultTotalRowCount = 50

// list of common key column quals that can be used with all chaos tables
func commonQuals(key ...*plugin.KeyColumn) []*plugin.KeyColumn {
	return append(key, []*plugin.KeyColumn{
		{
			Name:      "row_count",
			Require:   plugin.Optional,
			Operators: []string{"="},
		},
	}...)
}

// list of common columns that can be used with all chaos tables
func commonColumns(key ...*plugin.Column) []*plugin.Column {
	return append(key, []*plugin.Column{
		{
			Name:        "row_count",
			Type:        proto.ColumnType_INT,
			Description: "Total number of rows returned. This is an optional key column. Pass this as a qual to set the number of rows returned. Default is 50.",
			Transform:   transform.FromQual("row_count"),
		}}...)
}

// func getCommonData(ctx context.Context, d *plugin.QueryData) commonData {

// }
