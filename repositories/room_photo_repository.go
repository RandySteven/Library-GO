package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/utils"
)

type roomPhotoRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *roomPhotoRepository) Save(ctx context.Context, entity *models.RoomPhoto) (*models.RoomPhoto, error) {
	id, err := utils.Save[models.RoomPhoto](ctx, r.Trigger(), ``, &entity.Photo, &entity.RoomID)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomPhotoRepository) FindByID(ctx context.Context, id uint64) (*models.RoomPhoto, error) {
	result := &models.RoomPhoto{}
	err := utils.FindByID[models.RoomPhoto](ctx, r.Trigger(), ``, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roomPhotoRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.RoomPhoto, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roomPhotoRepository) Trigger() repositories_interfaces.Trigger {
	//TODO implement me
	panic("implement me")
}

func (r *roomPhotoRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *roomPhotoRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *roomPhotoRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *roomPhotoRepository) SetTx(tx *sql.Tx) {
	//TODO implement me
	panic("implement me")
}

func (r *roomPhotoRepository) GetTx(ctx context.Context) *sql.Tx {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.RoomPhotoRepository = &roomPhotoRepository{}

func newRoomPhotoRepository(db *sql.DB) *roomPhotoRepository {
	return &roomPhotoRepository{
		db: db,
	}
}
