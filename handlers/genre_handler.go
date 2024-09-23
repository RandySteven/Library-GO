package handlers

import (
	"context"
	"github.com/RandySteven/Library-GO/enums"
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"net/http"
)

type GenreHandler struct {
	usecase usecases_interfaces.GenreUsecase
}

func (g *GenreHandler) AddNewGenre(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (g *GenreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `genres`
	)
	result, customErr := g.usecase.GetAllGenres(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get genres`, &dataKey, result, nil)
}

var _ handlers_interfaces.GenreHandler = &GenreHandler{}

func newGenreHandler(usecase usecases_interfaces.GenreUsecase) *GenreHandler {
	return &GenreHandler{
		usecase: usecase,
	}
}
