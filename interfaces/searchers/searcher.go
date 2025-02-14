package searchers_interfaces

import "context"

type Searcher[T any] interface {
	SaveSearch(ctx context.Context, content *T) (result *T, err error)
	SearchList(ctx context.Context, keyword string) (result []*T, err error)
}
