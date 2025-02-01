package aws_client

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
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

func (c *AWSClient) UploadImageFile(ctx context.Context, fileRequest io.Reader, filePath string, fileHeader *multipart.FileHeader, width, height uint) (resultLocation *string, err error) {
	err = os.MkdirAll("./temp-images", os.ModePerm)
	if err != nil {
		return nil, err
	}

	if err = ctx.Err(); err != nil {
		return nil, err
	}

	tempFile, err := os.CreateTemp("./temp-images", "upload-*.png")
	if err != nil {
		return nil, fmt.Errorf("error during create temp file : %w", err)
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	imageFile, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("open file header issue to : %w", err)
	}
	defer imageFile.Close()

	fileBytes, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return nil, err
	}
	if err = ctx.Err(); err != nil {
		return nil, fmt.Errorf("operation canceled before writing to temp file: %w", err)
	}

	if _, err = tempFile.Write(fileBytes); err != nil {
		return nil, fmt.Errorf("failed to write to temp file: %w", err)
	}

	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("operation canceled before resizing image: %w", err)
	}

	err = utils.ResizeImage(tempFile.Name(), tempFile.Name(), width, height)
	if err != nil {
		return nil, err
	}

	fileExt := filepath.Ext(fileHeader.Filename)
	renamedImage := uuid.NewString() + fileExt
	if err = ctx.Err(); err != nil {
		return nil, fmt.Errorf("operation canceled before opening resized file: %w", err)
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err = ctx.Err(); err != nil {
		return nil, fmt.Errorf("operation canceled before uploading to S3: %w", err)
	}

	result, err := s3manager.NewUploader(c.SessionClient()).Upload(&s3manager.UploadInput{
		Bucket: aws.String("library-api-image"),
		Key:    aws.String(filePath + renamedImage),
		Body:   file,
	})
	if err != nil {
		log.Println("uploader issue ", err)
		return nil, err
	}

	return &result.Location, nil
}

func (c *AWSClient) UploadFileToS3(ctx context.Context, fileName, path string) (string, error) {
	buckets, err := c.ListBucket()
	if err != nil {
		return "", err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err = ctx.Err(); err != nil {
		return "", err
	}

	result, err := s3manager.NewUploader(c.SessionClient()).Upload(&s3manager.UploadInput{
		Bucket: aws.String(*buckets.Buckets[0].Name),
		Key:    aws.String(path + fileName),
		Body:   file,
	})
	if err != nil {
		return "", nil
	}
	return result.Location, nil
}
