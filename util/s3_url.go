package util

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetS3AccessUrl returns a presigned URL for the given object key.
func GetS3AccessUrl(s3ObjectKey string) string {
	ctx := context.TODO()

	resp, err := Presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(GetS3BucketName()),
		Key:    aws.String(s3ObjectKey),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		log.Fatalf("unable to get presigned URL, %v", err)
	}

	return resp.URL
}
