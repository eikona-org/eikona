package web

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-pg/pg/v10"
	"net/http"
	"github.com/imgProcessing/backend/v2/web/controller"
	"github.com/imgProcessing/backend/v2/web/middleware"
	"github.com/imgProcessing/backend/v2/web/service"
)

func Serve() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	r := gin.Default()
	r.GET("/ping", ping)
	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := r.Group("/api/test/", middleware.AuthorizeJWT())
	{
		apiRoutes.GET("/upload", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.POST("/api/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	r.Run(":8080") //TODO: Make this configurable
}

func ping(c *gin.Context){
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
