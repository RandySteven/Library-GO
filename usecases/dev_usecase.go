package usecases

import (
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
)

import (
	"context"
)

type devUsecase struct {
	awsCl  *aws_client.AWSClient
	pubsub rabbitmqs_client.PubSub
}

func (d *devUsecase) CreateBucket(ctx context.Context, name string) error {
	return d.awsCl.CreateBucket(name)
}

func (d *devUsecase) GetListBuckets(ctx context.Context) ([]string, error) {
	result, err := d.awsCl.ListBucket()
	if err != nil {
		return nil, err
	}
	buckets := result.Buckets
	response := []string{}
	for _, bucket := range buckets {
		response = append(response, *bucket.Name)
	}
	return response, nil
}

func (d *devUsecase) MessageBrokerCheckerHealth(ctx context.Context) (string, error) {
	requestID := (ctx.Value(enums.RequestID)).(string)
	err := d.pubsub.Send("dev_checker", "dev-send-message", fmt.Sprintf("Check dev healthy with request ID : %s", requestID))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("success health message broker : %s", requestID), nil
}

var _ usecases_interfaces.DevUsecase = &devUsecase{}

func newDevUsecase(awsCl *aws_client.AWSClient, pubsub rabbitmqs_client.PubSub) *devUsecase {
	return &devUsecase{
		awsCl:  awsCl,
		pubsub: pubsub,
	}
}
