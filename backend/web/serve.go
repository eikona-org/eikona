package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/controller"
	"github.com/imgProcessing/backend/v2/data"
	models "github.com/imgProcessing/backend/v2/data/models"
	"github.com/imgProcessing/backend/v2/helper"
	"github.com/imgProcessing/backend/v2/middleware"
	"github.com/imgProcessing/backend/v2/poc"
	"github.com/imgProcessing/backend/v2/service"
	//"github.com/go-pg/pg/v10"
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
	server.Use(gin.Recovery(), gin.Logger())

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
	server.GET("/ping", ping)
	server.GET("/poc", process)
	server.Run(":8080") //TODO: Make this configurable
}

func ping(c *gin.Context) {
	database := data.GetDbConnection()
	defer database.Close() //IMPORTANT!!! Close so Connection doesn't stay open indefinitely
	transaction, transactionError := database.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	database.Model(&models.Organization{
		Name:            "Test Default",
		MinioBucketName: "Blub",
	}).Insert()
	transaction.Commit()
	org := new (models.Organization)
	database.Model(org).Where("name = ?", "Test Default").Select()
	transaction.Commit()
	//test, _ := uuid.FromString("6fa6b261-6187-4f33-9abc-da69ca5e57f4")
	database.Model(&models.User{
		LoginName: "pascal@christen.ch",
		Hash:      "$123",
		
		OrganizationId: org.OrganizationId,
	}).Insert()



	transaction.Commit()
	users := new(models.User)
	//err := database.Model(users).Where("login_name = ?", "pascal@christen.ch").Select()
	database.Model(users).Select()
	fmt.Println(users)
	transaction.Commit()

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
