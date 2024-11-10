package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type RatingRepository interface {
	Saver[models.Rating]
	UnitOfWork
	FindRatingForBook(ctx context.Context, bookId uint64) (result *models.Rating, err error)
}
