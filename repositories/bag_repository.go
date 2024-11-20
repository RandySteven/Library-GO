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

type bagRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bagRepository) CheckBagExists(ctx context.Context, bag *models.Bag) (bool, error) {
	exists := 1
	err := b.Trigger().QueryRowContext(ctx, queries.SelectExistBookAlreadyInBag.ToString(), &bag.BookID, &bag.UserID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (b *bagRepository) FindBagByUser(ctx context.Context, userID uint64) (result []*models.Bag, err error) {
	rows, err := b.Trigger().QueryContext(ctx, queries.SelectBagByUserQuery.ToString(), userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		bag := new(models.Bag)
		err = rows.Scan(&bag.ID, &bag.UserID, &bag.BookID)
		if err != nil {
			return nil, err
		}
		result = append(result, bag)
	}
	return result, nil
}

func (b *bagRepository) Save(ctx context.Context, entity *models.Bag) (result *models.Bag, err error) {
	id, err := utils.Save[models.Bag](ctx, b.Trigger(), queries.InsertBagQuery, &entity.UserID, &entity.BookID)
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
	_, err := b.Trigger().ExecContext(ctx, selectStr, userId)
	if err != nil {
		return err
	}
	return nil
}

func (b *bagRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(b.db, b.tx)
}

func (b *bagRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.tx = tx
	return nil
}

func (b *bagRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *bagRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *bagRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *bagRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *bagRepository) DeleteUserBag(ctx context.Context, userId uint64) error {
	_, err := b.Trigger().ExecContext(ctx, queries.DeleteUserBagQuery.ToString(), userId)
	if err != nil {
		return err
	}
	return nil
}

var _ repositories_interfaces.BagRepository = &bagRepository{}

func newBagRepository(db *sql.DB) *bagRepository {
	return &bagRepository{
		db: db,
	}
}
