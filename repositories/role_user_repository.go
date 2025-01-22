package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roleUserRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *roleUserRepository) Save(ctx context.Context, entity *models.RoleUser) (*models.RoleUser, error) {
	id, err := utils.Save[models.RoleUser](ctx, r.Trigger(), queries.InsertRoleUserQuery, entity.RoleID, entity.UserID)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (r *roleUserRepository) FindByID(ctx context.Context, id uint64) (*models.RoleUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleUserRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.RoleUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleUserRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(r.db, r.tx)
}

func (r *roleUserRepository) BeginTx(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *roleUserRepository) CommitTx(ctx context.Context) error {
	return utils.CommitTx(ctx, r.tx)
}

func (r *roleUserRepository) RollbackTx(ctx context.Context) error {
	return utils.RollbackTx(ctx, r.tx)
}

func (r *roleUserRepository) SetTx(tx *sql.Tx) {
	r.tx = tx
}

func (r *roleUserRepository) GetTx(ctx context.Context) *sql.Tx {
	return r.tx
}

func (r *roleUserRepository) FindRoleUserByUserID(ctx context.Context, id uint64) (result *models.RoleUser, err error) {
	result = &models.RoleUser{}
	err = r.Trigger().QueryRowContext(ctx, queries.SelectRoleUserByUserIDQuery.ToString(), id).
		Scan(&result.ID, &result.RoleID, &result.UserID, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ repositories_interfaces.RoleUserRepository = &roleUserRepository{}

func newRoleUserRepository(db *sql.DB) *roleUserRepository {
	return &roleUserRepository{
		db: db,
	}
}
