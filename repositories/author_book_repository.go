package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type authorBookRepository struct {
	db repositories_interfaces.DB
}

func (a *authorBookRepository) Save(ctx context.Context, entity *models.AuthorBook) (result *models.AuthorBook, err error) {
	id, err := utils.Save[models.AuthorBook](ctx, a.db(ctx), queries.InsertAuthorBookQuery, &entity.AuthorID, &entity.BookID)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (a *authorBookRepository) FindAuthorBookByBookID(ctx context.Context, bookID uint64) (result []*models.AuthorBook, err error) {
	rows, err := a.db(ctx).QueryContext(ctx, queries.SelectAuthorBookByBookIDQuery.ToString(), bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		authorBook := new(models.AuthorBook)
		err = rows.Scan(
			&authorBook.ID, &authorBook.AuthorID, &authorBook.BookID, &authorBook.CreatedAt, &authorBook.UpdatedAt, &authorBook.DeletedAt,
			&authorBook.Book.ID, &authorBook.Book.Title, &authorBook.Book.Description, &authorBook.Book.Image, &authorBook.Book.Status, &authorBook.Book.CreatedAt, &authorBook.Book.UpdatedAt, &authorBook.Book.DeletedAt,
			&authorBook.Author.ID, &authorBook.Author.Name, &authorBook.Author.Nationality, &authorBook.Author.CreatedAt, &authorBook.Author.UpdatedAt, &authorBook.Author.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, authorBook)
	}
	return result, nil
}

var _ repositories_interfaces.AuthorBookRepository = &authorBookRepository{}

func newAuthorBookRepository(db repositories_interfaces.DB) *authorBookRepository {
	return &authorBookRepository{
		db: db,
	}
}
