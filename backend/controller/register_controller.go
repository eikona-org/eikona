package controller

import (
	"github.com/imgProcessing/backend/v2/dto"
	"github.com/imgProcessing/backend/v2/service"

	"github.com/gin-gonic/gin"
)

type RegisterController interface {
	Register(ctx *gin.Context) string
}

type registerController struct {
	registerService service.RegisterService
}

func RegisterHandler(registerService service.RegisterService) RegisterController {
	return &registerController{
		registerService: registerService,
	}
}

func (controller *registerController) Register(ctx *gin.Context) string {
	var credential dto.RegisterInformation
	err := ctx.ShouldBind(&credential)
	print(err)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.registerService.RegisterUser(credential.Email, credential.Password)
	print(isUserAuthenticated)
	return "ok"
}