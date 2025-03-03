package repositories

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"strconv"
	"strings"
)

type borrowRepository struct {
	dbx repositories_interfaces.DB
}

func (b *borrowRepository) Save(ctx context.Context, entity *models.Borrow) (result *models.Borrow, err error) {
	id, err := utils.Save[models.Borrow](ctx, b.dbx(ctx), queries.InsertBorrowQuery, &entity.UserID, &entity.BorrowReference)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (b *borrowRepository) FindByID(ctx context.Context, id uint64) (result *models.Borrow, err error) {
	result = &models.Borrow{}
	err = utils.FindByID[models.Borrow](ctx, b.dbx(ctx), queries.SelectBorrowByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *borrowRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Borrow, error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowRepository) FindByReferenceID(ctx context.Context, referenceID string) (result *models.Borrow, err error) {
	result = &models.Borrow{}
	err = b.dbx(ctx).QueryRowContext(ctx, queries.SelectBorrowByReference.ToString(), referenceID).
		Scan(
			&result.ID,
			&result.UserID,
			&result.BorrowReference,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *borrowRepository) FindByUserId(ctx context.Context, userId uint64) (result []*models.Borrow, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBorrowUserIdQuery.ToString(), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		borrow := &models.Borrow{}
		err = rows.Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BorrowReference,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, borrow)
	}

	return result, nil
}

func (b *borrowRepository) FindByMultipleBorrowID(ctx context.Context, borrowIDs []uint64) (result []*models.Borrow, err error) {
	queryIn := ` WHERE id IN (%s)`
	wildCards := []string{}
	for _, id := range borrowIDs {
		wildCards = append(wildCards, strconv.Itoa(int(id)))
	}
	wildCardStr := strings.Join(wildCards, ",")
	queryIn = fmt.Sprintf(queryIn, wildCardStr)
	selectStr := queries.SelectBorrowsQueryWithUser.ToString() + queryIn

	rows, err := b.dbx(ctx).QueryContext(ctx, selectStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := &models.Borrow{}
		err = rows.Scan(&res.ID, &res.UserID, &res.User.Name, &res.User.Email)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

var _ repositories_interfaces.BorrowRepository = &borrowRepository{}

func newBorrowRepository(dbx repositories_interfaces.DB) *borrowRepository {
	return &borrowRepository{
		dbx: dbx,
	}
}
