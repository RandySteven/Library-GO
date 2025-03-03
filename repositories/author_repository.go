package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"log"
)

type authorRepository struct {
	dbx repositories_interfaces.DB
}

//func (a *authorRepository) Trigger() repositories_interfaces.Trigger {
//	return utils.InitTrigger(a.db, a.tx)
//}

func (a *authorRepository) Save(ctx context.Context, entity *models.Author) (result *models.Author, err error) {
	id, err := utils.Save[models.Author](ctx, a.dbx(ctx), ``, &entity.Name, &entity.Nationality)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return
}

func (a *authorRepository) FindByID(ctx context.Context, id uint64) (result *models.Author, err error) {
	result = &models.Author{}
	err = utils.FindByID[models.Author](ctx, a.dbx(ctx), queries.SelectAuthorByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *authorRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Author, error) {
	return utils.FindAll[models.Author](ctx, a.dbx(ctx), queries.SelectAuthorQuery)
}

func (a *authorRepository) FindSelectedAuthorsByID(ctx context.Context, ids []uint64) (result []*models.Author, err error) {
	selectStr := utils.SelectIdIn(queries.SelectAuthorQuery, ids)
	log.Println(selectStr)
	rows, err := a.dbx(ctx).QueryContext(ctx, selectStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		author := new(models.Author)
		err = rows.Scan(&author.ID, &author.Name, &author.Nationality, &author.CreatedAt, &author.UpdatedAt, &author.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, author)
	}
	return result, nil
}

var _ repositories_interfaces.AuthorRepository = &authorRepository{}

func newAuthorRepository(dbx repositories_interfaces.DB) *authorRepository {
	return &authorRepository{
		dbx: dbx,
	}
}
