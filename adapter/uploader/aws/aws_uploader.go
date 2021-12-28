package aws_uploader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	comp "uploader/utils/image_compressor"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type AwsUploader struct {
	AWS_REGION     string
	AWS_ACCESS_KEY string
	AWS_SECRET     string
	S3_NAME        string
}

func NewAwsUploader(aws_region string, aws_access_key string, aws_secret string, s3_name string) *AwsUploader {
	return &AwsUploader{
		AWS_REGION:     aws_region,
		AWS_ACCESS_KEY: aws_access_key,
		AWS_SECRET:     aws_secret,
		S3_NAME:        s3_name,
	}
}

func putFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func upload(image []byte, name string, extension string, u *AwsUploader) (int64, error) {
	region := u.AWS_REGION
	access_key := u.AWS_ACCESS_KEY
	secret := u.AWS_SECRET
	bucket := u.S3_NAME

	Body := bytes.NewReader(image)

	input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &name,
		Body:   Body,
		ACL:    "public-read",
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(access_key, secret, "")),
	)
	if err != nil {
		return 0, errors.New("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	_, err = putFile(context.TODO(), client, input)
	if err != nil {
		return 0, errors.New("Got error uploading file:" + err.Error())
	}

	return Body.Size(), nil
}

func (u *AwsUploader) UploadImage(image []byte, id string, extension string) (int64, error) {
	compressedImage := comp.CompressImage(image)
	filename := fmt.Sprintf("%s.%s", id, extension)

	size, err := upload(compressedImage, filename, extension, u)

	if err != nil {
		return 0, err
	}

	return size, nil
}
