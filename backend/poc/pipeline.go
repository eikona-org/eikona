package poc

import (
	"bytes"
	"github.com/disintegration/gift"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// TODO: For now static file serving
// TODO: Minio sdk for future with caching for processed images
// TODO: Currently only png
// TODO: dynamic application of filters, dynamic image selection
func Process(queryArguments map[string][]string) []byte {
	pipeline := gift.New()

	applyQueryOperations(pipeline, queryArguments)

	absPathRaw, _ := filepath.Abs("./poc/fixtures/raw/image2.png")

	src := loadImage(absPathRaw)

	dst := image.NewRGBA(pipeline.Bounds(src.Bounds()))

	pipeline.Draw(dst, src)

	buffer := new(bytes.Buffer)
	encodeError := png.Encode(buffer, dst)
	if encodeError != nil {
		log.Fatalf("png.Encode failed: %v", encodeError)
	}

	return buffer.Bytes()
}

// TODO: check for mime type, only support jpeg and png
func loadImage(filename string) image.Image {
	file, fileError := os.Open(filename)
	if fileError != nil {
		log.Fatalf("os.Open failed: %v", fileError)
	}
	defer file.Close()
	img, _, decodeError := image.Decode(file)
	if decodeError != nil {
		log.Fatalf("image.Decode failed: %v", decodeError)
	}
	return img
}

// TODO: image quality settings via image encoding
func saveImage(filename string, img image.Image) {
	file, fileError := os.Create(filename)
	if fileError != nil {
		log.Fatalf("os.Create failed: %v", fileError)
	}
	defer file.Close()
	encodeError := png.Encode(file, img)
	if encodeError != nil {
		log.Fatalf("png.Encode failed: %v", encodeError)
	}
}