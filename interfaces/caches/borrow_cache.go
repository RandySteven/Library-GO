package caches_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BorrowCache interface {
	Cache[models.Borrow]
}
