package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"GET request time: ": time.DateTime})
	})

	r.Run(":8080")
}
