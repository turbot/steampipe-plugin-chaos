package chaos

import (
	"context"
	"errors"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
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
var maxPages = 5
var noMorePages = -1

// the number of times the function should fail/retry
var failureCount = 5
var errorCount = 0
var errorAfterPages = 3

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

	// define function to fetch a page of data - we will pass this to the sdk RetryHydrate call
	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		items, resp, err := getPage(nextPage)
		return ListResponse{
			Items: items,
			Resp:  resp,
		}, err
	}

	for {
		retryResp, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{shouldRetryError})
		listResponse := retryResp.(ListResponse)
		items := listResponse.Items
		resp := listResponse.Resp

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

	if pageNumber == maxPages {
		return nil, nil, errors.New("invalid page")
	}

	// after returning 3 pages, fail 5 times to return the next page before succeeding on the 6th attempt
	if pageNumber == errorAfterPages && errorCount < failureCount {
		errorCount++
		return nil, nil, errors.New(retriableError)
	}

	var items []Item
	for i := 0; i < 10; i++ {
		items = append(items, Item{Id: fmt.Sprintf("%d_%d", pageNumber, i), Page: pageNumber})
	}
	nextPage := pageNumber + 1
	if nextPage == maxPages {
		nextPage = noMorePages
	}
	response := PagingResponse{NextPage: nextPage}

	return items, &response, nil

}
