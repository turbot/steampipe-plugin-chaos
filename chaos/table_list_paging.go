package chaos

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type Item struct {
	Id   string
	Page int
}

type PagingResponse struct {
	NextPage int
}

type ListResponse struct {
	Items []Item
	Resp  *PagingResponse
}

// paging is 0 based, ranging from 0 to 4
var MaxPage = 5
var noMorePages = -1

// the number of times the function should fail/retry
var failureCount = 200

var errorAfterPages = 3
var retryListError = map[string]int{}
var listErrorString = "retriableError"
var listMutex = &sync.Mutex{}

func listPagingTable() *plugin.Table {
	return &plugin.Table{
		Name:             "chaos_list_paging",
		Description:      "Chaos table to test the list function supporting pagination fails to send results after some pages with a non fatal error",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: listPagingFunction,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "page", Type: proto.ColumnType_INT},
		},
	}
}

func listPagingFunction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	nextPage := 0

	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		items, resp, err := getPage(nextPage)
		return ListResponse{
			Items: items,
			Resp:  resp,
		}, err
	}

	for {
		listResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{shouldRetryError})
		items := listResponse.(ListResponse).Items
		resp := listResponse.(ListResponse).Resp

		if err != nil {
			return nil, err
		}

		for _, i := range items {
			d.StreamListItem(ctx, i)
		}
		if nextPage = resp.NextPage; nextPage == noMorePages {
			break
		}
	}
	return nil, nil
}

// This is a proxy of the API function to fetch the results
func getPage(pageNumber int) ([]Item, *PagingResponse, error) {
	listMutex.Lock()
	errorCount := retryListError[listErrorString]
	retryListError[listErrorString] = errorCount + 1
	listMutex.Unlock()

	if errorCount < failureCount {
		if pageNumber == errorAfterPages {
			return nil, nil, errors.New(listErrorString)
		}
	}

	listMutex.Lock()
	retryListError[listErrorString] = 0
	listMutex.Unlock()

	if pageNumber == MaxPage {
		return nil, nil, errors.New("invalid page")
	}
	var items []Item
	for i := 0; i < 10; i++ {
		items = append(items, Item{Id: fmt.Sprintf("%d_%d", pageNumber, i), Page: pageNumber})
	}
	nextPage := pageNumber + 1
	if nextPage == MaxPage {
		nextPage = noMorePages
	}
	response := PagingResponse{NextPage: nextPage}

	return items, &response, nil

}
