package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roomRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *roomRepository) Save(ctx context.Context, entity *models.Room) (*models.Room, error) {
	id, err := utils.Save[models.Room](ctx, r.Trigger(), queries.InsertRoomQuery, &entity.Name, &entity.Thumbnail)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roomRepository) FindByID(ctx context.Context, id uint64) (*models.Room, error) {
	result := &models.Room{}
	err := utils.FindByID[models.Room](ctx, r.Trigger(), ``, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roomRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Room, error) {
	return utils.FindAll[models.Room](ctx, r.Trigger(), ``)
}

func (r *roomRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(r.db, r.tx)
}

func (r *roomRepository) BeginTx(ctx context.Context) error {
	if r.tx == nil {
		tx, err := r.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		r.tx = tx
	}
	return nil
}

func (r *roomRepository) CommitTx(ctx context.Context) error {
	return r.tx.Commit()
}

func (r *roomRepository) RollbackTx(ctx context.Context) error {
	return r.tx.Rollback()
}

func (r *roomRepository) SetTx(tx *sql.Tx) {
	r.tx = tx
}

func (r *roomRepository) GetTx(ctx context.Context) *sql.Tx {
	return r.tx
}

var _ repositories_interfaces.RoomRepository = &roomRepository{}

func newRoomRepository(db *sql.DB) *roomRepository {
	return &roomRepository{
		db: db,
	}
}
