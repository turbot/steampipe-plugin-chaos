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

func commonQuals(key ...*plugin.KeyColumn) []*plugin.KeyColumn {
	return append(key, []*plugin.KeyColumn{
		{
			Name:      "total_row_count",
			Require:   plugin.Optional,
			Operators: []string{"="},
		},
	}...)
}

func commonColumns(key ...*plugin.Column) []*plugin.Column {
	return append(key, []*plugin.Column{
		{
			Name:        "total_row_count",
			Type:        proto.ColumnType_INT,
			Description: "Column that returns the total number of rows. This is an optional key column. Set this to a desired number to get a specific number of rows. Default is 50.",
			Transform:   transform.FromQual("total_row_count"),
		}}...)
}

// func getCommonData(ctx context.Context, d *plugin.QueryData) commonData {

// }
