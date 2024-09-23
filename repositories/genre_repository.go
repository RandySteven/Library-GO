package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type genreRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (g *genreRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = g.db
	if g.tx != nil {
		trigger = g.tx
	}
	return trigger
}

func (g *genreRepository) Save(ctx context.Context, entity *models.Genre) (result *models.Genre, err error) {
	id, err := utils.Save[models.Genre](ctx, g.InitTrigger(), queries.InsertGenreQuery, &entity.Genre)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (g *genreRepository) FindByID(ctx context.Context, id uint64) (result *models.Genre, err error) {
	result = &models.Genre{}
	err = utils.FindByID[models.Genre](ctx, g.InitTrigger(), queries.SelectGenreByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *genreRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.Genre, err error) {
	return utils.FindAll[models.Genre](ctx, g.InitTrigger(), queries.SelectGenresQuery)
}

func (g *genreRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (g *genreRepository) Update(ctx context.Context, entity *models.Genre) (result *models.Genre, err error) {
	//TODO implement me
	panic("implement me")
}

func (g *genreRepository) BeginTx(ctx context.Context) error {
	tx, err := g.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	g.tx = tx
	return nil
}

func (g *genreRepository) CommitTx(ctx context.Context) error {
	return g.tx.Commit()
}

func (g *genreRepository) RollbackTx(ctx context.Context) error {
	return g.tx.Rollback()
}

func (g *genreRepository) SetTx(tx *sql.Tx) {
	g.tx = tx
}

func (g *genreRepository) GetTx(ctx context.Context) *sql.Tx {
	return g.tx
}

var _ repositories_interfaces.GenreRepository = &genreRepository{}

func newGenreRepository(db *sql.DB) *genreRepository {
	return &genreRepository{
		db: db,
	}
}
