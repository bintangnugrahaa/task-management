package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Tusk API",
		})
	})

	router.Run("192.168.80.1:8080")
}
