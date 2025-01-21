package aws_client

import (
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
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

func (c *AWSClient) UploadImageFile(fileRequest io.Reader, filePath string, fileHeader *multipart.FileHeader, width, height uint) (resultLocation *string, err error) {
	tempFile, err := ioutil.TempFile("./temp-images", "upload-*.png")
	if err != nil {
		log.Println("err temp file", err)
		return nil, err
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(fileRequest)
	if err != nil {
		log.Println("err file bytes", err)
		return nil, err
	}
	tempFile.Write(fileBytes)

	imageFile, err := fileHeader.Open()
	if err != nil {
		log.Println("err image header open ", err)
		return nil, err
	}
	defer imageFile.Close()

	err = utils.ResizeImage(tempFile.Name(), tempFile.Name(), width, height)
	if err != nil {
		log.Println("err resize : ", err)
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to resize image`, err)
	}

	renamedImage := utils.RenameFileWithDateAndUUID(tempFile.Name()[len(`./temp-images/`):])

	buckets, err := c.ListBucket()
	if err != nil {
		log.Println("err bucket ", err)
		return nil, err
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		log.Println("err file ", err)
		return nil, err
	}
	defer file.Close()

	result, err := s3manager.NewUploader(c.SessionClient()).Upload(&s3manager.UploadInput{
		Bucket: aws.String(*buckets.Buckets[0].Name),
		Key:    aws.String(filePath + renamedImage),
		Body:   file,
	})
	if err != nil {
		log.Println("uploader issue ", err)
		return nil, err
	}

	_ = os.Remove(tempFile.Name())

	return &result.Location, nil
}

func (c *AWSClient) UploadFileToS3(fileName, path string) (string, error) {
	buckets, err := c.ListBucket()
	if err != nil {
		return "", err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()
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
