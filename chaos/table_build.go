package chaos

import (
	"context"
	"errors"
	"fmt"
	log "log"
	"strconv"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	t "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type FailType string

const (
	FailNone  FailType = "None"
	FailError          = "Error"
	FailPanic          = "Panic"
)

type chaosTable struct {
	name        string
	description string
	columnCount int
	getError    FailType
	getDelay    bool
	listError   FailType
	// the number of rows returned before a list error/hydrate error is raised
	listErrorRows    int
	listDelay        bool
	rowCount         int
	hydrateError     FailType
	hydrateDelay     bool
	itemFromKeyError FailType
	transformError   FailType
	transformDelay   bool
}

const columnPrefix = "column_"
const defaultColumnCount = 10
const defaultRowCount = 10
const delayValue = 5 * time.Second

func buildTable(tableDef *chaosTable) *plugin.Table {
	return &plugin.Table{
		Name:        tableDef.name,
		Description: tableDef.description,
		List: &plugin.ListConfig{
			Hydrate: getList(tableDef),
		},
		Get: &plugin.GetConfig{
			KeyColumns:  plugin.SingleColumn(columnPrefix + "0"),
			ItemFromKey: buildInputKey,
			Hydrate:     getGet(tableDef),
		},
		Columns: getColumns(tableDef),
	}

}

func getColumns(tableDef *chaosTable) []*plugin.Column {
	var columns []*plugin.Column = []*plugin.Column{{
		Name: "id",
		Type: proto.ColumnType_INT,
	}}
	columnCount := tableDef.columnCount
	if columnCount == 0 {
		columnCount = defaultColumnCount
	}
	for i := 0; i < columnCount; i++ {
		columnName := fmt.Sprintf("%s_%d", "column", i)
		item := &plugin.Column{
			Name: columnName,
			Type: proto.ColumnType_STRING,
		}
		columns = append(columns, item)
	}
	hydrateColumn := &plugin.Column{
		Name:    "hydrate_column",
		Type:    proto.ColumnType_STRING,
		Hydrate: getHydrate(tableDef),
	}
	transformColumn := &plugin.Column{
		Name:      "transform_column",
		Type:      proto.ColumnType_STRING,
		Transform: t.From(getTransform(tableDef)),
	}
	columns = append(columns, transformColumn)
	if tableDef.hydrateDelay || tableDef.hydrateError == FailError || tableDef.hydrateError == FailPanic {
		columns = append(columns, hydrateColumn)
	}
	return columns
}

//// item from key ////

func buildInputKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	keyInput := quals["column_0"].GetStringValue()
	item := keyInput
	return item, nil
}

//// list function ////

func getList(tableDef *chaosTable) plugin.HydrateFunc {

	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		log.Printf("[DEBUG] INSIDE LIST CALL")
		if tableDef.listDelay {
			time.Sleep(delayValue)
		}

		rowCount := tableDef.rowCount
		if rowCount == 0 {
			rowCount = defaultRowCount
		}
		rowsStreamed := 0

		for i := 0; i < rowCount; i++ {

			log.Printf("[DEBUG] ROW LOOP streamed %d error limit %d", rowsStreamed, tableDef.listErrorRows)
			if rowsStreamed >= tableDef.listErrorRows {
				if tableDef.listError == FailError {
					log.Printf("[DEBUG] LIST ERROR ")
					return nil, errors.New("LIST ERROR")
				}
				if tableDef.listError == FailPanic {
					panic("LIST PANIC")
				}
			}

			log.Printf("[DEBUG] STREAM LIST ITEM")

			item := populateItem(i, d.Table)
			d.StreamListItem(ctx, item)
			rowsStreamed++

		}
		return nil, nil
	}
}

/// get the 'Get' hydrate function ///
func getGet(tableDef *chaosTable) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		log.Printf("[DEBUG] INSIDE GET CALL")
		if tableDef.getError == FailError {
			return nil, errors.New("GET ERROR")
		}
		if tableDef.getError == FailPanic {
			panic("GET PANIC")
		}
		if tableDef.getDelay {
			time.Sleep(delayValue)
		}
		item := h.Item.(string)
		rowNumber, _ := strconv.Atoi(item[strings.LastIndex(item, "-")+1:])
		column := populateItem(rowNumber, d.Table)
		return column, nil
	}
}

/// get a hydrate function based on the table def ///
func getHydrate(tableDef *chaosTable) plugin.HydrateFunc {
	if tableDef.hydrateError == FailError {
		return hydrateError
	}
	if tableDef.hydrateError == FailPanic {
		return hydratePanic
	}
	if tableDef.hydrateDelay {
		time.Sleep(delayValue)
	}
	return hydrateColumn
}

//func hydrateError(rowCount int) func(context.Context, *plugin.QueryData, *plugin.HydrateData) (interface{}, error) {
//	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
//		panic("NO")
//		log.Printf("[WARN] hydrateError func rowCount %d row %v", rowCount, h.Item.(map[string]interface{})["id"])
//		if h.Item.(map[string]interface{})["id"] == rowCount {
//			return nil, errors.New("HYDRATE ERROR")
//		}
//		return map[string]interface{}{}, nil
//	}
//}

func hydrateError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(map[string]interface{})["id"].(int)
	if id > 10 {
		return nil, fmt.Errorf("HYDRATE ERROR %d", id)
	}
	return nil, nil
}

func hydratePanic(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	panic("HYDRATE PANIC")
}

func hydrateColumn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	time.Sleep(delayValue)
	key := h.Item.(map[string]interface{})
	columnVal := key["hydrate_column"].(string)
	item := map[string]interface{}{"hydrate_column": columnVal}
	return item, nil
}

//// Transform functions ////
func getTransform(tableDef *chaosTable) t.TransformFunc {
	return func(_ context.Context, d *t.TransformData) (interface{}, error) {
		if tableDef.transformError == FailError {
			return nil, errors.New("TRANSFORM ERROR")
		}
		if tableDef.transformError == FailPanic {
			panic("TRANSFORM PANIC")
		}
		if tableDef.transformDelay {
			time.Sleep(delayValue)
		}
		key := d.HydrateItem.(map[string]interface{})
		columnVal := key["transform_column"].(string)
		return columnVal, nil
	}
}
