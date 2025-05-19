package minioclient

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

func InitMinIO() {
	endpoint := "localhost:9000"
	accessKey := "miniodomain"
	secretKey := "miniodomain123"
	useSSL := false

	var err error
	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}
	bucketName := "evidence-tasks"
	ctx := context.Background()

	exists, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("Failed to check bucket existence: %v", err)
	}
	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
		log.Printf("Bucket %s created successfully", bucketName)
	} else {
		log.Printf("Bucket %s already exists", bucketName)
	}
}
