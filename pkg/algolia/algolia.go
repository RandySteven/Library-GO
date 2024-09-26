package algolia_client

import (
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/search"
)

type AlgoliaAPISearchClient struct {
	search *search.APIClient
}

func NewAlgoliaSearch(config *configs.Config) (*AlgoliaAPISearchClient, error) {
	algolia := config.Config.Algolia
	client, err := search.NewClient(algolia.AppID, algolia.ApiKey)
	if err != nil {
		return nil, err
	}
	return &AlgoliaAPISearchClient{
		search: client,
	}, nil
}

func (a *AlgoliaAPISearchClient) SearchClient() *search.APIClient {
	return a.search
}

func (a *AlgoliaAPISearchClient) SaveObject(indexName string, record map[string]any) (resp any, err error) {
	resp, err = a.search.SaveObject(
		a.search.NewApiSaveObjectRequest(indexName, record))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlgoliaAPISearchClient) Search(indexName string, query string) ([]search.SearchResult, error) {
	resp, err := a.search.Search(
		a.search.NewApiSearchRequest(
			search.NewEmptySearchMethodParams().SetRequests(
				[]search.SearchQuery{
					*search.SearchForHitsAsSearchQuery(
						search.NewEmptySearchForHits().SetIndexName(indexName).SetQuery(query),
					),
				},
			),
		),
	)
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}
