package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/data"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/poc"
	"github.com/imgProcessing/backend/v2/service"
	"net/http"
)

var (
	//dbConn         = data.GetDbConnection()
	dbConn         = data.Init2()
	userHelper     = helper.NewUserHelper(dbConn)
	jwtService     = service.NewJWTService()
	authService    = service.NewAuthService(userHelper)
	authController = controller.NewAuthController(authService, jwtService)
)

func Serve() {
	server := gin.Default()
	server.GET("/ping", ping)
	server.GET("/poc", process)
	server.GET("/userRepoTest", userRepoTest)
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

func ping(c *gin.Context){
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

func userRepoTest(c *gin.Context){
	testEmail := "testuser@imgprocessing.io"
	testPassword := "My5up3r53cr3tP455w0rd"
	testOrganizationName := "Testing Inc."
	userRepo := repositories.UserRepository{}
	organizationRepo := repositories.OrganizationRepository{}

	organization := organizationRepo.CreateNew(testOrganizationName)
	user, userFound := userRepo.InsertOrUpdate(testEmail, []byte(testPassword), organization.OrganizationId)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"insertSuccess": false,
		})
		return
	}

	passwordVerified := userRepo.VerifyCredential(testEmail, testPassword)

	c.JSON(http.StatusOK, gin.H{
		"insertSuccess": true,
		"existingUserFound": userFound,
		"passwordVerified": passwordVerified,
		"userEmail": user.Email,
	})
}