package handlers

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/enums"
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"net/http"
)

type StoryGeneratorHandler struct {
	usecase usecases_interfaces.StoryGeneratorUsecase
}

func (s *StoryGeneratorHandler) GenerateStory(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.StoryGeneratorRequest{}
		dataKey = `story`
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := s.usecase.GenerateStory(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to generate story`, &dataKey, result, nil)
}

var _ handlers_interfaces.StoryGeneratorHandler = &StoryGeneratorHandler{}

func newStoryGeneratorHandler(usecase usecases_interfaces.StoryGeneratorUsecase) *StoryGeneratorHandler {
	return &StoryGeneratorHandler{
		usecase: usecase,
	}
}
