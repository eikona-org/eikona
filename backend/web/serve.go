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
	imgRepo          = repositories.ImageRepository{}
	orgRepo          = repositories.OrganizationRepository{}
	procRepo         = repositories.ProcessRepository{}
	userRepo         = repositories.UserRepository{}
	storageClient    = storage.NewClient()
	authService      = service.NewAuthService(userRepo, orgRepo, storageClient)
	jwtService       = service.NewJWTService()
	renderService    = service.NewRenderService(imgRepo, procRepo, storageClient)
	imageService    = service.NewImageService(imgRepo, userRepo)
	authController   = controller.NewAuthController(authService, jwtService)
	renderController = controller.NewRenderController(renderService)
	imageController = controller.NewImageController(imageService, jwtService)
)

func Serve() {
	server := gin.Default()
	// Used for simpler developing - can be removed later or adjusted only for public api route
	server.Use(CORS)
	server.GET("/api/ping", ping)
	server.GET("/api/poc", process)
	// Login -> POST /api/login
	server.POST("/api/login", authController.Login)
	// Register -> POST /api/register
	server.POST("/api/register", authController.Register)
	// Render -> GET /api/render/<org-id>/<img-id>/<proc-id>
	server.GET("/api/render/:organization/:identifier/:process", renderController.Render)

	// Auth Path
	apiRoutes := server.Group("/api/auth", middleware.AuthorizeJWT(jwtService))
	{
		// -> GET /api/auth/getAllImages
		apiRoutes.GET("/getAllImages", imageController.AllImages)
		// -> POST /api/auth/upload
		apiRoutes.GET("/upload", func(ctx *gin.Context) {
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
	c.Header("Content-Type", "application/json")

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
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getimages(c *gin.Context) {
	id := [5]string{"1","2","3","4","5"}
	name := [5]string{"A","B","C","D","E"}

	parseData := make([]map[string]interface{}, 0, 0)

	for counter,_ := range id {
		var singleMap = make(map[string]interface{})
		singleMap["img"] = "https://pascalchristen.ch/images/thumbs/6.jpg"
		singleMap["id"] = id[counter]
		singleMap["name"] = name[counter]
		parseData = append(parseData, singleMap)
	}
	c.JSON(200, parseData)
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
