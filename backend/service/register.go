package service

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	RegisterUser(email string, password string) string
}
type registerInformation struct {
	email    string
	password string
	PasswordHash	string
}

func DatabaseRegisterService() RegisterService {
	return &registerInformation{
		email:    "pascal",
		password: "testing",
	}
}

func (info *registerInformation) RegisterUser(email string, password string) string {
	if !checkEmail(email) {
		return "Email is not valid"
	}
	registerInformation.hashAndSalt(*info,password)
	return "true"
}

func (u registerInformation) hashAndSalt(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	print(u.PasswordHash)
	print(u.email)
	return nil
}

func checkEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}