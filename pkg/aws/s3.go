package aws_client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func (c *AWSClient) CreateBucket(name string) error {
	_, err := c.s3.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *AWSClient) ListBucket() (result *s3.ListBucketsOutput, err error) {
	result = &s3.ListBucketsOutput{}
	result, err = c.s3.ListBuckets(nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *AWSClient) UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) (resultLocation *string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	return &result.Location, nil
}
