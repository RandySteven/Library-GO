package searchers_interfaces

import (
	"github.com/RandySteven/Library-GO/entities/indexes"
)

type BookSearcher interface {
	Searcher[indexes.BookIndex]
}
