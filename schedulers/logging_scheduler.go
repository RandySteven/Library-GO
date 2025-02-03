package schedulers

import (
	"context"
	"github.com/RandySteven/Library-GO/enums"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
	"os"
)

type loggerScheduler struct {
	awsClient aws_client.AWS
}

func (l *loggerScheduler) UploadLoggerScheduler(ctx context.Context) error {

	file, err := utils.WriteLogFile()
	if err != nil {
		return err
	}

	fileName := file.Name()
	_, err = l.awsClient.UploadFileToS3(ctx, fileName, enums.LogsPath)
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

func newLoggingScheduler(awsClient aws_client.AWS) *loggerScheduler {
	return &loggerScheduler{
		awsClient: awsClient,
	}
}
