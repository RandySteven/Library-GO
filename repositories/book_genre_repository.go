package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"log"
)

type bookGenreRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bookGenreRepository) Save(ctx context.Context, entity *models.BookGenre) (result *models.BookGenre, err error) {
	log.Println("book id ", entity.BookID)
	log.Println("genre id ", entity.GenreID)
	id, err := utils.Save[models.BookGenre](ctx, b.InitTrigger(), queries.InsertBookGenreQuery, &entity.BookID, &entity.GenreID)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bookGenreRepository) FindByID(ctx context.Context, id uint64) (result *models.BookGenre, err error) {
	return
}

func (b *bookGenreRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.BookGenre, error) {
	//TODO implement me
	return nil, nil
}

func (b *bookGenreRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	return
}

func (b *bookGenreRepository) Update(ctx context.Context, entity *models.BookGenre) (result *models.BookGenre, err error) {
	//TODO implement me
	return
}

func (b *bookGenreRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = b.db
	if b.tx != nil {
		trigger = b.tx
	}
	return trigger
}

func (b *bookGenreRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.tx = tx
	return nil
}

func (b *bookGenreRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *bookGenreRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *bookGenreRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *bookGenreRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *bookGenreRepository) FindBookGenreByBookID(ctx context.Context, bookID uint64) (result []*models.BookGenre, err error) {
	rows, err := b.InitTrigger().QueryContext(ctx, queries.SelectBookGenreByBookIDQuery.ToString(), bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bookGenre := new(models.BookGenre)
		err = rows.Scan(&bookGenre.ID, &bookGenre.BookID, &bookGenre.GenreID, &bookGenre.CreatedAt, &bookGenre.UpdatedAt, &bookGenre.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, bookGenre)
	}
	return result, nil
}

func (b *bookGenreRepository) FindBookGenreByGenreID(ctx context.Context, genreID uint64) (result []*models.BookGenre, err error) {
	rows, err := b.InitTrigger().QueryContext(ctx, queries.SelectBookGenreByGenreIDQuery.ToString(), genreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bookGenre := new(models.BookGenre)
		err = rows.Scan(&bookGenre.ID, &bookGenre.BookID, &bookGenre.GenreID, &bookGenre.CreatedAt, &bookGenre.UpdatedAt, &bookGenre.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, bookGenre)
	}
	return result, nil
}

var _ repositories_interfaces.BookGenreRepository = &bookGenreRepository{}

func newBookGenreRepository(db *sql.DB) *bookGenreRepository {
	return &bookGenreRepository{
		db: db,
	}
}
