package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type eventRepository struct {
	db *sql.DB
	tx *sql.Tx
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

func (e *eventRepository) InitTrigger() repositories_interfaces.Trigger {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) SetTx(tx *sql.Tx) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) GetTx(ctx context.Context) *sql.Tx {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.EventRepository = &eventRepository{}

func newEventRepository(db *sql.DB) *eventRepository {
	return &eventRepository{
		db: db,
	}
}
