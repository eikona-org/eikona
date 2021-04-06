package service

import (
	"github.com/disintegration/gift"
	"github.com/google/uuid"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/storage"
	"image"
)

type RenderService interface {
	Render(orgUuid uuid.UUID, imgUuid uuid.UUID, procUuid uuid.UUID) *helper.ImageWrapper
}

type renderService struct {
	imageRepository   repositories.ImageRepository
	processRepository repositories.ProcessRepository
	storageClient     storage.Client
}

func NewRenderService(imgRep repositories.ImageRepository, procRep repositories.ProcessRepository, client storage.Client) RenderService {
	return &renderService{
		imageRepository:   imgRep,
		processRepository: procRep,
		storageClient:     client,
	}
}

func (service *renderService) Render(imgUuid uuid.UUID, orgUuid uuid.UUID, procUuid uuid.UUID) *helper.ImageWrapper {
	imageModel := service.getImage(imgUuid, orgUuid)
	processModel := service.getProcess(procUuid)

	if nil == imageModel || nil == processModel {
		panic("Invalid parameters")
	}

	imgWrapper := service.loadImage(imageModel.Owner.MinioBucketName, imageModel.MinioObjectName)

	return service.encodeImage(service.processImage(imgWrapper))
}

// TODO: Temporary, doesn't really do anything
func (service *renderService) processImage(imgWrapper *helper.ImageWrapper) *helper.ImageWrapper {
	pipeline := gift.New()
	imgWrapper.ProcessedImage = image.NewRGBA(pipeline.Bounds(imgWrapper.Image.Bounds()))

	pipeline.Draw(imgWrapper.ProcessedImage, imgWrapper.Image)

	return imgWrapper
}

func (service *renderService) loadImage(bucketName string, objectName string) *helper.ImageWrapper {
	object := service.storageClient.GetObject(bucketName, objectName)

	return helper.LoadImage(object)
}

func (service *renderService) encodeImage(imgWrapper *helper.ImageWrapper) *helper.ImageWrapper {
	helper.EncodeImage(imgWrapper)

	return imgWrapper
}

func (service *renderService) getImage(imgUuid uuid.UUID, orgUuid uuid.UUID) *datamodels.Image {
	return service.imageRepository.Find(imgUuid, orgUuid)
}

// TODO: Refactor when process is linked to a org
func (service *renderService) getProcess(procUuid uuid.UUID) *datamodels.Process {
	return service.processRepository.Find(procUuid)
}
