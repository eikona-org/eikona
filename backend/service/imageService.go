package service

import (
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/data/repositories"
)

type ImageService interface {
	AllImages(email string) []datamodels.Image
}

type imageService struct {
	imageRepository repositories.ImageRepository
	userRepository  repositories.UserRepository
}

func NewImageService(imageRepo repositories.ImageRepository, userRepo repositories.UserRepository) ImageService {
	return &imageService{
		imageRepository: imageRepo,
		userRepository:  userRepo,
	}
}

func (service *imageService) AllImages(email string) []datamodels.Image {
	user := service.userRepository.FindByEmail(email)
	return service.imageRepository.AllImages(user.OrganizationId)
}
