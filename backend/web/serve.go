package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"net/http"
)

var context *pg.DB

func Serve(dbContext *pg.DB) {
	context = dbContext

	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":8080")
}

func ping(c *gin.Context){
	//You can work with context in here e.g.
	//  context.Model(&data.Image{
	//  	Id: 0,
	//  	DateUploaded: time.Now(),
	//  	Hash: "SomeHash1234",
	//  	Location: "/path/to/image",
	//  }).Insert()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
