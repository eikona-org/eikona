package controller

import (
	"github.com/gin-gonic/gin"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/service"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login godoc
// @Tags Authorization
// @Summary Login
// @Description Verify User Credentials returning a JSON Web Token
// @Accept  json
// @Produce  json
// @Param user body webmodels.LoginCredentials true "User Data"
// @Success 200 {string} string "Token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO webmodels.LoginCredentials
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(*datamodels.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.Email)
		ctx.JSON(http.StatusOK, gin.H{
			"token": generatedToken,
		})
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register godoc
// @Tags Authorization
// @Summary Register
// @Description Register a User with an name, email, password
// @Accept  json
// @Produce  json
// @Param user body webmodels.RegisterInformation true "User Data"
// @Success 201 {string} string "ok"
// @Failure 400,409 {string} string "error"
// @Router /register [post]
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO webmodels.RegisterInformation
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsValidEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid email", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		c.authService.CreateUser(registerDTO)
		ctx.JSON(http.StatusCreated, "Account created")
	}
}
