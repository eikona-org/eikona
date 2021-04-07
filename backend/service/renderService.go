package service

import (
	"github.com/google/uuid"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/storage"
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
	image := service.getImage(imgUuid, orgUuid)
	process := service.getProcess(procUuid)

	if nil == image || nil == process {
		panic("Invalid parameters")
	}

	imgWrapper := service.loadImage(image.Owner.MinioBucketName, image.MinioObjectName)

	return service.encodeImage(
		service.processImage(imgWrapper, process.ProcessingSteps),
	)
}

func (service *renderService) processImage(imgWrapper *helper.ImageWrapper, procSteps []datamodels.ProcessingStep) *helper.ImageWrapper {
	pipelineService := NewPipelineService(imgWrapper, procSteps)

	return pipelineService.Process()
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
