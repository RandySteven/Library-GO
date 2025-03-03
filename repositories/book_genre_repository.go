package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type bookGenreRepository struct {
	dbx repositories_interfaces.DB
}

func (b *bookGenreRepository) Save(ctx context.Context, entity *models.BookGenre) (result *models.BookGenre, err error) {
	id, err := utils.Save[models.BookGenre](ctx, b.dbx(ctx), queries.InsertBookGenreQuery, &entity.BookID, &entity.GenreID)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bookGenreRepository) FindBookGenreByBookID(ctx context.Context, bookID uint64) (result []*models.BookGenre, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBookGenreByBookIDQuery.ToString(), bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bookGenre := new(models.BookGenre)
		err = rows.Scan(
			&bookGenre.ID, &bookGenre.BookID, &bookGenre.GenreID, &bookGenre.CreatedAt, &bookGenre.UpdatedAt, &bookGenre.DeletedAt,
			&bookGenre.Book.ID, &bookGenre.Book.Title, &bookGenre.Book.Description, &bookGenre.Book.Image, &bookGenre.Book.Status, &bookGenre.Book.CreatedAt, &bookGenre.Book.UpdatedAt, &bookGenre.Book.DeletedAt,
			&bookGenre.Genre.ID, &bookGenre.Genre.Genre, &bookGenre.Genre.CreatedAt, &bookGenre.Genre.UpdatedAt, &bookGenre.Genre.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, bookGenre)
	}
	return result, nil
}

func (b *bookGenreRepository) FindBookGenreByGenreID(ctx context.Context, genreID uint64) (result []*models.BookGenre, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBookGenreByGenreIDQuery.ToString(), genreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bookGenre := new(models.BookGenre)
		err = rows.Scan(&bookGenre.ID, &bookGenre.BookID, &bookGenre.GenreID, &bookGenre.CreatedAt, &bookGenre.UpdatedAt, &bookGenre.DeletedAt,
			&bookGenre.Book.ID, &bookGenre.Book.Title, &bookGenre.Book.Description, &bookGenre.Book.Image, &bookGenre.Book.Status, &bookGenre.Book.CreatedAt, &bookGenre.Book.UpdatedAt, &bookGenre.Book.DeletedAt,
			&bookGenre.Genre.ID, &bookGenre.Genre.Genre, &bookGenre.Genre.CreatedAt, &bookGenre.Genre.UpdatedAt, &bookGenre.Genre.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, bookGenre)
	}
	return result, nil
}

var _ repositories_interfaces.BookGenreRepository = &bookGenreRepository{}

func newBookGenreRepository(dbx repositories_interfaces.DB) *bookGenreRepository {
	return &bookGenreRepository{
		dbx: dbx,
	}
}
