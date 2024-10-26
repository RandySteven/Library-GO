package caches_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BookCache interface {
	Cache[models.Book]
}
