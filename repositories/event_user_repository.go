package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type eventUserRepository struct {
	dbx repositories_interfaces.DB
}

func (e *eventUserRepository) Save(ctx context.Context, entity *models.EventUser) (*models.EventUser, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) FindByID(ctx context.Context, id uint64) (*models.EventUser, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.EventUser, error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.EventUserRepository = &eventUserRepository{}

func newEventUserRepository(dbx repositories_interfaces.DB) *eventUserRepository {
	return &eventUserRepository{
		dbx: dbx,
	}
}
