package aws_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/aws/aws-sdk-go-v2/config"
	credentials2 "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type (
	AWS interface {
		SessionClient() *session.Session
		BedrockClient() *bedrockruntime.Client
		S3() *s3.S3
		GeneratePromptResult(ctx context.Context, request any) (outputText string, err error)
		ListBucket() (result *s3.ListBucketsOutput, err error)
		UploadImageFile(ctx context.Context, fileRequest io.Reader, filePath string, fileHeader *multipart.FileHeader, width, height uint) (resultLocation *string, err error)
		CreateBucket(name string) error
		UploadFileToS3(ctx context.Context, fileName, path string) (string, error)
	}

	AWSClient struct {
		session *session.Session
		s3      *s3.S3
		brc     *bedrockruntime.Client
	}
)

var _ AWS = &AWSClient{}

func NewAWSClient(configYml *configs.Config) (*AWSClient, error) {
	if configYml == nil || &configYml.Config.AWS == nil {
		return nil, errors.New("AWS configuration is required")
	}

	awsCfg := configYml.Config.AWS
	if awsCfg.AccessKeyID == "" || awsCfg.SecretAccessKey == "" || awsCfg.Region == "" {
		return nil, errors.New("AWS access key ID, secret access key, and region are required")
	}

	bedrockCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsCfg.Region),
		config.WithCredentialsProvider(credentials2.NewStaticCredentialsProvider(
			awsCfg.AccessKeyID,
			awsCfg.SecretAccessKey,
			"",
		)),
		config.WithHTTPClient(&http.Client{
			Timeout: 120 * time.Second,
		}))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsCfg.Region),
		Credentials: credentials.NewStaticCredentials(
			awsCfg.AccessKeyID,
			awsCfg.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &AWSClient{
		session: sess,
		s3:      s3.New(sess),
		brc:     bedrockruntime.NewFromConfig(bedrockCfg),
	}, nil
}

func (b *AWSClient) BedrockClient() *bedrockruntime.Client {
	return b.brc
}

func (b *AWSClient) SessionClient() *session.Session {
	return b.session
}

func (b *AWSClient) S3() *s3.S3 {
	return b.s3
}
