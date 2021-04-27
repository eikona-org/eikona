package webmodels

type LoginCredentials struct {
	Email    string `form:"email" json:"email" binding:"required" example:"test@testit.example"`
	Password string `form:"password" json:"password" binding:"required" example:"Secure4Life!;)"`
}
