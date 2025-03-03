package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roomChatUserRepository struct {
	dbx repositories_interfaces.DB
}

func (r *roomChatUserRepository) Save(ctx context.Context, entity *models.RoomChatUser) (*models.RoomChatUser, error) {
	id, err := utils.Save[models.RoomChatUser](ctx, r.dbx(ctx), queries.InsertIntoRoomChatUsersQuery, entity.RoomChatID, entity.UserID)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomChatUserRepository) FindByID(ctx context.Context, id uint64) (*models.RoomChatUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roomChatUserRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.RoomChatUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roomChatUserRepository) FindUserRooms(ctx context.Context, userId uint64) (result []*models.RoomChatUser, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.RoomChatUserRepository = &roomChatUserRepository{}

func newRoomChatUserRepository(dbx repositories_interfaces.DB) *roomChatUserRepository {
	return &roomChatUserRepository{
		dbx: dbx,
	}
}
