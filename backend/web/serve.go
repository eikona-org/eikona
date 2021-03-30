package web

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/poc"
	"net/http"
)

func Serve() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/poc", process)
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
