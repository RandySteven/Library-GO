package elastics_client

import (
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticClient struct {
	client *elasticsearch.Client
}

func NewElasticClient(config *configs.Config) (*ElasticClient, error) {
	elastic := config.Config.ElasticSearch
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("https://%s:%s@%s:%s",
				elastic.AccessKey,
				elastic.SecretKey,
				elastic.Host,
				elastic.Port,
			),
		},
	})
	if err != nil {
		return nil, err
	}

	return &ElasticClient{
		client: es,
	}, nil
}

func (c *ElasticClient) Ping() {
	c.client.Ping.WithHuman()
}

func (c *ElasticClient) Client() *elasticsearch.Client {
	return c.client
}
