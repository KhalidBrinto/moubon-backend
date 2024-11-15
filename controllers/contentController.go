package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBannerImages(c *gin.Context) {

	var content *models.ContentImage

	// Bind the incoming JSON to the Category struct
	if err := c.BindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the category into the database
	if err := config.DB.Create(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the newly created category
	c.JSON(http.StatusCreated, gin.H{"message": "content added successfully"})
}

func GetBannerImages(c *gin.Context) {
	response := map[string]interface{}{
		"banner": []string{},
	}

	var content []*models.ContentImage

	config.DB.Model(&content).Find(&content).Order("id DESC")

	if len(content) != 0 {
		for _, image := range content {
			switch image.Position {
			case "banner":
				response["banner"] = append(response["banner"].([]string), image.Image)
			}

		}
	}
	c.JSON(http.StatusOK, response)
}
