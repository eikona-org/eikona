package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/data"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/poc"
	"github.com/imgProcessing/backend/v2/service"
	//"github.com/go-pg/pg/v10"
	"net/http"
)

var (
	db             = data.Init2()
	userHelper     = helper.NewUserHelper(db)
	jwtService     = service.NewJWTService()
	authService    = service.NewAuthService(userHelper)
	authController = controller.NewAuthController(authService, jwtService)
)

func Serve() {
	server := gin.Default()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/ping", ping)

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
	server.GET("/ping", ping)
	server.GET("/poc", process)
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
