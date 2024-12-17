package searchers_interfaces

import "context"

type Searcher[T any] interface {
	SearchList(ctx context.Context, keyword string) (result []*T, err error)
}
