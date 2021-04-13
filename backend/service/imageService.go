package service

import (
	"github.com/imgProcessing/backend/v2/data/repositories"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
)

type ImageService interface {
	AllImages(email string) []webmodels.Image
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

func (service *imageService) AllImages(email string) []webmodels.Image {
	user := service.userRepository.FindByEmail(email)
	return service.imageRepository.AllImages(user.OrganizationId)
}
