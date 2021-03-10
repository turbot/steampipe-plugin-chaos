package chaos

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

type APIResult struct {
	Item     *[]Item
	response *pagingResponse
}

type Item struct {
	id   string
	page int
}

type pagingResponse struct {
	NextPage int
}

// paging is 0 based, ranging from 0 to 4
var MaxPage = 5
var noMorePages = -1

func listPagingTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaos_list_paging",
		Description: "Chaos table to test the default transform functionality from specified json tags",
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
	for {
		items, resp, err := getPage(nextPage)
		if err != nil {
			return nil, err
		}

		for _, i := range items {
			log.Printf("[ERROR] ITEM VALUES======> %v", i)
			d.StreamListItem(ctx, i)
		}
		if nextPage = resp.NextPage; nextPage == noMorePages {
			break
		}
	}
	return nil, nil
}

// This is a proxy of the API function to fetch the results
func getPage(pageNumber int) ([]Item, *pagingResponse, error) {
	if pageNumber == MaxPage {
		return nil, nil, errors.New("invalid page")
	}
	var items []Item
	for i := 0; i < 100; i++ {
		items = append(items, Item{id: fmt.Sprintf("%d_%d", pageNumber, i), page: pageNumber})
	}
	nextPage := pageNumber + 1
	if nextPage == MaxPage {
		nextPage = noMorePages
	}
	response := pagingResponse{NextPage: nextPage}

	return items, &response, nil

}
