package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

type Client interface {
	CreateBucket(bucketName string)
	RemoveBucket(bucketName string)
	CreateObject(bucketName string, objectName string, reader io.Reader, size int64)
	RemoveObject(bucketName string, objectName string)
	GetObject(bucketName string, objectName string) *minio.Object
}

type client struct {
	client *minio.Client
}

func NewClient() Client {
	minioClient := initMinioClient()

	return &client{
		client: minioClient,
	}
}

func initMinioClient() *minio.Client {
	endpoint := os.Getenv("MINIO_HOST")
	accessKeyID := os.Getenv("MINIO_USER")
	secretAccessKey := os.Getenv("MINIO_PASSWORD")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // TODO: This is currently only for dev
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
