package repositories

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"log"
	"strconv"
	"strings"
)

type bagRepository struct {
	dbx repositories_interfaces.DB
}

func (b *bagRepository) CheckBagExists(ctx context.Context, bag *models.Bag) (bool, error) {
	exists := 1
	err := b.dbx(ctx).QueryRowContext(ctx, queries.SelectExistBookAlreadyInBag.ToString(), &bag.BookID, &bag.UserID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (b *bagRepository) FindBagByUser(ctx context.Context, userID uint64) (result []*models.Bag, err error) {
	rows, err := b.dbx(ctx).QueryContext(ctx, queries.SelectBagByUserQuery.ToString(), userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		bag := new(models.Bag)
		err = rows.Scan(&bag.ID, &bag.UserID, &bag.BookID,
			&bag.User.ID, &bag.User.Name, &bag.User.Address,
			&bag.User.Email, &bag.User.PhoneNumber, &bag.User.Password, &bag.User.DoB,
			&bag.User.ProfilePicture, &bag.User.CreatedAt, &bag.User.UpdatedAt, &bag.User.DeletedAt,
			&bag.User.VerifiedAt, &bag.Book.ID, &bag.Book.Title, &bag.Book.Description, &bag.Book.Image,
			&bag.Book.Status, &bag.Book.CreatedAt, &bag.Book.UpdatedAt, &bag.Book.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, bag)
	}
	return result, nil
}

func (b *bagRepository) Save(ctx context.Context, entity *models.Bag) (result *models.Bag, err error) {
	id, err := utils.Save[models.Bag](ctx, b.dbx(ctx), queries.InsertBagQuery, &entity.UserID, &entity.BookID)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bagRepository) DeleteByUserAndSelectedBooks(ctx context.Context, userId uint64, bookIds []uint64) error {
	queryIn := ` AND book_id IN (%s)`
	wildCards := []string{}
	for _, id := range bookIds {
		wildCards = append(wildCards, strconv.Itoa(int(id)))
	}
	wildCardStr := strings.Join(wildCards, ",")
	queryIn = fmt.Sprintf(queryIn, wildCardStr)
	selectStr := queries.DeleteUserBagQuery.ToString() + queryIn
	log.Printf(selectStr)
	_, err := b.dbx(ctx).ExecContext(ctx, selectStr, userId)
	if err != nil {
		return err
	}
	return nil
}

func (b *bagRepository) DeleteUserBag(ctx context.Context, userId uint64) error {
	_, err := b.dbx(ctx).ExecContext(ctx, queries.DeleteUserBagQuery.ToString(), userId)
	if err != nil {
		return err
	}
	return nil
}

var _ repositories_interfaces.BagRepository = &bagRepository{}

func newBagRepository(dbx repositories_interfaces.DB) *bagRepository {
	return &bagRepository{
		dbx: dbx,
	}
}
