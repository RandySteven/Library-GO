package aws_client

import (
	"context"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSClient struct {
	s3  *s3.S3
	brc *bedrockruntime.Client
}

func NewAWSClient(configYml *configs.Config) (*AWSClient, error) {
	awsCfg := configYml.Config.AWS
	bedrockCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsCfg.Region))
	if err != nil {
		return nil, err
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsCfg.Region),
		Credentials: credentials.NewStaticCredentials(
			awsCfg.AccessKeyID,
			awsCfg.SecretAccessKey,
			"",
		),
	})
	return &AWSClient{
		s3:  s3.New(sess),
		brc: bedrockruntime.NewFromConfig(bedrockCfg),
	}, nil
}

func (b *AWSClient) BedrockClient() *bedrockruntime.Client {
	return b.brc
}
