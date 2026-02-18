package routes

import (
	"github.com/gin-gonic/gin"
	"go-edge/cmd/internal/types"
)

func AuthRoute(c *gin.Context) {
	var json types.Json

	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if json.User == "admin" || json.Password == "adminadmin" {
		c.JSON(200, gin.H{"token": "dummy-token"})
	} else {
		c.JSON(401, gin.H{"error": "unauthorized"})
	}
}
