package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
)

type storyGeneratorUsecase struct {
	awsClient          *aws_client.AWSClient
	storyGeneratorRepo repositories_interfaces.StoryGeneratorRepository
}

func (s *storyGeneratorUsecase) GenerateStory(ctx context.Context, request *requests.StoryGeneratorRequest) (result *responses.StoryGeneratorResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.StoryGeneratorUsecase = &storyGeneratorUsecase{}

func newStoryGeneratorUsecase(awsClient *aws_client.AWSClient) *storyGeneratorUsecase {
	return &storyGeneratorUsecase{
		awsClient: awsClient,
	}
}
