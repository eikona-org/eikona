package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

func (c *client) CreateBucket(bucketName string) {
	err := c.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "eu-central-1"})
	if err != nil {
		exists, errBucketExists := c.client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			panic(exists)
		} else {
			panic(err)
		}
	}
}

func (c *client) RemoveBucket(bucketName string) {
	err := c.client.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		panic(err)
	}
}

func (c *client) CreateObject(bucketName string, objectName string, reader io.Reader, size int64) {
	_, err := c.client.PutObject(context.Background(), bucketName, objectName, reader, size, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
}

func (c *client) RemoveObject(bucketName string, objectName string) {
	err := c.client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		panic(err)
	}
}

func (c *client) GetObject(bucketName string, objectName string) *minio.Object {
	object, err := c.client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}

	return object
}
