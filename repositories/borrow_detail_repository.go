package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type borrowDetailRepository struct {
	dbx repositories_interfaces.DB
}

func (b *borrowDetailRepository) FindByID(ctx context.Context, id uint64) (*models.BorrowDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.BorrowDetail, err error) {
	return utils.FindAll[models.BorrowDetail](ctx, b.dbx(ctx), queries.SelectBorrowDetailQuery)
}

func (b *borrowDetailRepository) Save(ctx context.Context, entity *models.BorrowDetail) (result *models.BorrowDetail, err error) {
	result = entity
	id, err := utils.Save[models.BorrowDetail](ctx, b.dbx(ctx), queries.InsertBorrowDetailQuery, &entity.BorrowID, &entity.BookID, &entity.ReturnedDate)
	if err != nil {
		return nil, err
	}
	result.ID = *id
	return result, nil
}

func (b *borrowDetailRepository) FindByBorrowID(ctx context.Context, borrowID uint64) (results []*models.BorrowDetail, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBorrowDetailByBorrowIDQuery.ToString(), borrowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		borrowDetail := &models.BorrowDetail{}
		err = rows.Scan(
			&borrowDetail.ID, &borrowDetail.BorrowID, &borrowDetail.BookID, &borrowDetail.BorrowedDate, &borrowDetail.ReturnedDate, &borrowDetail.VerifiedReturnDate,
			&borrowDetail.CreatedAt, &borrowDetail.UpdatedAt, &borrowDetail.DeletedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, borrowDetail)
	}
	return results, nil
}

func (b *borrowDetailRepository) FindCurrReturnDate(ctx context.Context) (results []*models.BorrowDetail, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBorrowDetailReturnedDateTodayQuery.ToString())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		borrowDetail := &models.BorrowDetail{}
		err = rows.Scan(
			&borrowDetail.ID,
			&borrowDetail.BorrowID,
			&borrowDetail.BookID,
			&borrowDetail.BorrowedDate,
			&borrowDetail.ReturnedDate,
			&borrowDetail.VerifiedReturnDate,
			&borrowDetail.CreatedAt,
			&borrowDetail.UpdatedAt,
			&borrowDetail.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, borrowDetail)
	}
	return results, nil
}

func (b *borrowDetailRepository) FindByBorrowIDAndBookID(ctx context.Context, borrowID uint64, bookID uint64) (result *models.BorrowDetail, err error) {
	result = &models.BorrowDetail{}
	err = b.dbx(ctx).QueryRowContext(ctx, queries.SelectBorrowDetailByBorrowAndBookQuery.ToString(), borrowID, bookID).
		Scan(
			&result.ID,
			&result.BorrowID,
			&result.BookID,
			&result.BorrowedDate,
			&result.ReturnedDate,
			&result.VerifiedReturnDate,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *borrowDetailRepository) UpdateReturnDateByBorrowIDAndBookID(ctx context.Context, borrowID uint64, bookID uint64) (result *models.BorrowDetail, err error) {
	result = &models.BorrowDetail{}
	_, err = b.dbx(ctx).ExecContext(ctx, queries.UpdateBorrowReturnDateByBorrowAndBookQuery.ToString(), borrowID, bookID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *borrowDetailRepository) FindByBookID(ctx context.Context, bookID uint64) (results []*models.BorrowDetail, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBorrowDetailWithBookIDQuery.ToString(), bookID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		result := &models.BorrowDetail{}
		err = rows.Scan(&result.ID, &result.BorrowID, &result.BookID, &result.BorrowedDate, &result.ReturnedDate, &result.VerifiedReturnDate, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

var _ repositories_interfaces.BorrowDetailRepository = &borrowDetailRepository{}

func newBorrowDetailRepository(dbx repositories_interfaces.DB) *borrowDetailRepository {
	return &borrowDetailRepository{
		dbx: dbx,
	}
}
