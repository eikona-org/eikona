package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/storage"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
	"mime/multipart"
)

type ImageService interface {
	GetAllImages(email string) []webmodels.Image
	Insert(fileHeader *multipart.FileHeader, email string) error
}

type imageService struct {
	imageRepository repositories.ImageRepository
	userRepository  repositories.UserRepository
	storageClient   storage.Client
}

func NewImageService(imageRepo repositories.ImageRepository, userRepo repositories.UserRepository, storage storage.Client) ImageService {
	return &imageService{
		imageRepository: imageRepo,
		userRepository:  userRepo,
		storageClient:   storage,
	}
}

func (service *imageService) GetAllImages(email string) []webmodels.Image {
	user := service.userRepository.FindByEmail(email)
	images := service.imageRepository.GetAll(user.OrganizationId)

	var apiImageModels []webmodels.Image

	for _, image := range *images {
		apiImageModels = append(apiImageModels, webmodels.Image{
			ImageId:  image.ImageId,
			Name:     image.Name,
			Uploaded: image.Uploaded,
		})
	}

	return apiImageModels
}

func (service *imageService) Insert(fileHeader *multipart.FileHeader, email string) error {
	user := service.userRepository.FindByEmail(email)
	if user == nil {
		return errors.New("user not found")
	}

	minioName := uuid.NewString()
	file, error := fileHeader.Open()
	if error != nil {
		return error
	}
	service.storageClient.CreateObject(user.Organization.MinioBucketName, minioName, file, fileHeader.Size)

	image := service.imageRepository.Insert(fileHeader.Filename, minioName, user.UserId)
	if image == nil {
		service.storageClient.RemoveObject(user.Organization.MinioBucketName, minioName)
		return errors.New("error while saving image to database")
	}

	return nil
}
