package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/data"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/poc"
	"github.com/imgProcessing/backend/v2/service"
	"net/http"
)

var (
	userRepo       = repositories.UserRepository{}
	orgRepo        = repositories.OrganizationRepository{}
	minioClient    = data.InitMinioClient()
	minioRepo      = repositories.NewMinioRepository(minioClient)
	jwtService     = service.NewJWTService()
	authService    = service.NewAuthService(userRepo, orgRepo, minioRepo)
	authController = controller.NewAuthController(authService, jwtService)
)

func Serve() {
	server := gin.Default()
	server.GET("/api/ping", ping)
	server.GET("/api/poc", process)
	//Login -> POST /api/login
	server.POST("/api/login", authController.Login)
	//Register -> POST /api/register
	server.POST("/api/register", authController.Register)

	//Auth Path
	apiRoutes := server.Group("/api/auth", middleware.AuthorizeJWT(jwtService))
	{
		// -> POST /api/auth/upload
		apiRoutes.GET("/upload", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Upload",
			})
		})
	}
	server.Run(":8080") //TODO: Make this configurable

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

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
