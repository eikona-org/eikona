package service

import (
	data2 "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/web/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user models.RegisterInformation) data2.User
	FindByEmail(email string) data2.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userHelper helper.UserHelper
}

func NewAuthService(userHelper helper.UserHelper) AuthService {
	return &authService{
		userHelper: userHelper,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userHelper.VerifyCredential(email, password)
	if v, ok := res.(data2.User); ok {
		//TODO remove hash use from db -> v.hash
		//comparedPassword := comparePassword(v.PasswordHashSalt, []byte(password))
		comparedPassword := comparePassword("$2a$04$JKaM506hJ0RdnF7eOkEpHuTEeJJp9PbkVsK027bkhm6ibLyzSKPlW", []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user models.RegisterInformation) data2.User {
	userToCreate := data2.User{}
	res := service.userHelper.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) data2.User {
	return service.userHelper.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userHelper.IsDuplicateEmail(email)
	//TODO add !(==)
	return res == nil
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
