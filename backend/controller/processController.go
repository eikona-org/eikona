package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/service"
)

type ProcessController interface {
	ListAllProcesses(context *gin.Context)
	ListAllProcessingStepTypes(context *gin.Context)
}

type processController struct {
	processService service.ProcessService
	jwtService     service.JWTService
}

func NewProcessController(processServ service.ProcessService, jwtServ service.JWTService) ProcessController {
	return &processController{
		processService: processServ,
		jwtService:     jwtServ,
	}
}

func (c *processController) ListAllProcesses(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	email := c.getEmailByToken(authHeader)
	var processes = c.processService.GetAllProcesses(email)
	if processes == nil {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.JSON(http.StatusOK, processes)
}

func (c *processController) ListAllProcessingStepTypes(context *gin.Context) {
	var processes = c.processService.GetAllProcessingStepTypes()
	if processes == nil {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.JSON(http.StatusOK, processes)
}

func (c *processController) getEmailByToken(token string) string {
	const BEARER_SCHEMA = "Bearer "
	tokenString := token[len(BEARER_SCHEMA):]
	aToken, err := c.jwtService.ValidateToken(tokenString)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	email := fmt.Sprintf("%v", claims["email"])
	return email
}
