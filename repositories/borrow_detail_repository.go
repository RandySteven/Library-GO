package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type borrowDetailRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *borrowDetailRepository) Save(ctx context.Context, entity *models.BorrowDetail) (result *models.BorrowDetail, err error) {
	result = entity
	id, err := utils.Save[models.BorrowDetail](ctx, b.InitTrigger(), queries.InsertBorrowDetailQuery, &entity.BorrowID, &entity.BookID, &entity.ReturnedDate)
	if err != nil {
		return nil, err
	}
	result.ID = *id
	return result, nil
}

func (b *borrowDetailRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = b.db
	if b.tx != nil {
		trigger = b.tx
	}
	return trigger
}

func (b *borrowDetailRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.tx = tx
	return nil
}

func (b *borrowDetailRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *borrowDetailRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *borrowDetailRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *borrowDetailRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *borrowDetailRepository) FindByBorrowID(ctx context.Context, borrowID uint64) (results []*models.BorrowDetail, err error) {
	rows, err := b.InitTrigger().QueryContext(ctx, queries.SelectBorrowDetailByBorrowIDQuery.ToString(), borrowID)
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
	rows, err := b.InitTrigger().QueryContext(ctx, queries.SelectBorrowDetailReturnedDateTodayQuery.ToString())
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
	err = b.InitTrigger().QueryRowContext(ctx, queries.SelectBorrowDetailByBorrowAndBookQuery.ToString(), borrowID, bookID).
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
	_, err = b.InitTrigger().ExecContext(ctx, queries.UpdateBorrowReturnDateByBorrowAndBookQuery.ToString(), borrowID, bookID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ repositories_interfaces.BorrowDetailRepository = &borrowDetailRepository{}

func newBorrowDetailRepository(db *sql.DB) *borrowDetailRepository {
	return &borrowDetailRepository{
		db: db,
	}
}
