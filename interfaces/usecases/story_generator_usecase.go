package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type StoryGeneratorUsecase interface {
	GenerateStory(ctx context.Context, request *requests.StoryGeneratorRequest) (result *responses.StoryGeneratorResponse, customErr *apperror.CustomError)
}
