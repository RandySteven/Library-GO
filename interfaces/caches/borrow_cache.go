package caches_interfaces

import (
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type BorrowCache interface {
	Cache[responses.BorrowDetailResponse, responses.BorrowListResponse]
}
