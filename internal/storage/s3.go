package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/BIC-Final-Project/backend/configs/env"
	awsConfig "github.com/BIC-Final-Project/backend/internal/aws"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Service interface {
	Upload(file *multipart.FileHeader, folder string) (*S3Response, error)
	Update(oldKey string, file *multipart.FileHeader, folder string) (*S3Response, error)
	Delete(key string) error
}

type s3Service struct {
	env env.EnvVars
}

type S3Response struct {
	Key      *string
	Location *string
}

const bucketName = "sportgather-community"

// Upload implements S3Service.
func (s *s3Service) Upload(file *multipart.FileHeader, folderName string) (*S3Response, error) {
	// Check if the file is nil or empty
	if file == nil || file.Size == 0 {
		return nil, nil
	}

	cfg, err := awsConfig.AwsConfig()
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	imageKey, err := utils.GenerateCUID()
	if err != nil {
		return nil, err
	}

	img, err := file.Open()
	if err != nil {
		return nil, err
	}

	format := filepath.Ext(file.Filename)
	imageKey = fmt.Sprintf("%s/%s%s", folderName, imageKey, format)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(imageKey),
		Body:        img,
		ContentType: aws.String(file.Header.Get("Content-Type")),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return nil, err
	}

	return &S3Response{
		Key:      aws.String(imageKey),
		Location: &result.Location,
	}, nil
}

// Update implements S3Service.
func (s *s3Service) Update(oldImageKey string, file *multipart.FileHeader, folder string) (*S3Response, error) {
	errChan := make(chan error, 2)
	resChan := make(chan *S3Response, 1)

	go func() {
		if err := s.Delete(oldImageKey); err != nil {
			errChan <- err
		}
		errChan <- nil
	}()

	go func() {
		res, err := s.Upload(file, folder)
		if err != nil {
			errChan <- err
		}
		resChan <- res
		errChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return nil, err
		}
	}

	return <-resChan, nil
}

// Delete implements S3Service.
func (s *s3Service) Delete(imageKey string) error {
	cfg, err := awsConfig.AwsConfig()
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imageKey),
	})
	if err != nil {
		return err
	}

	if err := InvalidateImage(imageKey); err != nil {
		return err
	}

	return nil
}

func NewS3Service(env env.EnvVars) S3Service {
	return &s3Service{env}
}
