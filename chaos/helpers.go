package chaos

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
)

func populateItem(rowNumber int, table *plugin.Table) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = rowNumber
	for _, column := range table.Columns {
		var columnVal interface{}
		switch column.Type {
		case proto.ColumnType_STRING:
			columnVal = fmt.Sprintf("%s-%v", column.Name, rowNumber)
			break
		case proto.ColumnType_BOOL:
			columnVal = rowNumber%2 == 0
			break
		case proto.ColumnType_DATETIME:
			columnVal = time.Now()
			break
		case proto.ColumnType_INT:
			columnVal = rowNumber
			break
		case proto.ColumnType_DOUBLE:
			columnVal = float64(rowNumber)
			break
		case proto.ColumnType_CIDR:
			columnVal = "10.0.0.10/32"
			break
		case proto.ColumnType_IPADDR:
			columnVal = "10.0.0.2"
			break
		case proto.ColumnType_JSON:
			columnVal = `{"Version": "2012-10-17","Statement": [{"Action": ["iam:GetContextKeysForCustomPolicy","iam:GetContextKeysForPrincipalPolicy","iam:SimulateCustomPolicy","iam:SimulatePrincipalPolicy"],"Effect": "Allow","Resource": "*"}]}`
			break
		}
		row[column.Name] = columnVal
	}
	return row

}

func randomTimeDelay(minMs int, maxMs int) time.Duration {
	delta := rand.Intn(maxMs-minMs) + minMs
	timeDelay := time.Duration(delta) * time.Millisecond
	return timeDelay
}
