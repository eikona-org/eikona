package service

import (
	"github.com/badoux/checkmail"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/storage"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user webmodels.RegisterInformation) datamodels.User
	FindByEmail(email string) *datamodels.User
	IsDuplicateEmail(email string) bool
	IsValidEmail(email string) bool
}

type authService struct {
	userRepository         repositories.UserRepository
	organizationRepository repositories.OrganizationRepository
	storageClient          storage.Client
}

func NewAuthService(userRep repositories.UserRepository, orgRep repositories.OrganizationRepository, client storage.Client) AuthService {
	return &authService{
		userRepository:         userRep,
		organizationRepository: orgRep,
		storageClient:          client,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(*datamodels.User); ok {
		comparedPassword := comparePassword(v.PasswordHashSalt, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user webmodels.RegisterInformation) datamodels.User {
	organization := service.organizationRepository.CreateNew(user.Email)
	service.storageClient.CreateBucket(organization.MinioBucketName)
	res, _ := service.userRepository.InsertOrUpdate(user.Email, []byte(user.Password), organization.OrganizationId)
	return *res
}

func (service *authService) FindByEmail(email string) *datamodels.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	return service.userRepository.IsDuplicateEmail(email)
}

func (service *authService) IsValidEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}
