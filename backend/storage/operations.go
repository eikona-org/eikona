package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (m *client) CreateBucket(bucketName string) {
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

func (m *client) RemoveBucket(bucketName string) {
	err := m.client.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		panic(err)
	}
}
