package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadDumpToS3(cfg Config, file string) {
	ctx := context.Background()

	awsCfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AWSKey, cfg.AWSSecret, ""),
		),
	)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	client := s3.NewFromConfig(awsCfg)

	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(cfg.AWSBucket),
		Key:    aws.String(filepath.Base(file)),
		Body:   f,
	})

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}
