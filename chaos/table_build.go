package chaos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	t "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type FailType string

const (
	FailNone       FailType = "None"
	FailError               = "Error"
	FailPanic               = "Panic"
	FatalError              = "fatalError"
	RetryableError          = "retriableError"
	IgnorableError          = "resourceNotFound"
)

type chaosTable struct {
	listBuildConfig  *listBuildConfig
	getBuildConfig   *getBuildConfig
	name             string
	description      string
	columnCount      int
	hydrateError     FailType
	hydrateDelay     bool
	itemFromKeyError FailType
	transformError   FailType
	transformDelay   bool
	errorType        FailType
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
			Hydrate: buildListHydrate(tableDef.listBuildConfig),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    buildGetHydrate(tableDef.getBuildConfig),
		},
		Columns: buildColumns(tableDef),
	}

}

func buildColumns(tableDef *chaosTable) []*plugin.Column {
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
		Hydrate: buildHydrate(tableDef),
	}
	transformColumn := &plugin.Column{
		Name:      "transform_column",
		Type:      proto.ColumnType_STRING,
		Transform: t.From(buildTransform(tableDef)),
	}
	columns = append(columns, transformColumn)
	if tableDef.hydrateDelay || tableDef.hydrateError == FailError || tableDef.hydrateError == FailPanic {
		columns = append(columns, hydrateColumn)
	}
	return columns
}

/// get a hydrate function based on the table def ///
func buildHydrate(tableDef *chaosTable) plugin.HydrateFunc {
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
func buildTransform(tableDef *chaosTable) t.TransformFunc {
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

// factory function which returns a list call with the behaviour determined by the list config
func buildListHydrate(buildConfig *listBuildConfig) plugin.HydrateFunc {
	if buildConfig == nil {
		buildConfig = &listBuildConfig{
			rowCount:    defaultRowCount,
			columnCount: defaultColumnCount,
		}
	}

	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		// if listDelay is specified, sleep
		if buildConfig.listDelay {
			time.Sleep(delayValue)
		}

		log.Printf("[TRACE] ABOUT TO START STREAMING. pid %d, cols %v", os.Getpid(), d.QueryContext.Columns)

		for i := 0; i < buildConfig.rowCount; i++ {
			//log.Printf("[WARN] ROW LOOP streamed %d.   pid %d, cols %v", os.Getpid(), d.QueryContext.Columns)
			time.Sleep(50 * time.Millisecond)
			// listErrorRows is the number of rows to return successfully before raising an error
			// if we stream that many rows, let's raise an error
			if i == buildConfig.listErrorRows {
				switch buildConfig.listError {
				case RetryableError:
					// failureCount is the number of times the error occurs before we succeed
					if listTableErrorCount <= buildConfig.failureCount {
						listTableErrorCount++
						return nil, errors.New(RetryableError)
					}
					// if we have failed 'failureCount' times, reset listTableErrorCount and fall through to return item
					listTableErrorCount = 0
				case IgnorableError:
					return nil, errors.New(IgnorableError)
				case FailError:
					return nil, errors.New(FatalError)
				case FailPanic:
					panic(FailPanic)

				}

			}
			item := populateItem(i, d.Table)
			d.StreamListItem(ctx, item)

		}

		log.Printf("[TRACE] END  STREAMING. pid %d, cols %v", os.Getpid(), d.QueryContext.Columns)
		return nil, nil
	}
}
