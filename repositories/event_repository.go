package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type eventRepository struct {
	dbx repositories_interfaces.DB
}

func (e *eventRepository) Save(ctx context.Context, entity *models.Event) (result *models.Event, err error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) FindByID(ctx context.Context, id uint64) (result *models.Event, err error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) Update(ctx context.Context, entity *models.Event) (result *models.Event, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.EventRepository = &eventRepository{}

func newEventRepository(dbx repositories_interfaces.DB) *eventRepository {
	return &eventRepository{
		dbx: dbx,
	}
}
