package poc

import (
	"bytes"
	"context"
	"github.com/disintegration/gift"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

const IMAGE_TYPE_PNG string = "png"
const IMAGE_TYPE_JPEG string = "jpeg"

func isSupportedImageType(imgType string) bool {
	return imgType == IMAGE_TYPE_PNG || imgType == IMAGE_TYPE_JPEG
}

type ImgWrapper struct {
	img     image.Image
	imgType string
}

func (t *ImgWrapper) init(img image.Image, imgType string) {
	if !isSupportedImageType(imgType) {
		log.Fatal("Unsupported image format present")
	}

	t.img = img
	t.imgType = imgType
}

func Process(queryArguments map[string][]string) *bytes.Buffer {
	pipeline := gift.New()

	applyQueryOperations(pipeline, queryArguments)

	sourceImage := loadImage("testbucket", "image1.jpg")
	//sourceImage := loadImage("testbucket", "image2.png")

	processedImage := image.NewRGBA(pipeline.Bounds(sourceImage.img.Bounds()))

	pipeline.Draw(processedImage, sourceImage.img)

	return encodeImage(sourceImage, processedImage)
}

func initializeMinioClient() *minio.Client {
	endpoint := os.Getenv("MINIO_HOST")
	accessKeyID := os.Getenv("MINIO_USER")
	secretAccessKey := os.Getenv("MINIO_PASSWORD")

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // Only for dev
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

func loadImage(bucketName string, objectName string) *ImgWrapper {
	minioClient := initializeMinioClient()

	reader, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer reader.Close()

	//reader.Seek(0, 0)

	img, imgType, decodeError := image.Decode(reader)
	if decodeError != nil {
		log.Fatalf("image.Decode failed: %v", decodeError)
	}

	imgWrapper := new(ImgWrapper)
	imgWrapper.init(img, imgType)

	return imgWrapper
}

func encodeImage(imgWrapper *ImgWrapper, processedImage *image.RGBA) *bytes.Buffer {
	buffer := new(bytes.Buffer)

	switch imgWrapper.imgType {
	case IMAGE_TYPE_PNG:
		encodeError := png.Encode(buffer, processedImage)
		if encodeError != nil {
			log.Fatalf("png.Encode failed: %v", encodeError)
		}
	case IMAGE_TYPE_JPEG:
		encodeError := jpeg.Encode(buffer, processedImage, nil)
		if encodeError != nil {
			log.Fatalf("jpeg.Encode failed: %v", encodeError)
		}
		break
	default:
		log.Fatal("Unsupported image format present")
	}

	return buffer
}
