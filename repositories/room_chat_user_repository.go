package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roomChatUserRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *roomChatUserRepository) Save(ctx context.Context, entity *models.RoomChatUser) (*models.RoomChatUser, error) {
	id, err := utils.Save[models.RoomChatUser](ctx, r.Trigger(), queries.InsertIntoRoomChatUsersQuery, entity.RoomChatID, entity.UserID)
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

func (r *roomChatUserRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(r.db, r.tx)
}

func (r *roomChatUserRepository) BeginTx(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *roomChatUserRepository) CommitTx(ctx context.Context) error {
	return r.tx.Commit()
}

func (r *roomChatUserRepository) RollbackTx(ctx context.Context) error {
	return r.tx.Rollback()
}

func (r *roomChatUserRepository) SetTx(tx *sql.Tx) {
	r.tx = tx
}

func (r *roomChatUserRepository) GetTx(ctx context.Context) *sql.Tx {
	return r.tx
}

func (r *roomChatUserRepository) FindUserRooms(ctx context.Context, userId uint64) (result []*models.RoomChatUser, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.RoomChatUserRepository = &roomChatUserRepository{}

func newRoomChatUserRepository(db *sql.DB) *roomChatUserRepository {
	return &roomChatUserRepository{
		db: db,
	}
}
