package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/service"
	//"github.com/go-pg/pg/v10"
	"net/http"
)

var (
	loginService    = service.StaticLoginService()
	jwtService      = service.JWTAuthService()
	loginController = controller.LoginHandler(loginService, jwtService)
)

func Serve() {

	server := gin.Default()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/ping", ping)

	//Login -> POST /api/login
	server.POST("/api/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
		}
	})

	//Register -> POST /api/register
	server.POST("/api/register", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Register",
		})
	})

	//Auth Path
	apiRoutes := server.Group("/api/auth", middleware.AuthorizeJWT())
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
	//You can work with dbContext in here e.g.
	//  database := data.GetDbConnection()
	//  defer database.Close() //IMPORTANT!!! Close so Connection doesn't stay open indefinitely
	//  transaction, transactionError := db.Begin()
	//  if transactionError != nil {
	//    panic(transactionError)
	//  }
	//  database.Model(&data.Image{
	//      Name: "TestImage",
	//  	DateUploaded: time.Now(),
	//  }).Insert()
	//  transaction.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
