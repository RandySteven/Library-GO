package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roomRepository struct {
	dbx repositories_interfaces.DB
}

func (r *roomRepository) Save(ctx context.Context, entity *models.Room) (*models.Room, error) {
	id, err := utils.Save[models.Room](ctx, r.dbx(ctx), queries.InsertRoomQuery, &entity.Name, &entity.Thumbnail)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomRepository) FindByID(ctx context.Context, id uint64) (*models.Room, error) {
	result := &models.Room{}
	err := utils.FindByID[models.Room](ctx, r.dbx(ctx), ``, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roomRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Room, error) {
	return utils.FindAll[models.Room](ctx, r.dbx(ctx), ``)
}

var _ repositories_interfaces.RoomRepository = &roomRepository{}

func newRoomRepository(dbx repositories_interfaces.DB) *roomRepository {
	return &roomRepository{
		dbx: dbx,
	}
}
