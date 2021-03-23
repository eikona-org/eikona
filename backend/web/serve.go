package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/poc"
	//"github.com/go-pg/pg/v10"
	"net/http"
)

func Serve() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/poc", process)
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

func process(c *gin.Context){
	data := poc.Process(c.Request.URL.Query())

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "filename=test.png")
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprintf("%d", len(data)))
	c.Writer.Write(data)
}
