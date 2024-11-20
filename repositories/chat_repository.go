package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type chatRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (c *chatRepository) Save(ctx context.Context, entity *models.Chat) (*models.Chat, error) {
	id, err := utils.Save[models.Chat](ctx, c.Trigger(), queries.InsertIntoChatQuery,
		&entity.RoomChatID, &entity.UserID, &entity.Chat)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (c *chatRepository) FindByID(ctx context.Context, id uint64) (*models.Chat, error) {
	return nil, nil
}

func (c *chatRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(c.db, c.tx)
}

func (c *chatRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) SetTx(tx *sql.Tx) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) GetTx(ctx context.Context) *sql.Tx {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) FindChatByRoomID(ctx context.Context, roomChatID uint64) (result []*models.Chat, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.ChatRepository = &chatRepository{}

func newChatRepository(db *sql.DB) *chatRepository {
	return &chatRepository{
		db: db,
	}
}
