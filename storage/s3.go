package storage

import (
	"OctaneServer/config"
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client *s3.Client
}

func (storage *S3) Open() error {
	cred := credentials.NewStaticCredentialsProvider(config.CurrentConfig.Storage.AccessKey, config.CurrentConfig.Storage.SecretKey, "")
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithCredentialsProvider(cred), awsConfig.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: config.CurrentConfig.Storage.Name,
		}, nil
	})))

	if err != nil {
		return err
	}

	storage.client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	return nil
}

func (storage *S3) Save(key string, file []byte) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(config.CurrentConfig.Storage.Name),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	}

	_, err := storage.client.PutObject(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (storage *S3) Delete(key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(config.CurrentConfig.Storage.Name),
		Key:    aws.String(key),
	}
	_, err := storage.client.DeleteObject(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}
