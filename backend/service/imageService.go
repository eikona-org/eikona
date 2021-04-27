package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/eikona-org/eikona/v2/data/repositories"
	"github.com/eikona-org/eikona/v2/storage"
	"github.com/eikona-org/eikona/v2/web/webmodels"
	"mime/multipart"
)

type ImageService interface {
	GetAllImages(email string) []webmodels.Image
	Insert(fileHeader *multipart.FileHeader, email string) error
}

type imageService struct {
	imageRepository repositories.ImageRepository
	userRepository  repositories.UserRepository
	orgRepository repositories.OrganizationRepository
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
	org := service.orgRepository.Find(user.OrganizationId)
	if org == nil {
		return errors.New("organization not found")
	}

	minioName := uuid.NewString()
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	service.storageClient.CreateObject(org.MinioBucketName, minioName, file, fileHeader.Size)

	image := service.imageRepository.Insert(fileHeader.Filename, minioName, org.OrganizationId)
	if image == nil {
		service.storageClient.RemoveObject(user.Organization.MinioBucketName, minioName)
		return errors.New("error while saving image to database")
	}

	return nil
}
