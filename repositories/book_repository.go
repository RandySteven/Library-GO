package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type bookRepository struct {
	dbx repositories_interfaces.DB
}

func (b *bookRepository) Save(ctx context.Context, entity *models.Book) (result *models.Book, err error) {
	id, err := utils.Save[models.Book](ctx, b.dbx(ctx), queries.InsertBookQuery, entity.Title, entity.Description, entity.Image, entity.Status)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bookRepository) FindByID(ctx context.Context, id uint64) (result *models.Book, err error) {
	result = &models.Book{}
	err = utils.FindByID[models.Book](ctx, b.dbx(ctx), queries.SelectBookByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *bookRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Book, error) {
	if skip == 0 && take == 0 {
		return utils.FindAll[models.Book](ctx, b.dbx(ctx), queries.SelectBooksQuery)
	}
	if skip == 1 {
		skip = 0
	} else {
		skip = skip*take - take
	}
	books := []*models.Book{}
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBookPaginateQuery.ToString(), take, skip)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.Image,
			&book.Status, &book.PDFFile, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b *bookRepository) FindSelectedBooksId(ctx context.Context, ids []uint64) (results []*models.Book, err error) {
	selectStr := utils.SelectIdIn(queries.SelectBooksQuery, ids)
	rows, err := b.dbx(ctx).QueryContext(ctx, selectStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := new(models.Book)
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image, &book.Status, &book.PDFFile, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, book)
	}
	return results, nil
}

func (b *bookRepository) FindBookStatus(ctx context.Context, id uint64, status enums.BookStatus) (isExist bool, err error) {
	result := &models.Book{}
	err = b.dbx(ctx).QueryRowContext(ctx, queries.SelectBookAndStatus.ToString(), id, status).Scan(
		&result.ID, &result.Title, &result.Description, &result.Image, &result.Status, &result.PDFFile, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return false, err
	}
	return result.Status == status, nil
}

func (b *bookRepository) UpdateBookStatus(ctx context.Context, id uint64, status enums.BookStatus) error {
	_, err := b.dbx(ctx).ExecContext(ctx, queries.UpdateBookStatusQuery.ToString(), status, id)
	if err != nil {
		return err
	}
	return nil
}

var _ repositories_interfaces.BookRepository = &bookRepository{}

func newBookRepository(dbx repositories_interfaces.DB) *bookRepository {
	return &bookRepository{
		dbx: dbx,
	}
}
