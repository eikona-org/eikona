package dto

//Login credential
type LoginCredentials struct {
	Email    string `form:"username"`
	Password string `form:"password"`
}