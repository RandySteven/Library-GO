package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type roleUserRepository struct {
	dbx repositories_interfaces.DB
}

func (r *roleUserRepository) Save(ctx context.Context, entity *models.RoleUser) (*models.RoleUser, error) {
	id, err := utils.Save[models.RoleUser](ctx, r.dbx(ctx), queries.InsertRoleUserQuery, entity.RoleID, entity.UserID)
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

func (r *roleUserRepository) FindRoleUserByUserID(ctx context.Context, id uint64) (result *models.RoleUser, err error) {
	result = &models.RoleUser{}
	err = r.dbx(ctx).QueryRowContext(ctx, queries.SelectRoleUserByUserIDQuery.ToString(), id).
		Scan(&result.ID, &result.RoleID, &result.UserID, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ repositories_interfaces.RoleUserRepository = &roleUserRepository{}

func newRoleUserRepository(dbx repositories_interfaces.DB) *roleUserRepository {
	return &roleUserRepository{
		dbx: dbx,
	}
}
