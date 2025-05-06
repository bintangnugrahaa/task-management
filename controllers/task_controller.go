package controllers

import (
	"net/http"
	"tusk/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

func (t *TaskController) Create(c *gin.Context) {
	task := models.Task{}
	errBindJson := c.ShouldBindJSON(&task)
	if errBindJson != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errBindJson.Error()})
		return
	}

	errDB := t.DB.Create(&task).Error
	if errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDB.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}
