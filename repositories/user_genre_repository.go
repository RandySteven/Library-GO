package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type userGenreRepository struct {
	db *sql.DB
}

func (u *userGenreRepository) Save(ctx context.Context, entity *models.UserGenre) (*models.UserGenre, error) {
	id, err := utils.Save[models.UserGenre](ctx, u.db, queries.InsertUserGenreQuery, entity.UserID, entity.GenreID)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (u *userGenreRepository) FindByID(ctx context.Context, id uint64) (*models.UserGenre, error) {
	panic("implement me")
}

func (u *userGenreRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.UserGenre, error) {
	panic("implement me")
}

var _ repositories_interfaces.UserGenreRepository = &userGenreRepository{}
