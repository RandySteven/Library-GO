package handlers_test

import (
	"github.com/RandySteven/Library-GO/handlers"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DevHandlerTestSuite struct {
	suite.Suite
}

func (suite *DevHandlerTestSuite) SetupSuite() {

}

func TestDevHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DevHandlerTestSuite))
}

func (suite *DevHandlerTestSuite) TestHealthCheck() {
	suite.Run("success hit health check", func() {
		req := httptest.NewRequest(http.MethodGet, "/dev", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		devHandler := &handlers.DevHandler{}
		devHandler.HealthCheck(w, req)
		suite.Equal(w.Code, http.StatusOK)
	})
}
