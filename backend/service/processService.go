package service

import (
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/data/repositories"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
)

type ProcessService interface {
	GetAllProcesses(email string) []webmodels.Process
	GetAllProcessingStepTypes() []webmodels.ProcessingStepType
}

type processService struct {
	processRepository repositories.ProcessRepository
	userRepository    repositories.UserRepository
}

func NewProcessService(processRepo repositories.ProcessRepository, userRepo repositories.UserRepository) ProcessService {
	return &processService{
		processRepository: processRepo,
		userRepository:    userRepo,
	}
}

func (service *processService) GetAllProcesses(email string) []webmodels.Process {

	user := service.userRepository.FindByEmail(email)
	processes := service.processRepository.GetAll(user.OrganizationId)

	var apiProcessModels []webmodels.Process

	for _, process := range *processes {
		apiProcessModels = append(apiProcessModels, webmodels.Process{
			ProcessId: process.ProcessId,
			Name:      process.Name,
		})
	}

	return apiProcessModels
}

func (service *processService) GetAllProcessingStepTypes() []webmodels.ProcessingStepType {
	// TODO: Maybe also move these into pipelineOperations
	return []webmodels.ProcessingStepType{
		{
			Id:      datamodels.Blur,
			Name:    "Blur",
			Options: []string{"sigma"},
		},
		{
			Id:      datamodels.Brightness,
			Name:    "Brightness",
			Options: []string{"percentage"},
		},
		{
			Id:      datamodels.Contrast,
			Name:    "Contrast",
			Options: []string{"percentage"},
		},
		{
			Id:      datamodels.ContrastSigmoid,
			Name:    "Sigmoid Contrast",
			Options: []string{"midpoint", "factor"},
		},
		{
			Id:      datamodels.Gamma,
			Name:    "Gamma",
			Options: []string{"gamma"},
		},
		{
			Id:      datamodels.Grayscale,
			Name:    "Grayscale",
			Options: []string{},
		},
		{
			Id:      datamodels.Hue,
			Name:    "Hue",
			Options: []string{"shift"},
		},
		{
			Id:      datamodels.Invert,
			Name:    "Invert",
			Options: []string{},
		},
		{
			Id:      datamodels.Resize,
			Name:    "Resizing",
			Options: []string{"width", "height"},
		},
		{
			Id:      datamodels.Rotate90,
			Name:    "Rotate 90 degrees",
			Options: []string{},
		},
		{
			Id:      datamodels.Rotate180,
			Name:    "Rotate 180 degrees",
			Options: []string{},
		},
		{
			Id:      datamodels.Rotate270,
			Name:    "Rotate 270 degrees",
			Options: []string{},
		},
		{
			Id:      datamodels.Saturation,
			Name:    "Saturation",
			Options: []string{"percentage"},
		},
		{
			Id:      datamodels.Sharpen,
			Name:    "Sharpen",
			Options: []string{"sigma", "amount", "threshold"},
		},
	}
}
