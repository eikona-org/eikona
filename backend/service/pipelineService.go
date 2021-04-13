package service

import (
	"github.com/disintegration/gift"
	data "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/helper"
	"image"
)

type PipelineService interface {
	Process() *helper.ImageWrapper
}

type pipelineService struct {
	imageWrapper    *helper.ImageWrapper
	processingSteps []data.ProcessingStep
	pipeline        *gift.GIFT
}

func NewPipelineService(imgWrapper *helper.ImageWrapper, procSteps []data.ProcessingStep) PipelineService {
	return &pipelineService{
		imageWrapper:    imgWrapper,
		processingSteps: procSteps,
		pipeline:        gift.New(),
	}
}

func (service *pipelineService) Process() *helper.ImageWrapper {
	service.imageWrapper.ProcessedImage = image.NewRGBA(service.pipeline.Bounds(service.imageWrapper.Image.Bounds()))

	service.pipeline.Draw(service.imageWrapper.ProcessedImage, service.imageWrapper.Image)

	return service.imageWrapper
}
