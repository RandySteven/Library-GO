package usecases_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookUsecaseTestSuite struct {
	suite.Suite
}

func TestBookUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BookUsecaseTestSuite))
}
