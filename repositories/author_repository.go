package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"log"
	"strconv"
	"strings"
)

type authorRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (a *authorRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(a.db, a.tx)
}

func (a *authorRepository) Save(ctx context.Context, entity *models.Author) (result *models.Author, err error) {
	id, err := utils.Save[models.Author](ctx, a.Trigger(), ``, &entity.Name, &entity.Nationality)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return
}

func (a *authorRepository) FindByID(ctx context.Context, id uint64) (result *models.Author, err error) {
	result = &models.Author{}
	err = utils.FindByID[models.Author](ctx, a.Trigger(), queries.SelectAuthorByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *authorRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) BeginTx(ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	a.tx = tx
	return nil
}

func (a *authorRepository) CommitTx(ctx context.Context) error {
	return a.tx.Commit()
}

func (a *authorRepository) RollbackTx(ctx context.Context) error {
	return a.tx.Rollback()
}

func (a *authorRepository) SetTx(tx *sql.Tx) {
	a.tx = tx
}

func (a *authorRepository) GetTx(ctx context.Context) *sql.Tx {
	return a.tx
}

func (a *authorRepository) FindSelectedAuthorsByID(ctx context.Context, ids []uint64) (result []*models.Author, err error) {
	queryIn := ` WHERE id IN (%s)`
	wildCards := []string{}
	for _, id := range ids {
		wildCards = append(wildCards, strconv.Itoa(int(id)))
	}
	wildCardStr := strings.Join(wildCards, ",")
	queryIn = fmt.Sprintf(queryIn, wildCardStr)
	selectStr := queries.SelectAuthorQuery.ToString() + queryIn
	log.Println(selectStr)
	rows, err := a.Trigger().QueryContext(ctx, selectStr)
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

func newAuthorRepository(db *sql.DB) *authorRepository {
	return &authorRepository{
		db: db,
	}
}
