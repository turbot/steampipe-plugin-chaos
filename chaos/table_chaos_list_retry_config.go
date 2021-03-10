package chaos

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	log "log"
// 	"sync"

// 	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
// 	"github.com/turbot/steampipe-plugin-sdk/plugin"
// )

// // var retryListError = map[string]int{}
// // var listErrorString = "retriableError"
// // var listMutex = &sync.Mutex{}

// func listRetryConfigTable() *plugin.Table {
// 	return &plugin.Table{
// 		Name:        "chaos_list_retry_config",
// 		Description: "Chaos table to test the List function with Retry config in case of non fatal error",
// 		List: &plugin.ListConfig{
// 			Hydrate: listRetryConfigList,
// 			RetryConfig: &plugin.RetryConfig{
// 				ShouldRetryError: shouldRetryError,
// 			},
// 		},
// 		Columns: []*plugin.Column{
// 			{Name: "retry_column", Type: proto.ColumnType_STRING},
// 		},
// 	}
// }

// func listRetryConfigList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	log.Println("[INFO] INSIDE LIST CALL")

// 	var failureCount = 2

// 	listMutex.Lock()
// 	errorCount := retryListError[listErrorString]
// 	retryListError[listErrorString] = errorCount + 1
// 	listMutex.Unlock()

// 	if errorCount < failureCount {
// 		return nil, errors.New(listErrorString)
// 	}
// 	listMutex.Lock()
// 	retryListError[listErrorString] = 0
// 	listMutex.Unlock()
// 	for i := 0; i < 5; i++ {
// 		columnValue := fmt.Sprintf("%s-%v", "columnValue", i)
// 		item := map[string]interface{}{"retry_column": columnValue}
// 		d.StreamListItem(ctx, item)
// 	}
// 	return nil, nil
// }
