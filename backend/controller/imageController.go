package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/eikona-org/eikona/v2/service"
	_ "github.com/eikona-org/eikona/v2/web/webmodels"
	"net/http"
)

type ImageController interface {
	ListAllImages(context *gin.Context)
	UploadImage(context *gin.Context)
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

// ListAllImages godoc
// @Tags Images
// @Summary List all organization images
// @Description List all the images of an organization
// @Security jwtAuth
// @Accept json
// @Produce json
// @Success 200 {object} []webmodels.Image
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/images [get]
func (c *imageController) ListAllImages(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	email := c.getEmailByToken(authHeader)
	var images = c.imageService.GetAllImages(email)
	if images == nil {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.JSON(http.StatusOK, images)
}

// UploadImage godoc
// @Tags Images
// @Summary Upload an image
// @Description Upload an image
// @Security jwtAuth
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/upload [post]
func (c *imageController) UploadImage(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	email := c.getEmailByToken(authHeader)
	file, uploadError := context.FormFile("file") //TODO: Make sure frontend sets the same field name
	if uploadError != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occured while uploading the image",
			"error":   "No file received",
		})
		return
	}

	insertError := c.imageService.Insert(file, email)
	if insertError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error occured while saving the image",
			"error":   insertError.Error(),
		})
		return
	}

	context.Status(http.StatusOK)
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
