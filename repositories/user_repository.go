package repositories

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"strconv"
	"strings"
)

type userRepository struct {
	dbx repositories_interfaces.DB
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return u.findUser(ctx, enums.OnboardByEmail, email)
}

func (u *userRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	return u.findUser(ctx, enums.OnboardByPhone, phoneNumber)
}

func (u *userRepository) Save(ctx context.Context, entity *models.User) (result *models.User, err error) {
	result = &models.User{}
	id, err := utils.Save[models.User](ctx, u.dbx(ctx), queries.InsertUserQuery,
		&entity.Name, &entity.Address, &entity.Email, &entity.PhoneNumber, &entity.Password, &entity.DoB)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (u *userRepository) FindByID(ctx context.Context, id uint64) (result *models.User, err error) {
	result = &models.User{}
	err = utils.FindByID[models.User](ctx, u.dbx(ctx), queries.SelectUserByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindSelectedUsersByID(ctx context.Context, ids []uint64) (result []*models.User, err error) {
	queryIn := ` WHERE id IN (%s)`
	wildCards := []string{}
	for _, id := range ids {
		wildCards = append(wildCards, strconv.Itoa(int(id)))
	}
	wildCardStr := strings.Join(wildCards, ",")
	queryIn = fmt.Sprintf(queryIn, wildCardStr)
	selectStr := queries.SelectUsersQuery.ToString() + queryIn
	rows, err := u.dbx(ctx).QueryContext(ctx, selectStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := new(models.User)
		err = rows.Scan(&user.ID, &user.Name, &user.Address, &user.Email, &user.PhoneNumber, &user.Password, &user.DoB,
			&user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

var _ repositories_interfaces.UserRepository = &userRepository{}

func newUserRepository(dbx repositories_interfaces.DB) *userRepository {
	return &userRepository{
		dbx: dbx,
	}
}

func (u *userRepository) findUser(ctx context.Context, by enums.OnboardMethod, identifier string) (*models.User, error) {
	var query string
	switch by {
	case enums.OnboardByPhone:
		query = queries.SelectUserByPhoneNumberQuery.ToString()
	case enums.OnboardByEmail:
		query = queries.SelectUserByEmailQuery.ToString()
	}
	result := &models.User{}
	err := u.dbx(ctx).QueryRowContext(ctx, query, identifier).Scan(
		&result.ID, &result.Name, &result.Address, &result.Email,
		&result.PhoneNumber, &result.Password, &result.DoB, &result.ProfilePicture, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt, &result.VerifiedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
