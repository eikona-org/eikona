package service

import (
	"github.com/imgProcessing/backend/v2/data/repositories"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
)

type ProcessService interface {
	AllProcesses(email string) []webmodels.Process
}

type processService struct {
	processRepository repositories.ProcessRepository
	userRepository  repositories.UserRepository
}

func NewProcessService(processRepo repositories.ProcessRepository, userRepo repositories.UserRepository) ProcessService {
	return &processService{
		processRepository: processRepo,
		userRepository:  userRepo,
	}
}

func (service *processService) AllProcesses(email string) []webmodels.Process {
	user := service.userRepository.FindByEmail(email)
	return service.processRepository.AllProcesses(user.OrganizationId)
}
