package main

import (
	"net/http"
	"tusk/config"
	"tusk/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})
	config.CreateOwnerAccount(db)

	// router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Tusk API",
		})
	})

	router.Static("attachments", "./attachments")
	router.Run("192.168.80.1:8080")
}
