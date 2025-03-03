package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roomChatRepository struct {
	dbx repositories_interfaces.DB
}

func (r *roomChatRepository) Save(ctx context.Context, entity *models.RoomChat) (*models.RoomChat, error) {
	id, err := utils.Save[models.RoomChat](ctx, r.dbx(ctx), queries.InsertRoomChatQuery, entity.RoomName)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomChatRepository) FindByID(ctx context.Context, id uint64) (*models.RoomChat, error) {
	result := &models.RoomChat{}
	err := utils.FindByID[models.RoomChat](ctx, r.dbx(ctx), queries.SelectRoomChatByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roomChatRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.RoomChat, error) {
	return utils.FindAll[models.RoomChat](ctx, r.dbx(ctx), queries.SelectRoomChatQuery)
}

var _ repositories_interfaces.RoomChatRepository = &roomChatRepository{}

func newRoomChatRepository(dbx repositories_interfaces.DB) *roomChatRepository {
	return &roomChatRepository{
		dbx: dbx,
	}
}
