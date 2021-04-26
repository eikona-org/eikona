package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/service"
	"github.com/imgProcessing/backend/v2/storage"
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
	imageService      = service.NewImageService(imgRepo, userRepo, storageClient)
	processService    = service.NewProcessService(procRepo, userRepo)
	authController    = controller.NewAuthController(authService, jwtService)
	renderController  = controller.NewRenderController(renderService)
	imageController   = controller.NewImageController(imageService, jwtService)
	processController = controller.NewProcessController(processService, jwtService)
)

func Serve() {
	server := gin.Default()
	// Used for simpler developing - can be removed later or adjusted only for public api route
	server.Use(CORS())

	publicEndpoints := server.Group("/api")
	{
		// Login -> POST /api/login
		publicEndpoints.POST("/login", authController.Login)
		// Register -> POST /api/register
		publicEndpoints.POST("/register", authController.Register)
		// Dynamic render -> GET /api/render/dynamic/<img-id>?<params>
		publicEndpoints.GET("/render/dynamic/:identifier", renderController.DynamicRender)
		// Pipeline render -> GET /api/render/pipeline/<img-id>/<proc-id>
		publicEndpoints.GET("/render/pipeline/:identifier/:process", renderController.PipelineRender)
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
		protectedEndpoints.POST("/upload", imageController.UploadImage)
	}

	server.Run(":8080") //TODO: Make this configurable
}

// Used for simpler developing - can be removed later or adjusted only for public api route
// https://github.com/gin-contrib/cors/issues/29
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
