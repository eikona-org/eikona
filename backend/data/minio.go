package data

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

func InitMinioClient() *minio.Client {
	endpoint := os.Getenv("MINIO_HOST")
	accessKeyID := os.Getenv("MINIO_USER")
	secretAccessKey := os.Getenv("MINIO_PASSWORD")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // Only for dev
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
