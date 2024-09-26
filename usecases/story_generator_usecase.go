package usecases

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
)

const (
	templatePromptFile = "./files/txt/prompt.template.txt"
)

type (
	storyGeneratorUsecase struct {
		awsClient          *aws_client.AWSClient
		storyGeneratorRepo repositories_interfaces.StoryGeneratorRepository
	}

	StoryPromptRequest struct {
		InputText            string     `json:"inputText"`
		TextGenerationConfig TextConfig `json:"textGenerationConfig"`
	}

	TextConfig struct {
		MaxTokenCount int      `json:"maxTokenCount"`
		StopSequences []string `json:"stopSequences,omitempty"`
		Temperature   float64  `json:"temperature"`
		TopP          float64  `json:"topP"`
	}
)

func (s *storyGeneratorUsecase) GenerateStory(ctx context.Context, request *requests.StoryGeneratorRequest) (result *responses.StoryGeneratorResponse, customErr *apperror.CustomError) {
	promptTemplate, err := utils.ReadFileContent(templatePromptFile)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to read template file", err)
	}

	// Create the prompt request
	storyPromptRequest := &StoryPromptRequest{
		InputText: fmt.Sprintf(promptTemplate,
			request.Title, request.Setting, request.Genre,
			request.Style, request.MainCharacter, request.SideCharacter, request.Ending),
		TextGenerationConfig: TextConfig{
			MaxTokenCount: 280,
			StopSequences: []string{},
			Temperature:   0.7,
			TopP:          1.0,
		},
	}

	// Generate the prompt result
	promptRes, err := s.awsClient.GeneratePromptResult(ctx, storyPromptRequest)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to generate prompt result", err)
	}

	return &responses.StoryGeneratorResponse{
		Result: promptRes,
	}, nil
}

var _ usecases_interfaces.StoryGeneratorUsecase = &storyGeneratorUsecase{}

func newStoryGeneratorUsecase(awsClient *aws_client.AWSClient) *storyGeneratorUsecase {
	return &storyGeneratorUsecase{
		awsClient: awsClient,
	}
}
