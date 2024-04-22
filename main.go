package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", Hello)
	r.Run(":9001")
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "hello",
	})
}
