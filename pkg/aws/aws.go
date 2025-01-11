package aws_client

import (
	"context"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/aws/aws-sdk-go-v2/config"
	credentials2 "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) (resultLocation *string, err error)
		CreateBucket(name string) error
	}

	AWSClient struct {
		session *session.Session
		s3      *s3.S3
		brc     *bedrockruntime.Client
	}
)

func NewAWSClient(configYml *configs.Config) (*AWSClient, error) {
	awsCfg := configYml.Config.AWS
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
	if err != nil {
		return nil, err
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
