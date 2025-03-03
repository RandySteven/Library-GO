package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type ratingRepository struct {
	dbx repositories_interfaces.DB
}

func (r *ratingRepository) Save(ctx context.Context, entity *models.Rating) (result *models.Rating, err error) {
	id, err := utils.Save[models.Rating](ctx, r.dbx(ctx), queries.InsertIntoRatingQuery, &entity.BookID, &entity.UserID, &entity.Score)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *ratingRepository) FindRatingForBook(ctx context.Context, bookId uint64) (result *models.Rating, err error) {
	result = &models.Rating{}
	err = r.dbx(ctx).QueryRowContext(ctx, queries.SelectRatingValue.ToString(), bookId).
		Scan(&result.BookID, &result.Score)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ratingRepository) FindSortedLimitRating(ctx context.Context, sorted string, limit uint64) (result []*models.Rating, err error) {
	rows, err := r.dbx(ctx).QueryContext(ctx, queries.SelectRatingSortedLimitQuery.ToString(), limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rating := &models.Rating{}
		err = rows.Scan(&rating.BookID, &rating.Score,
			&rating.Book.Title, &rating.Book.Image, &rating.Book.Status,
			&rating.Book.CreatedAt, &rating.Book.UpdatedAt, &rating.Book.DeletedAt)
		result = append(result, rating)
	}
	return result, nil
}

var _ repositories_interfaces.RatingRepository = &ratingRepository{}

func newRatingRepository(dbx repositories_interfaces.DB) *ratingRepository {
	return &ratingRepository{
		dbx: dbx,
	}
}
