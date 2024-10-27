package caches_interfaces

import "github.com/RandySteven/Library-GO/entities/payloads/responses"

type GenreCache interface {
	MultiDataCache[responses.ListGenresResponse]
}
