package service

import (
	"github.com/disintegration/gift"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/pipelineOperations"
	"image"
)

type PipelineService interface {
	Process() *helper.ImageWrapper
}

type pipelineService struct {
	imageWrapper    *helper.ImageWrapper
	processingSteps []datamodels.ProcessingStep
	pipeline        *gift.GIFT
}

func NewPipelineService(imgWrapper *helper.ImageWrapper, procSteps []datamodels.ProcessingStep) PipelineService {
	return &pipelineService{
		imageWrapper:    imgWrapper,
		processingSteps: procSteps,
		pipeline:        gift.New(),
	}
}

func (service *pipelineService) Process() *helper.ImageWrapper {
	service.applyProcessingSteps()

	service.imageWrapper.ProcessedImage = image.NewRGBA(service.pipeline.Bounds(service.imageWrapper.Image.Bounds()))

	service.pipeline.Draw(service.imageWrapper.ProcessedImage, service.imageWrapper.Image)

	return service.imageWrapper
}

func (service *pipelineService) applyProcessingSteps() {
	for _, step := range service.processingSteps {
		service.applyOperation(step)
	}
}

func (service *pipelineService) applyOperation(procStep datamodels.ProcessingStep) {
	switch procStep.ProcessingStepType {
	case datamodels.Blur:
		pipelineOperations.ApplyBlurOperation(service.pipeline, procStep.ParameterJson)
		break
	case datamodels.Brightness:
		pipelineOperations.ApplyBrightnessOperation(service.pipeline, procStep.ParameterJson)
		break
	case datamodels.Contrast:
		pipelineOperations.ApplyContrastOperation(service.pipeline, procStep.ParameterJson)
		break
	case datamodels.Gamma:
		pipelineOperations.ApplyGammaOperation(service.pipeline, procStep.ParameterJson)
		break
	case datamodels.Grayscale:
		pipelineOperations.ApplyGrayscaleOperation(service.pipeline)
		break
	case datamodels.Invert:
		pipelineOperations.ApplyInvertOperation(service.pipeline)
		break
	case datamodels.Resize:
		pipelineOperations.ApplyResizeOperation(service.pipeline, procStep.ParameterJson)
		break
	case datamodels.Sharpen:
		pipelineOperations.ApplySharpenOperation(service.pipeline, procStep.ParameterJson)
		break
	}
}
