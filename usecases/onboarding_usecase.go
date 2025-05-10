package usecases

import (
	"context"
	"database/sql"
	"encoding/json"
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
	oauth2_client "github.com/RandySteven/Library-GO/pkg/oauth2"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"
	"time"
)

type onboardingUsecase struct {
	userRepo     repositories_interfaces.UserRepository
	roleUserRepo repositories_interfaces.RoleUserRepository
	pubSub       rabbitmqs_client.PubSub
	transaction  repositories_interfaces.Transaction
	oauth2       oauth2_client.Oauth2
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

func (o *onboardingUsecase) GoogleLogin(ctx context.Context) (customErr *apperror.CustomError) {
	result := o.oauth2.LoginAuth(ctx)
	log.Println(result)
	return
}

func (o *onboardingUsecase) GoogleCallback(ctx context.Context) (result *responses.UserLoginResponse, customErr *apperror.CustomError) {
	user := &models.User{}
	roleUser := &models.RoleUser{}
	accesssToken, err := o.oauth2.CallbackAuth(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to callback oauth2`, err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accesssToken)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to hit oauth2 user info`, err)
	}

	oauth2Response := &responses.Oauth2Response{}
	err = json.NewDecoder(resp.Body).Decode(oauth2Response)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to read oauth2 user info`, err)
	}

	customErr = o.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {
		user, _ = o.userRepo.FindByEmail(ctx, oauth2Response.Email)
		if user != nil {
			return nil
		}

		user = &models.User{
			Name:           oauth2Response.Name,
			Email:          oauth2Response.Email,
			Password:       "",
			PhoneNumber:    "",
			ProfilePicture: oauth2Response.Picture,
		}

		user, err = o.userRepo.Save(ctx, user)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user`, err)
		}

		roleUser, err = o.roleUserRepo.Save(ctx, &models.RoleUser{
			UserID: user.ID,
			RoleID: uint64(enums.Member),
		})
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to register user role`, err)
		}
		return nil
	})
	if customErr != nil {
		return nil, customErr
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

var _ usecases_interfaces.OnboardingUsecase = &onboardingUsecase{}

func newOnboardingUsecase(
	userRepo repositories_interfaces.UserRepository,
	roleUserRepo repositories_interfaces.RoleUserRepository,
	pubSub rabbitmqs_client.PubSub,
	transaction repositories_interfaces.Transaction,
	oauth2 oauth2_client.Oauth2) *onboardingUsecase {
	return &onboardingUsecase{
		userRepo:     userRepo,
		roleUserRepo: roleUserRepo,
		pubSub:       pubSub,
		transaction:  transaction,
		oauth2:       oauth2,
	}
}
