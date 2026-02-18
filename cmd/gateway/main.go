package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	//r.Use(Auth())
	//r.Use(RateLimit())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"GET request time: ": time.Now().String(), "formatted": time.Now().Format("2006-01-02 15:04:05")})
	})

	r.Run(":8080")
}
