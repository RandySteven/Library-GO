package usecases

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BagUsecaseTestSuite struct {
	suite.Suite
}

func (b *BookUsecaseTestSuite) SetupSuite() {

}

func TestBagUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BagUsecaseTestSuite))
}
