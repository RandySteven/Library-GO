package aws_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"time"
)

type BedrockResponse struct {
	InputTextTokenCount int `json:"inputTextTokenCount"`
	Results             []struct {
		TokenCount       int    `json:"tokenCount"`
		OutputText       string `json:"outputText"`
		CompletionReason string `json:"completionReason"`
	} `json:"results"`
}

func (a *AWSClient) GeneratePromptResult(ctx context.Context, request any) (outputText string, err error) {
	// Marshalling the request into JSON
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	// Setting a longer timeout for the context, if needed
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second) // 60-second timeout
	defer cancel()

	// Making the request to the model
	output, err := a.brc.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		Body:        payloadBytes,
		ModelId:     aws.String("amazon.titan-text-lite-v1"),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("application/json"),
	})
	if err != nil {
		return "", err
	}

	// Optionally, you can unmarshal the result if it's JSON formatted
	var response BedrockResponse
	if err := json.Unmarshal(output.Body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if len(response.Results) > 0 {
		outputText = response.Results[0].OutputText
	}
	// For debugging or inspecting the full response
	log.Printf("Full response: %+v\n", response)

	return outputText, nil
}
