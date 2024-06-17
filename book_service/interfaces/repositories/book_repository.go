package repositories_interfaces

import (
	"github.com/RandySteven/Library-GO/book_service/entities/models"
	"github.com/RandySteven/Library-GO/repositories"
)

type (
	BookRepository interface {
		repositories.Repository[models.Book]
	}
)
