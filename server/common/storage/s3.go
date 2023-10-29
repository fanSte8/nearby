package storage

import (
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Config struct {
	BucketName      string
	Region          string
	AccessKeyID     string
	AccessKeySecret string
}

type S3Storage struct {
	bucketName string
	s3Client   *s3.S3
}

func NewS3Storage(cfg S3Config) S3Storage {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
		},
	))

	return S3Storage{
		s3Client:   s3.New(sess),
		bucketName: cfg.BucketName,
	}
}

func (storage S3Storage) Save(key string, file multipart.File) error {
	_, err := storage.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(storage.bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func (storage S3Storage) GetURL(key string) (string, error) {
	req, _ := storage.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(storage.bucketName),
		Key:    aws.String(key),
	})

	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}
