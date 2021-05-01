package controller

import (
	"fmt"
	"github.com/eikona-org/eikona/v2/helper"
	"github.com/eikona-org/eikona/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type RenderController interface {
	DynamicRender(ctx *gin.Context)
	PipelineRender(ctx *gin.Context)
}

type renderController struct {
	renderService service.RenderService
	cacheService  service.CacheService
}

func NewRenderController(renderService service.RenderService, cacheService service.CacheService) RenderController {
	return &renderController{
		renderService: renderService,
		cacheService:  cacheService,
	}
}

// DynamicRender godoc
// @Tags Render
// @Summary Render image dynamically
// @Description Render an image dynamically based on the query parameters
// @Produce image/jpeg,image/png
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Param image-id path string true "Id of the image that should be rendered"
// @Param resize-w query int false "Resize width parameter"
// @Param resize-h query int false "Resize height parameter"
// @Router /render/dynamic/{image-id} [get]
func (c *renderController) DynamicRender(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error has occured while rendering: %s\n", r)
			response := helper.BuildErrorResponse("Failed to process request", "An error occurred", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}()

	imgUuid := uuid.MustParse(ctx.Param("identifier"))
	queryParameters := ctx.Request.URL.Query()

	cacheKey := fmt.Sprintf("%s/%s", imgUuid.String(), queryParameters.Encode())

	cacheValue, cacheHit := c.cacheService.CheckCache(cacheKey)

	var imgWrapper *helper.ImageWrapper
	if cacheHit {
		imgWrapper = &helper.ImageWrapper{ImageType: cacheValue.ImageType, EncodedImage: &cacheValue.EncodedImage}
	} else {
		imgWrapper = c.renderService.DynamicRender(imgUuid, helper.ExtractProcessingSteps(queryParameters))
		c.cacheService.AddToCache(cacheKey, imgWrapper)
	}

	ctx.DataFromReader(
		http.StatusOK,
		int64(len(imgWrapper.EncodedImage.Bytes())),
		imgWrapper.GetMimeType(),
		imgWrapper.EncodedImage,
		nil,
	)
}

// PipelineRender godoc
// @Tags Render
// @Summary Render image based on pipeline
// @Description Render an image based on a previously defined pipeline
// @Produce image/jpeg,image/png
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Param image-id path string true "Id of the image that should be rendered"
// @Param pipeline-id path string true "Id of the pipeline that should be used for rendering"
// @Router /render/pipeline/{image-id}/{pipeline-id} [get]
func (c *renderController) PipelineRender(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error has occured while rendering: %s\n", r)
			response := helper.BuildErrorResponse("Failed to process request", "An error occurred", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}()

	imgUuid := uuid.MustParse(ctx.Param("identifier"))
	procUuid := uuid.MustParse(ctx.Param("process"))

	cacheKey := fmt.Sprintf("%s/%s", imgUuid.String(), procUuid.String())

	cacheValue, cacheHit := c.cacheService.CheckCache(cacheKey)

	var imgWrapper *helper.ImageWrapper
	if cacheHit {
		imgWrapper = &helper.ImageWrapper{ImageType: cacheValue.ImageType, EncodedImage: &cacheValue.EncodedImage}
	} else {
		imgWrapper = c.renderService.PipelineRender(imgUuid, procUuid)
		c.cacheService.AddToCache(cacheKey, imgWrapper)
	}

	ctx.DataFromReader(
		http.StatusOK,
		int64(len(imgWrapper.EncodedImage.Bytes())),
		imgWrapper.GetMimeType(),
		imgWrapper.EncodedImage,
		nil,
	)
}
