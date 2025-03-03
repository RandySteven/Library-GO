package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type userGenreRepository struct {
	dbx repositories_interfaces.DB
}

func (u *userGenreRepository) Save(ctx context.Context, entity *models.UserGenre) (*models.UserGenre, error) {
	id, err := utils.Save[models.UserGenre](ctx, u.dbx(ctx), queries.InsertUserGenreQuery, entity.UserID, entity.GenreID)
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
