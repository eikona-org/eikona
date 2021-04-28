package service

import (
	datamodels "github.com/eikona-org/eikona/v2/data/datamodels"
	"github.com/eikona-org/eikona/v2/data/repositories"
	"github.com/eikona-org/eikona/v2/web/webmodels"
)

type ProcessService interface {
	GetAllProcesses(email string) []webmodels.Process
	GetAllProcessingStepTypes() []webmodels.ProcessingStepType
	CreateProcess(dto webmodels.CreateProcess, email string) webmodels.Process
}

type processService struct {
	processRepository repositories.ProcessRepository
	processingStepRepository repositories.ProcessingStepRepository
	userRepository    repositories.UserRepository
	orgRepository     repositories.OrganizationRepository
}

func NewProcessService(processRepo repositories.ProcessRepository, processingStepRepo repositories.ProcessingStepRepository, userRepo repositories.UserRepository) ProcessService {
	return &processService{
		processRepository: processRepo,
		processingStepRepository: processingStepRepo,
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

func (service *processService) CreateProcess(dto webmodels.CreateProcess, email string) webmodels.Process {
	user := service.userRepository.FindByEmail(email)
	org := service.orgRepository.Find(user.OrganizationId)

	process := service.processRepository.Create(dto.Name, org.OrganizationId)

	for i := 0; i < len(dto.ProcessingSteps); i++ {
		err := service.processingStepRepository.AddToProcess(
			process.ProcessId,
			dto.ProcessingSteps[i].ProcessingStepType,
			dto.ProcessingSteps[i].ParameterJson,
			dto.ProcessingSteps[i].ExecutionPosition)
		if err != nil {
			return webmodels.Process{}
		}
	}

	return webmodels.Process{
		ProcessId: process.ProcessId,
		Name:      process.Name,
	}
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
			Id:      datamodels.Crop,
			Name:    "Crop (top left)",
			Options: []string{"width", "height"},
		},
		{
			Id:      datamodels.CropCenter,
			Name:    "Crop (center)",
			Options: []string{"width", "height"},
		},
		{
			Id:      datamodels.Fill,
			Name:    "Fill",
			Options: []string{"width", "height"},
		},
		{
			Id:      datamodels.Fit,
			Name:    "Fit",
			Options: []string{"width", "height"},
		},
		{
			Id:      datamodels.ContrastSigmoid,
			Name:    "Sigmoid Contrast",
			Options: []string{"midpoint", "factor"},
		},
		{
			Id:      datamodels.FlipH,
			Name:    "Flip horizontally",
			Options: []string{},
		},
		{
			Id:      datamodels.FlipV,
			Name:    "Flip vertically",
			Options: []string{},
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
			Id:      datamodels.Rotate,
			Name:    "Rotate",
			Options: []string{"angle"},
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
		{
			Id:      datamodels.Thumbnail,
			Name:    "Thumbnail",
			Options: []string{},
		},
		{
			Id:      datamodels.Transpose,
			Name:    "Transpose",
			Options: []string{},
		},
		{
			Id:      datamodels.Transverse,
			Name:    "Transverse",
			Options: []string{},
		},
	}
}
