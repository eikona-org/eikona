package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/service"
	"net/http"
)

type RenderController interface {
	Render(ctx *gin.Context)
}

type renderController struct {
	renderService service.RenderService
}

func NewRenderController(renderService service.RenderService) RenderController {
	return &renderController{
		renderService: renderService,
	}
}

func (c *renderController) Render(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error has occured while rendering: %s\n", r)
			response := helper.BuildErrorResponse("Failed to process request", "An error occurred", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}()

	imgUuid := uuid.MustParse(ctx.Param("identifier"))
	procUuid := uuid.MustParse(ctx.Param("process"))

	imgWrapper := c.renderService.Render(imgUuid, procUuid)

	ctx.DataFromReader(
		http.StatusOK,
		int64(len(imgWrapper.EncodedImage.Bytes())),
		imgWrapper.GetMimeType(),
		imgWrapper.EncodedImage,
		nil,
	)
}
