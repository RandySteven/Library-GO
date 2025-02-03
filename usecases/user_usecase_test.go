package usecases_test

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/usecases"
	"testing"
	"time"

	"github.com/RandySteven/Library-GO/mocks"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	ctx      context.Context
	userRepo *mocks.UserRepository
	usecase  usecases_interfaces.UserUsecase
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.userRepo = new(mocks.UserRepository)
	suite.usecase = usecases.NewUsecases()
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (suite *UserUsecaseTestSuite) TestGetUserDetail() {
	suite.Run("success to get user detail", func() {
		// Arrange
		userID := uint64(1)
		expectedUser := &models.User{
			ID:             userID,
			Name:           "John Doe",
			Address:        "123 Main St",
			PhoneNumber:    "123-456-7890",
			ProfilePicture: "profile.jpg",
			DoB:            time.Now(),
		}
		expectedResponse := &responses.UserDetailResponse{
			ID:             userID,
			Name:           "John Doe",
			Address:        "123 Main St",
			PhoneNumber:    "123-456-7890",
			ProfilePicture: "profile.jpg",
			DoB:            time.Now(),
		}

		suite.userRepo.On("FindByID", suite.ctx, userID).Return(expectedUser, nil).Once()

		// Act
		result, customErr := suite.usecase.GetUserDetail(suite.ctx, userID)

		// Assert
		suite.NoError(customErr)
		suite.Equal(expectedResponse, result)
		suite.userRepo.AssertExpectations(suite.T())
	})

}
