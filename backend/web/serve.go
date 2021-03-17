package web

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-pg/pg/v10"
	"net/http"
)

func Serve() {

	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":8080") //TODO: Make this configurable
}

func ping(c *gin.Context){
	//You can work with dbContext in here e.g.
	//  databaseContext := data.GetContext()
	//  databaseContext.Model(&data.Image{
	//  	Id: 0,
	//  	DateUploaded: time.Now(),
	//  	Hash: "SomeHash1234",
	//  	Location: "/path/to/image",
	//  }).Insert()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
