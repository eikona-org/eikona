package helper

import (
	"bytes"
	"fmt"
	"github.com/minio/minio-go/v7"
	"image"
	"image/jpeg"
	"image/png"
)

const ImageTypePng string = "png"
const ImageTypeJpeg string = "jpeg"

func isSupportedImageType(imgType string) bool {
	return imgType == ImageTypePng || imgType == ImageTypeJpeg
}

type ImageWrapper struct {
	Image          image.Image
	ImageType      string
	ProcessedImage *image.RGBA
	EncodedImage   *bytes.Buffer
}

func (w *ImageWrapper) init(img image.Image, imgType string) {
	if !isSupportedImageType(imgType) {
		panic("Unsupported image format present")
	}

	w.Image = img
	w.ImageType = imgType
}

func (w *ImageWrapper) GetMimeType() string {
	return fmt.Sprintf("image/%s", w.ImageType)
}

func LoadImage(object *minio.Object) *ImageWrapper {
	defer object.Close()

	img, imgType, err := image.Decode(object)
	if err != nil {
		panic("Image decoding failed")
	}

	imgWrapper := new(ImageWrapper)
	imgWrapper.init(img, imgType)

	return imgWrapper
}

func EncodeImage(imgWrapper *ImageWrapper) {
	buffer := new(bytes.Buffer)

	switch imgWrapper.ImageType {
	case ImageTypePng:
		encodeError := png.Encode(buffer, imgWrapper.ProcessedImage)
		if encodeError != nil {
			panic("PNG encoding failed")
		}
	case ImageTypeJpeg:
		encodeError := jpeg.Encode(buffer, imgWrapper.ProcessedImage, nil)
		if encodeError != nil {
			panic("JPEG encoding failed")
		}
		break
	default:
		panic("Unsupported image format present")
	}

	imgWrapper.EncodedImage = buffer
}
