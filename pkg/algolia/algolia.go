package algolia

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
