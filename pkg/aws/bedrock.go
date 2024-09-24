package aws_client

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/aws"
)

func (a *AWSClient) GeneratePromptResult(ctx context.Context, request any) (output *bedrockruntime.InvokeModelWithResponseStreamOutput, err error) {
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	output, err = a.brc.InvokeModelWithResponseStream(ctx, &bedrockruntime.InvokeModelWithResponseStreamInput{
		Body:        payloadBytes,
		ModelId:     aws.String(""),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}
