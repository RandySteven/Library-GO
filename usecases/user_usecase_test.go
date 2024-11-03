package usecases

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserUsecaseTestSuite struct {
	suite.Suite
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
