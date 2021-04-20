package service

import (
	"github.com/imgProcessing/backend/v2/data/repositories"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
)

type ImageService interface {
	GetAllImages(email string) []webmodels.Image
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
