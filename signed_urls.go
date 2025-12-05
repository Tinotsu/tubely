package main

import (
	"time"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"context"
)

func generatePresignedURL(s3Client *s3.Client, bucket, key string, expireTime time.Duration) (string, error) {
	presignedClient := s3.NewPresignClient(s3Client)
	ctx := context.Background()
	expiration := s3.WithPresignExpires(expireTime)
	objectInput := new(s3.GetObjectInput)
	objectInput.Bucket = &bucket
	objectInput.Key = &key
	objectSigned, err := presignedClient.PresignGetObject(ctx, objectInput, expiration)
	return objectSigned.URL, err
}
