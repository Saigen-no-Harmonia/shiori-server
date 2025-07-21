package util

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	S3Client     *s3.Client
	s3BucketName string
	initOnce     sync.Once
	Presigner    *s3.PresignClient
)

// InitS3 sets up the AWS SDK clients (called once)
func InitS3() {
	initOnce.Do(func() {
		ctx := context.TODO()
		s3BucketName = os.Getenv("S3_BUCKET_NAME")
		if s3BucketName == "" {
			log.Fatal("S3_BUCKET_NAME is not set")
		}

		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("unable to load AWS SDK config, %v", err)
		}

		S3Client = s3.NewFromConfig(cfg)
		Presigner = s3.NewPresignClient(S3Client)
	})
}

func GetS3BucketName() string {
	return s3BucketName
}
