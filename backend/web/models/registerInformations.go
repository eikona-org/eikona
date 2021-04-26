package models

type RegisterInformation struct {
	Name     string `json:"name" form:"name" binding:"required" example:"Max Mustermann"`
	Email    string `form:"email" json:"email" binding:"required" example:"test@testit.example"`
	Password string `form:"password" json:"password" binding:"required" example:"Secure4Life!;)"`
}
