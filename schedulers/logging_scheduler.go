package schedulers

import (
	"context"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
	"os"
)

type loggerScheduler struct {
	awsClient *aws_client.AWSClient
}

func (l *loggerScheduler) UploadLoggerScheduler(ctx context.Context) error {

	file, err := utils.WriteLogFile()
	if err != nil {
		return err
	}

	fileName := file.Name()
	_, err = l.awsClient.UploadFileToS3(fileName, "logs/")
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
