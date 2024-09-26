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

func (a *AWSClient) GeneratePromptResult(ctx context.Context, request any) (result string, err error) {
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

	// Since output.Body is already []byte, directly convert to string
	result = string(output.Body)
	log.Println("Generated text:", result)

	// Optionally, you can unmarshal the result if it's JSON formatted
	var response map[string]interface{}
	if err := json.Unmarshal(output.Body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// For debugging or inspecting the full response
	log.Printf("Full response: %+v\n", response)

	return result, nil
}
