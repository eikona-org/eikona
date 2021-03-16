package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/imgProcessing/backend/v2/poc"
	"net/http"
)

var context *pg.DB

func Serve(dbContext *pg.DB) {
	context = dbContext

	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/poc", process)
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

func process(c *gin.Context){
	data := poc.Process(c.Request.URL.Query())

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "filename=test.png")
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprintf("%d", len(data)))
	c.Writer.Write(data)
}
