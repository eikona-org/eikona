package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/service"
	"net/http"
)

type ImageController interface {
	ListAllImages(context *gin.Context)
	//AllProcess(context *gin.Context)
	//UploadImage(context *gin.Context)
	//UpdateImage(context *gin.Context)
	//DeleteImage(context *gin.Context)
	//CreateProcess(context *gin.Context)
	//UpdateProcess(context *gin.Context)
	//DeleteProcess(context *gin.Context)
}

type imageController struct {
	imageService service.ImageService
	jwtService   service.JWTService
}

func NewImageController(imgServ service.ImageService, jwtServ service.JWTService) ImageController {
	return &imageController{
		imageService: imgServ,
		jwtService:   jwtServ,
	}
}

func (c *imageController) ListAllImages(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	email := c.getEmailByToken(authHeader)
	var images = c.imageService.GetAllImages(email)
	if images == nil {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.JSON(http.StatusOK, images)
}

func (c *imageController) getEmailByToken(token string) string {
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
