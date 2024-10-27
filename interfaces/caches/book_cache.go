package caches_interfaces

import (
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type BookCache interface {
	Cache[responses.BookDetailResponse, responses.ListBooksResponse]
}
