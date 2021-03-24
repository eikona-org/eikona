package dto

type LoginCredentials struct {
	Email string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}