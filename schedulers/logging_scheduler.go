package schedulers

import (
	"context"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type loggerScheduler struct {
	awsClient *aws_client.AWSClient
}

func (l *loggerScheduler) UploadLoggerScheduler(ctx context.Context) error {
	buckets, err := l.awsClient.ListBucket()
	if err != nil {
		return err
	}

	file, err := utils.WriteLogFile()
	if err != nil {
		return err
	}
	fileName := file.Name()
	_, err = l.awsClient.UploadFile(s3manager.NewUploader(l.awsClient.SessionClient()), fileName, *buckets.Buckets[0].Name, "logs/"+fileName)
	if err != nil {
		return err
	}

	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}

var _ schedulers_interfaces.LoggerScheduler = &loggerScheduler{}

func newLoggingScheduler(awsClient *aws_client.AWSClient) *loggerScheduler {
	return &loggerScheduler{
		awsClient: awsClient,
	}
}
