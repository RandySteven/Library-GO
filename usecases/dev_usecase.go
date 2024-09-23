package usecases

import (
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
)

import (
	"context"
)

type devUsecase struct {
	awsCl *aws_client.AWSClient
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

var _ usecases_interfaces.DevUsecase = &devUsecase{}

func newDevUsecase(awsCl *aws_client.AWSClient) *devUsecase {
	return &devUsecase{
		awsCl: awsCl,
	}
}
