package models

type RegisterInformation struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
