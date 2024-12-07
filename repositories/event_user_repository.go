package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type eventUserRepository struct {
	db *sql.DB
	tx *sql.Tx
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

func (e *eventUserRepository) Trigger() repositories_interfaces.Trigger {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) SetTx(tx *sql.Tx) {
	//TODO implement me
	panic("implement me")
}

func (e *eventUserRepository) GetTx(ctx context.Context) *sql.Tx {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.EventUserRepository = &eventUserRepository{}

func newEventUserRepository(db *sql.DB) *eventUserRepository {
	return &eventUserRepository{
		db: db,
	}
}
