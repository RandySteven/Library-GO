package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	jwt2 "github.com/RandySteven/Library-GO/pkg/jwt"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"sync"
	"time"
)

type onboardingUsecase struct {
	userRepo     repositories_interfaces.UserRepository
	roleUserRepo repositories_interfaces.RoleUserRepository
	pubSub       rabbitmqs_client.PubSub
	transaction  repositories_interfaces.Transaction
}

func (o *onboardingUsecase) RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (result *responses.UserRegisterResponse, customErr *apperror.CustomError) {
	var (
		user        = &models.User{}
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
		err         error
	)

	customErr = o.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {
		go func() {
			defer wg.Done()
			_, err = o.userRepo.FindByEmail(ctx, request.Email)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to find user by email "`+request.Email+`"`, err)
					return
				}
				return
			}
		}()

		go func() {
			defer wg.Done()
			_, err = o.userRepo.FindByPhoneNumber(ctx, request.PhoneNumber)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to find user by phone number "`+request.PhoneNumber+`"`, err)
					return
				}
				return
			}
		}()

		go func() {
			wg.Wait()
			close(customErrCh)
		}()

		dob, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", request.Year, request.Month, request.Day))
		user = &models.User{
			Name:        fmt.Sprintf("%s %s", request.FirstName, request.LastName),
			Email:       request.Email,
			Password:    utils.HashPassword(request.Password),
			PhoneNumber: request.PhoneNumber,
			Address:     request.Address,
			DoB:         dob,
		}
		user, err = o.userRepo.Save(ctx, user)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user`, err)
		}

		select {
		case customErr = <-customErrCh:
			return customErr
		}
	})
	if customErr != nil {
		return nil, customErr
	}
	return &responses.UserRegisterResponse{
		ID:        utils.HashID(user.ID),
		Token:     uuid.NewString(),
		CreatedAt: time.Now(),
	}, nil
}

func (o *onboardingUsecase) LoginUser(ctx context.Context, request *requests.UserLoginRequest) (result *responses.UserLoginResponse, customErr *apperror.CustomError) {
	user, err := o.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NewCustomError(apperror.ErrNotFound, `failed to login consumers not found`, err)
		}
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to connect db`, err)
	}
	isPassExists := utils.ComparePassword(request.Password, user.Password)
	if !isPassExists {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `invalid credentials`, err)
	}

	roleUser, err := o.roleUserRepo.FindRoleUserByUserID(ctx, user.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get role user`, err)
	}

	claims := &jwt2.JWTClaim{
		UserID: user.ID,
		RoleID: roleUser.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Applications",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(jwt2.JwtKey)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to generate token`, err)
	}
	result = &responses.UserLoginResponse{
		Token: token,
	}
	return result, nil
}

func (o *onboardingUsecase) GetLoginUser(ctx context.Context) (result *responses.LoginUserResponse, customErr *apperror.CustomError) {
	id := ctx.Value(enums.UserID).(uint64)
	user, err := o.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user login`, err)
	}

	roleUser, err := o.roleUserRepo.FindRoleUserByUserID(ctx, user.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get role user`, err)
	}

	result = &responses.LoginUserResponse{
		ID:             user.ID,
		RoleID:         roleUser.RoleID,
		Name:           user.Name,
		Email:          user.Email,
		Password:       user.Password,
		PhoneNumber:    user.PhoneNumber,
		Address:        user.Address,
		ProfilePicture: user.ProfilePicture,
		DoB:            user.DoB,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
	return result, nil
}

func (o *onboardingUsecase) VerifyToken(ctx context.Context, token string) (customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.OnboardingUsecase = &onboardingUsecase{}

func newOnboardingUsecase(
	userRepo repositories_interfaces.UserRepository,
	roleUserRepo repositories_interfaces.RoleUserRepository,
	pubSub rabbitmqs_client.PubSub,
	transaction repositories_interfaces.Transaction) *onboardingUsecase {
	return &onboardingUsecase{
		userRepo:     userRepo,
		roleUserRepo: roleUserRepo,
		pubSub:       pubSub,
		transaction:  transaction,
	}
}
