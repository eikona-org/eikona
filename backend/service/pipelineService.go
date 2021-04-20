package service

import (
	"encoding/json"
	"github.com/disintegration/gift"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/helper"
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
		if !isSupportedProcessingStepType(step.ProcessingStepType) {
			continue
		}

		// TODO: Refactor
		if step.ProcessingStepType == datamodels.Resize {
			service.applyResizeOperation(step.ParameterJson)
		}
	}
}

// TODO: Refactor
func isSupportedProcessingStepType(procStepType datamodels.ProcessingStepType) bool {
	return procStepType == datamodels.Resize
}

// TODO: Refactor
type resizeParameters struct {
	Width  int
	Height int
}

// TODO: Refactor
func (service *pipelineService) applyResizeOperation(params string) {
	b := []byte(params)
	var parameters resizeParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	service.pipeline.Add(
		gift.Resize(
			parameters.Height,
			parameters.Width,
			gift.LanczosResampling,
		),
	)
}
