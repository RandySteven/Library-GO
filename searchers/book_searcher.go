package searchers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/RandySteven/Library-GO/entities/indexes"
	searchers_interfaces "github.com/RandySteven/Library-GO/interfaces/searchers"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
)

type bookSearcher struct {
	elastic *elasticsearch.Client
}

func (b *bookSearcher) SaveSearch(ctx context.Context, content *indexes.BookIndex) (result *indexes.BookIndex, err error) {
	jsonData, err := json.Marshal(&content)
	if err != nil {
		return nil, err
	}
	_, err = b.elastic.Indices.Create("books")
	if err != nil {
		return nil, err
	}
	_, err = b.elastic.Index("books", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (b *bookSearcher) SearchList(ctx context.Context, keyword string) (result []*indexes.BookIndex, err error) {
	query := `
		{
			"query": {
				"query_string": {
					"fields": []
				}
			}
		}
	`
	search, err := b.elastic.Search(
		b.elastic.Search.WithIndex("books"),
		b.elastic.Search.WithBody(strings.NewReader(query)))
	if err != nil {
		return nil, err
	}
	search.Body.Close()
	return nil, nil
}

var _ searchers_interfaces.BookSearcher = &bookSearcher{}

func newBookSearcher(elastic *elasticsearch.Client) *bookSearcher {
	return &bookSearcher{
		elastic: elastic,
	}
}
