package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type chatRepository struct {
	dbx repositories_interfaces.DB
}

func (c *chatRepository) Save(ctx context.Context, entity *models.Chat) (*models.Chat, error) {
	id, err := utils.Save[models.Chat](ctx, c.dbx(ctx), queries.InsertIntoChatQuery,
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

func (c *chatRepository) FindChatByRoomID(ctx context.Context, roomChatID uint64) (result []*models.Chat, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.ChatRepository = &chatRepository{}

func newChatRepository(dbx repositories_interfaces.DB) *chatRepository {
	return &chatRepository{
		dbx: dbx,
	}
}
