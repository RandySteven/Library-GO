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
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	jwt2 "github.com/RandySteven/Library-GO/pkg/jwt"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type onboardingUsecase struct {
	userRepo repositories_interfaces.UserRepository
}

func (o *onboardingUsecase) refreshTx(ctx context.Context) {
	tx := o.userRepo.GetTx(ctx)
	o.userRepo.SetTx(tx)
}

func (o *onboardingUsecase) RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (result *responses.UserRegisterResponse, customErr *apperror.CustomError) {
	var (
		user        = &models.User{}
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
	)
	err := o.userRepo.BeginTx(ctx)
	if err != nil {
		return
	}
	defer func() {
		defer o.userRepo.SetTx(nil)
		if r := recover(); r != nil {
			_ = o.userRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = o.userRepo.RollbackTx(ctx)
			return
		} else {
			if err = o.userRepo.CommitTx(ctx); err != nil {
				log.Println("failed to commit transaction")
				return
			}
			return
		}
	}()

	go func() {
		defer wg.Done()
		_, err = o.userRepo.FindByEmail(ctx, request.Email)
		if err != nil {
			if !errors.As(err, &sql.ErrNoRows) {
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
			if !errors.As(err, &sql.ErrNoRows) {
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
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user`, err)
	}

	select {
	case customErr = <-customErrCh:
		return nil, customErr
	default:
		return &responses.UserRegisterResponse{
			ID:        utils.HashID(user.ID),
			Token:     uuid.NewString(),
			CreatedAt: time.Now(),
		}, nil
	}

	return
}

func (o *onboardingUsecase) LoginUser(ctx context.Context, request *requests.UserLoginRequest) (result *responses.UserLoginResponse, customErr *apperror.CustomError) {
	user, err := o.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, apperror.NewCustomError(apperror.ErrNotFound, `failed to login email not found`, err)
		}
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to connect db`, err)
	}
	isPassExists := utils.ComparePassword(request.Password, user.Password)
	if !isPassExists {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `invalid credentials`, err)
	}
	claims := &jwt2.JWTClaim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Applications",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(jwt2.JWT_KEY)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to generate token`, err)
	}
	result = &responses.UserLoginResponse{
		Token: token,
	}
	return result, nil
}

func (o *onboardingUsecase) VerifyToken(ctx context.Context, token string) (customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.OnboardingUsecase = &onboardingUsecase{}

func newOnboardingUsecase(userRepo repositories_interfaces.UserRepository) *onboardingUsecase {
	return &onboardingUsecase{
		userRepo: userRepo,
	}
}
