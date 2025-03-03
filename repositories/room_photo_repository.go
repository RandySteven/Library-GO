package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/utils"
)

type roomPhotoRepository struct {
	dbx repositories_interfaces.DB
}

func (r *roomPhotoRepository) Save(ctx context.Context, entity *models.RoomPhoto) (*models.RoomPhoto, error) {
	id, err := utils.Save[models.RoomPhoto](ctx, r.dbx(ctx), ``, &entity.Photo, &entity.RoomID)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomPhotoRepository) FindByID(ctx context.Context, id uint64) (*models.RoomPhoto, error) {
	result := &models.RoomPhoto{}
	err := utils.FindByID[models.RoomPhoto](ctx, r.dbx(ctx), ``, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roomPhotoRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.RoomPhoto, error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.RoomPhotoRepository = &roomPhotoRepository{}

func newRoomPhotoRepository(dbx repositories_interfaces.DB) *roomPhotoRepository {
	return &roomPhotoRepository{
		dbx: dbx,
	}
}
