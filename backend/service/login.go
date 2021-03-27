package service

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)


type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "pascal",
		password: "testing",
	}
}

func DBLoginService() LoginService {
	return &loginInformation{
		//TODO get Password and Hash from DB
		email:    "pascal",
		password: "testing",
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	print(password)
	return info.email == email && comparePasswords(info.password, password)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePassword := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}