package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type genreRepository struct {
	dbx repositories_interfaces.DB
}

//func (g *genreRepository) .dbx(ctx) repositories_interfaces.Trigger {
//	return utils.InitTrigger(g.db, g.tx)
//}

func (g *genreRepository) Save(ctx context.Context, entity *models.Genre) (result *models.Genre, err error) {
	id, err := utils.Save[models.Genre](ctx, g.dbx(ctx), queries.InsertGenreQuery, &entity.Genre)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (g *genreRepository) FindByID(ctx context.Context, id uint64) (result *models.Genre, err error) {
	result = &models.Genre{}
	err = utils.FindByID[models.Genre](ctx, g.dbx(ctx), queries.SelectGenreByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *genreRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.Genre, err error) {
	return utils.FindAll[models.Genre](ctx, g.dbx(ctx), queries.SelectGenresQuery)
}

func (g *genreRepository) FindSelectedGenresByID(ctx context.Context, ids []uint64) (result []*models.Genre, err error) {
	selectStr := utils.SelectIdIn(queries.SelectGenresQuery, ids)
	rows, err := g.dbx(ctx).QueryContext(ctx, selectStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		genre := new(models.Genre)
		err = rows.Scan(&genre.ID, &genre.Genre, &genre.CreatedAt, &genre.UpdatedAt, &genre.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, genre)
	}
	return result, nil
}

var _ repositories_interfaces.GenreRepository = &genreRepository{}

func newGenreRepository(dbx repositories_interfaces.DB) *genreRepository {
	return &genreRepository{
		dbx: dbx,
	}
}
