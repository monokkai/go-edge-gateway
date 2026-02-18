package main

import (
	"github.com/gin-gonic/gin"
	"go-edge/cmd/internal/middleware"
	"go-edge/cmd/internal/routes"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.Auth())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"GET request time: ": time.Now().String(), "formatted": time.Now().Format("2006-01-02 15:04:05")})
	})

	r.POST("/auth", routes.AuthRoute)

	r.Run(":8080")
}
