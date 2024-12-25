package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roomChatRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *roomChatRepository) Save(ctx context.Context, entity *models.RoomChat) (*models.RoomChat, error) {
	id, err := utils.Save[models.RoomChat](ctx, r.Trigger(), queries.InsertRoomChatQuery, entity.RoomName)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomChatRepository) FindByID(ctx context.Context, id uint64) (*models.RoomChat, error) {
	result := &models.RoomChat{}
	err := utils.FindByID[models.RoomChat](ctx, r.Trigger(), queries.SelectRoomChatByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roomChatRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.RoomChat, error) {
	return utils.FindAll[models.RoomChat](ctx, r.Trigger(), queries.SelectRoomChatQuery)
}

func (r *roomChatRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(r.db, r.tx)
}

func (r *roomChatRepository) BeginTx(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *roomChatRepository) CommitTx(ctx context.Context) error {
	return r.tx.Commit()
}

func (r *roomChatRepository) RollbackTx(ctx context.Context) error {
	return r.tx.Rollback()
}

func (r *roomChatRepository) SetTx(tx *sql.Tx) {
	r.tx = tx
}

func (r *roomChatRepository) GetTx(ctx context.Context) *sql.Tx {
	return r.tx
}

var _ repositories_interfaces.RoomChatRepository = &roomChatRepository{}

func newRoomChatRepository(db *sql.DB) *roomChatRepository {
	return &roomChatRepository{
		db: db,
	}
}
