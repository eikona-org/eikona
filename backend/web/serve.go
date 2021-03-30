package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/data/repositories"
	"github.com/imgProcessing/backend/v2/poc"
	"net/http"
)

func Serve() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/poc", process)
	r.GET("/userRepoTest", userRepoTest)
	r.Run(":8080") //TODO: Make this configurable
}

func ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func process(c *gin.Context){
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