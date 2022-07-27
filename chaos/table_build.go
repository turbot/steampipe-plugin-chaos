package chaos

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
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
	listBuildConfig      *listBuildConfig
	getBuildConfig       *getBuildConfig
	hydrateBuildConfig   *hydrateBuildConfig
	transformBuildConfig *transformBuildConfig
	name                 string
	description          string
	columnCount          int
	cache                *plugin.TableCacheOptions
}

const defaultColumnCount = 10
const defaultRowCount = 10
const delayValue = 5 * time.Second

func buildTable(tableDef *chaosTable) *plugin.Table {
	return &plugin.Table{
		Name:        tableDef.name,
		Description: tableDef.description,
		List: &plugin.ListConfig{
			Hydrate: buildListHydrate(tableDef.listBuildConfig),
			//KeyColumns: []*plugin.KeyColumn{
			//	{
			//		Name:    "id",
			//		Require: plugin.Optional,
			//	},
			//},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    buildGetHydrate(tableDef.getBuildConfig),
		},
		Columns: buildColumns(tableDef),
		Cache:   tableDef.cache,
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
	return columns
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

		//var item map[string]interface{}
		//item = populateItem(0, d.Table)

		//for i := 0; i < 50000; i++ {
		//	if i%10000 == 0 {
		//		log.Printf("[WARN] st")
		//		runtime.GC()
		//	}
		//	d.FOO(ctx, populateItem(0, d.Table))
		//}
		for i := 0; i < buildConfig.rowCount; i++ {
			//time.Sleep(1 * time.Millisecond)
			//listErrorRows is the number of rows to return successfully before raising an error
			//if we stream that many rows, let's raise an error
			//if i == buildConfig.listRowsBeforeError {
			//	switch buildConfig.listError {
			//	case RetryableError:
			//		// failureCount is the number of times the error occurs before we succeed
			//		if listTableErrorCount < buildConfig.failureCount {
			//			log.Printf("[TRACE] return retriable error")
			//			listTableErrorCount++
			//			return nil, errors.New(RetryableError)
			//		}
			//
			//		// if we have failed 'failureCount' times, reset listTableErrorCount and fall through to return item
			//		log.Printf("[TRACE] retry worked - no error")
			//		listTableErrorCount = 0
			//	case IgnorableError:
			//		log.Printf("[TRACE] return ignorable error")
			//		return nil, errors.New(IgnorableError)
			//	case FailError:
			//		log.Printf("[TRACE] return fatal  error")
			//		return nil, errors.New(FatalError)
			//	case FailPanic:
			//		panic(FailPanic)
			//
			//	}
			//
			//}
			item := populateItem(i, d.Table)

			if d.QueryStatus.StreamingComplete {
				//break
				item = nil
			}

			//if i%1000 == 0 {
			//	log.Printf("[WARN] streamed 1000")
			//}
			d.StreamLeafListItem(ctx, item)
		}

		log.Printf("[WARN] END STREAMING. pid %d, cols %v", os.Getpid(), d.QueryContext.Columns)

		//time.Sleep(10 * time.Second)
		log.Printf("[WARN] END STREAMING. RET")

		//runtime.GC()
		return nil, nil
	}
}
