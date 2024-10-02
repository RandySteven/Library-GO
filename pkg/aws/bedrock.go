package aws_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second) // 60-second timeout
	defer cancel()

	output, err := a.brc.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		Body:        payloadBytes,
		ModelId:     aws.String("amazon.titan-text-lite-v1"),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("application/json"),
	})
	if err != nil {
		return "", err
	}

	var response BedrockResponse
	if err := json.Unmarshal(output.Body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if len(response.Results) > 0 {
		outputText = response.Results[0].OutputText
	}
	log.Printf("Full response: %+v\n", response)
	fileName := utils.GenerateStoryName()

	err = utils.GenerateStoryFile(fileName, outputText)
	if err != nil {
		return "", err
	}
	bucketsCh := make(chan *s3.ListBucketsOutput)
	errCh := make(chan error)
	go func() {
		buckets, err := a.ListBucket()
		if err != nil {
			errCh <- err
			return
		}
		bucketsCh <- buckets
	}()

	uploader := s3manager.NewUploader(a.session)

	select {
	case buckets := <-bucketsCh:
		resultLocation, err := a.UploadFile(uploader, "./temp-stories/"+fileName, *buckets.Buckets[0].Name, "stories/"+fileName)
		if err != nil {
			return "", err
		}
		return *resultLocation, nil

	case err = <-errCh:
		// Handle error in bucket listing
		return "", err
	}
}
