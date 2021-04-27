package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/eikona-org/eikona/v2/helper"
	"github.com/eikona-org/eikona/v2/service"
	"github.com/eikona-org/eikona/v2/web/webmodels"
	"net/http"
)

type ProcessController interface {
	ListAllProcesses(context *gin.Context)
	ListAllProcessingStepTypes(context *gin.Context)
	CreateProcess(context *gin.Context)
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

// ListAllProcesses godoc
// @Tags Processes
// @Summary List all organization processes
// @Description List all the processes of an organization
// @Security jwtAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} []webmodels.Process
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/processes [get]
func (c *processController) ListAllProcesses(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	email := c.getEmailByToken(authHeader)
	var processes = c.processService.GetAllProcesses(email)
	if processes == nil {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.JSON(http.StatusOK, processes)
}

// ListAllProcessingStepTypes godoc
// @Tags Processes
// @Summary List all processing steps
// @Description List all the avaialable processing steps
// @Security jwtAuth
// @Accept json
// @Produce json
// @Success 200 {object} []webmodels.ProcessingStepType
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/processingsteptypes [get]
func (c *processController) ListAllProcessingStepTypes(context *gin.Context) {
	var processes = c.processService.GetAllProcessingStepTypes()
	if processes == nil {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.JSON(http.StatusOK, processes)
}

// CreateProcess godoc
// @Tags Processes
// @Summary Create Process
// @Description Create a process
// @Security jwtAuth
// @Accept json
// @Produce json
// @Param name body webmodels.CreateProcess true "Name"
// @Success 200 {object} webmodels.Process
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/process [post]
func (c *processController) CreateProcess(context *gin.Context) {
	var createProzessDTO webmodels.CreateProcess
	errDTO := context.ShouldBind(&createProzessDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := context.GetHeader("Authorization")
	email := c.getEmailByToken(authHeader)

	var process = c.processService.CreateProcess(createProzessDTO, email)
	if process == (webmodels.Process{}) {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.JSON(http.StatusOK, process)
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
