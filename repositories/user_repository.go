package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type userRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (u *userRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = u.db
	if u.tx != nil {
		trigger = u.tx
	}
	return trigger
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return u.findUser(ctx, `email`, email)
}

func (u *userRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	return u.findUser(ctx, `phone`, phoneNumber)
}

func (u *userRepository) Save(ctx context.Context, entity *models.User) (result *models.User, err error) {
	result = &models.User{}
	id, err := utils.Save[models.User](ctx, u.InitTrigger(), queries.InsertUserQuery,
		&entity.Name, &entity.Address, &entity.Email, &entity.PhoneNumber, &entity.Password, &entity.DoB)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (u *userRepository) FindByID(ctx context.Context, id uint64) (result *models.User, err error) {
	err = utils.FindByID[models.User](ctx, u.InitTrigger(), queries.SelectUserByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(ctx context.Context, entity *models.User) (result *models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) BeginTx(ctx context.Context) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.tx = tx
	return nil
}

func (u *userRepository) CommitTx(ctx context.Context) error {
	return u.tx.Commit()
}

func (u *userRepository) RollbackTx(ctx context.Context) error {
	return u.tx.Rollback()
}

func (u *userRepository) SetTx(tx *sql.Tx) {
	u.tx = tx
}

func (u *userRepository) GetTx(ctx context.Context) *sql.Tx {
	return u.tx
}

var _ repositories_interfaces.UserRepository = &userRepository{}

func newUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) findUser(ctx context.Context, by string, identifier string) (*models.User, error) {
	var query string
	if by == "phone" {
		query = queries.SelectUserByPhoneNumberQuery.ToString()
	} else {
		query = queries.SelectUserByEmailQuery.ToString()
	}
	result := &models.User{}
	err := u.InitTrigger().QueryRowContext(ctx, query, identifier).Scan(
		&result.ID, &result.Name, &result.Address, &result.Email,
		&result.PhoneNumber, &result.Password, &result.DoB,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
