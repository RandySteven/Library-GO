package algolia

import (
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/search"
)

func NewAlgoliaSearch(config *configs.Config) (*search.APIClient, error) {
	algolia := config.Config.Algolia
	client, err := search.NewClient(algolia.AppID, algolia.ApiKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}
