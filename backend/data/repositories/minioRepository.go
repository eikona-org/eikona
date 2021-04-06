package repositories

import (
	"context"
	"github.com/minio/minio-go/v7"
)

type MinioRepository interface {
	CreateBucket(bucketName string)
	RemoveBucket(bucketName string)
}

type clientConnection struct {
	client *minio.Client
}

func NewMinioRepository(minioClient *minio.Client) MinioRepository {
	return &clientConnection{
		client: minioClient,
	}
}

func (m *clientConnection) CreateBucket(bucketName string) {
	err := m.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "eu-central-1"})
	if err != nil {
		exists, errBucketExists := m.client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			panic(exists)
		} else {
			panic(err)
		}
	}
}

func (m *clientConnection) RemoveBucket(bucketName string) {
	err := m.client.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		panic(err)
	}
}
