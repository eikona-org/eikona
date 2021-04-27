package webmodels

type CreateProcess struct {
	Name string `form:"name" json:"name" binding:"required" example:"Test Process"`
}
