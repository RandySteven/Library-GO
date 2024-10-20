package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type ratingRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *ratingRepository) Save(ctx context.Context, entity *models.Rating) (result *models.Rating, err error) {
	id, err := utils.Save[models.Rating](ctx, r.InitTrigger(), queries.InsertIntoRatingQuery, &entity.BookID, &entity.UserID, &entity.Score)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *ratingRepository) FindByID(ctx context.Context, id uint64) (result *models.Rating, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *ratingRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Rating, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ratingRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *ratingRepository) Update(ctx context.Context, entity *models.Rating) (result *models.Rating, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *ratingRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = r.db
	if r.tx != nil {
		trigger = r.tx
	}
	return trigger
}

func (r *ratingRepository) BeginTx(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *ratingRepository) CommitTx(ctx context.Context) error {
	return r.tx.Commit()
}

func (r *ratingRepository) RollbackTx(ctx context.Context) error {
	return r.tx.Rollback()
}

func (r *ratingRepository) SetTx(tx *sql.Tx) {
	r.tx = tx
}

func (r *ratingRepository) GetTx(ctx context.Context) *sql.Tx {
	return r.tx
}

func (r *ratingRepository) FindRatingForBook(ctx context.Context, bookId uint64) (result *models.Rating, err error) {
	result = &models.Rating{}
	err = r.InitTrigger().QueryRowContext(ctx, queries.SelectRatingValue.ToString(), bookId).
		Scan(&result.BookID, &result.Score)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ repositories_interfaces.RatingRepository = &ratingRepository{}

func newRatingRepository(db *sql.DB) *ratingRepository {
	return &ratingRepository{
		db: db,
	}
}
