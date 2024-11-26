package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type ratingUsecase struct {
	ratingRepo repositories_interfaces.RatingRepository
}

func (r *ratingUsecase) SubmitBookRating(ctx context.Context, request *requests.RatingRequest) (result *responses.RatingResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)

	rating, err := r.ratingRepo.Save(ctx, &models.Rating{
		UserID: userId,
		BookID: request.BookID,
		Score:  request.Rating,
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save rating`, err)
	}

	result = &responses.RatingResponse{
		ID: rating.ID,
		User: struct {
			ID uint64 `json:"id"`
		}{
			ID: userId,
		},
		Book: struct {
			ID uint64 `json:"id"`
		}{
			ID: rating.BookID,
		},
		Rating: rating.Score,
	}

	return result, nil
}

func (r *ratingUsecase) RatingBooksFilter(ctx context.Context, request *requests.RatingFilter) (results []*responses.SortedBookRatings, customErr *apperror.CustomError) {
	ratingBooks, err := r.ratingRepo.FindSortedLimitRating(ctx, request.Order, request.Limit)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get rating books sorted`, err)
	}

	for _, ratingBook := range ratingBooks {
		results = append(results, &responses.SortedBookRatings{
			Rating: ratingBook.Score,
			Book: struct {
				ID    uint64 `json:"id"`
				Title string `json:"title"`
				Image string `json:"image"`
			}{ID: ratingBook.BookID, Title: ratingBook.Book.Title, Image: ratingBook.Book.Image},
		})
	}

	return
}

var _ usecases_interfaces.RatingUsecase = &ratingUsecase{}

func newRatingUsecase(ratingRepo repositories_interfaces.RatingRepository) *ratingUsecase {
	return &ratingUsecase{
		ratingRepo: ratingRepo,
	}
}
