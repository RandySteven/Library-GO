package schedulers

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
)

type bookScheduler struct {
	cache      caches_interfaces.BookCache
	bookRepo   repositories_interfaces.BookRepository
	ratingRepo repositories_interfaces.RatingRepository
}

func (b *bookScheduler) RefreshBooksCache(ctx context.Context) (err error) {
	err = b.cache.Del(ctx, fmt.Sprintf(enums.BooksKey))
	if err != nil {
		return err
	}

	books, err := b.bookRepo.FindAll(ctx, 0, 0)
	if err != nil {
		return err
	}

	result := []*responses.ListBooksResponse{}
	for _, book := range books {
		result = append(result, &responses.ListBooksResponse{
			ID:        book.ID,
			Title:     book.Title,
			Image:     book.Image,
			Status:    book.Status.ToString(),
			Rating:    0,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
			DeletedAt: book.DeletedAt,
		})
	}

	err = b.cache.SetMultiData(ctx, result)
	if err != nil {
		return err
	}

	return
}

var _ schedulers_interfaces.BookScheduler = &bookScheduler{}

func newBookScheduler(
	bookRepo repositories_interfaces.BookRepository,
	ratingRepo repositories_interfaces.RatingRepository,
	cache caches_interfaces.BookCache) *bookScheduler {
	return &bookScheduler{
		bookRepo:   bookRepo,
		ratingRepo: ratingRepo,
		cache:      cache,
	}
}
