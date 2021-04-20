package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/poc"
	"github.com/imgProcessing/backend/v2/service"
	"github.com/imgProcessing/backend/v2/storage"
	"net/http"
)

var (
	imgRepo           = repositories.ImageRepository{}
	orgRepo           = repositories.OrganizationRepository{}
	procRepo          = repositories.ProcessRepository{}
	userRepo          = repositories.UserRepository{}
	storageClient     = storage.NewClient()
	authService       = service.NewAuthService(userRepo, orgRepo, storageClient)
	jwtService        = service.NewJWTService()
	renderService     = service.NewRenderService(imgRepo, procRepo, storageClient)
	imageService      = service.NewImageService(imgRepo, userRepo)
	processService    = service.NewProcessService(procRepo, userRepo)
	authController    = controller.NewAuthController(authService, jwtService)
	renderController  = controller.NewRenderController(renderService)
	imageController   = controller.NewImageController(imageService, jwtService)
	processController = controller.NewProcessController(processService, jwtService)
)

func Serve() {
	server := gin.Default()
	// Used for simpler developing - can be removed later or adjusted only for public api route
	server.Use(CORS)

	publicEndpoints := server.Group("/api")
	{
		// TODO: Remove when not needed
		server.GET("/api/poc", process)
		// Login -> POST /api/login
		server.POST("/api/login", authController.Login)
		// Register -> POST /api/register
		server.POST("/api/register", authController.Register)
		// Render -> GET /api/render/<img-id>/<proc-id>
		server.GET("/api/render/:identifier/:process", renderController.Render)
	}

	protectedEndpoints := publicEndpoints.Group("/auth", middleware.AuthorizeJWT(jwtService))
	{
		// -> GET /api/auth/images
		protectedEndpoints.GET("/images", imageController.ListAllImages)
		// -> GET /api/auth/processes
		protectedEndpoints.GET("/processes", processController.ListAllProcesses)
		// -> GET /api/auth/processingsteptypes
		protectedEndpoints.GET("/processingsteptypes", processController.ListAllProcessingStepTypes)
		// -> POST /api/auth/upload
		protectedEndpoints.GET("/upload", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Upload",
			})
		})
	}

	server.Run(":8080") //TODO: Make this configurable
}

// Used for simpler developing - can be removed later or adjusted only for public api route
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	//c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

// TODO: Remove after poc is redundant
func process(c *gin.Context) {
	data := poc.Process(c.Request.URL.Query())

	c.DataFromReader(
		http.StatusOK,
		int64(len(data.Bytes())),
		"image/png",
		data,
		nil,
	)
}
