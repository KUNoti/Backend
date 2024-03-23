package s3service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

type S3Service struct {
	uploader *s3manager.Uploader
	bucket   string
}

const (
	EventImageFolder = "event"
	//UserImageFolder  = "user"
)

type S3ServiceConfig struct {
	Region             string
	Bucket             string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}

func NewS3Service(config *S3ServiceConfig) (*S3Service, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "s3-user",
		Config: aws.Config{
			Region:      aws.String(config.Region),
			Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, ""),
		},
	})

	if err != nil {
		return nil, err
	}

	return &S3Service{
		uploader: s3manager.NewUploader(sess),
		bucket:   config.Bucket,
	}, nil
}

func (s *S3Service) Upload(folder string, file *multipart.FileHeader) (string, error) {
	fileExtension := strings.Split(file.Filename, ".")[1]
	f, err := file.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
		return "", err
	}

	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to close file")
		}
	}(f)

	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate UUID")
		return "", err
	}

	result, err := s.uploader.UploadWithContext(context.TODO(), &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s-%s", folder, id.String(), fileExtension)),
		Body:   f,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to upload file")
		return "", err
	}

	return result.Location, nil
}
