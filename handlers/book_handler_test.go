package handlers_test

import (
	"bytes"
	"context"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/handlers"
	mocks "github.com/RandySteven/Library-GO/mocks/interfaces/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BookHandlerTestSuite struct {
	suite.Suite
	usecases *mocks.BookUsecase
	handlers handlers.Handlers
	ctx      context.Context
}

func (suite *BookHandlerTestSuite) SetupSuite() {
	suite.usecases = new(mocks.BookUsecase)
}

func TestBookHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}

func (suite *BookHandlerTestSuite) TestAddBook() {
	suite.Run("success create book", func() {
		request := ``
		requestBytStr := []byte(request)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBytStr))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		bookHandler := handlers.BookHandler{}
		bookHandler.AddBook(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
}

func (suite *BookHandlerTestSuite) TestGetAllBooks() {
	suite.Run("success to get all books", func() {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		bookHandler := handlers.BookHandler{}
		ctx := context.WithValue(req.Context(), enums.RequestID, uuid.NewString())

		suite.usecases.On("GetAllBooks", ctx).
			Return(mock.AnythingOfType("[]*responses.ListBooksResponse"), nil).
			Once()

		bookHandler.GetAllBooks(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
}
