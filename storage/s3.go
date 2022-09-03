package storage

import (
	"OctaneServer/config"
	"bytes"
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
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

func (storage *S3) HouseKeeping() {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(config.CurrentConfig.Storage.Name),
	}
	output, err := storage.client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Warn().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.HouseKeeping)
		return
	}
	for _, v := range output.Contents {
		delete := true
		for _, w := range output.Contents {
			if strings.HasPrefix(*w.Key, strings.Split(*v.Key, "_")[0]) && v.Key != w.Key && v.LastModified.Before(*w.LastModified) {
				delete = true
			}
		}
		if delete {
			err = storage.Delete(*v.Key)
			if err != nil {
				log.Warn().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.HouseKeeping)
				return
			}
		}
	}
}
