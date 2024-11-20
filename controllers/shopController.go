package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

// CreateCategory creates a new category
func AddShop(c *gin.Context) {

	var shop *models.Shop

	// Bind the incoming JSON to the Category struct
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the category into the database
	if err := config.DB.Create(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the newly created category
	c.JSON(http.StatusCreated, gin.H{"message": "Shop added successfully"})
}

// GetCategories retrieves all categories with their products
func GetShops(c *gin.Context) {
	var shops []*models.Shop

	model := config.DB.Model(&shops)

	pg := paginate.New()
	page := pg.With(model).Request(c.Request).Response(&shops)

	if page.Error {
		c.JSON(http.StatusInternalServerError, gin.H{"error": page.ErrorMessage})
		return
	}

	// Return the categories list
	c.JSON(http.StatusOK, page)
}

// UpdateCategory updates a category by its ID
func UpdateShop(c *gin.Context) {

	shopID := c.Param("id")
	var shop *models.Shop

	// Fetch the category from the database
	if err := config.DB.First(&shop, shopID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Bind the updated data to the category
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated category
	if err := config.DB.Save(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated category
	c.JSON(http.StatusOK, gin.H{"message": "Shop updated"})
}

// DeleteCategory deletes a category by its ID
func DeleteShop(c *gin.Context) {

	shopID := c.Param("id")
	var shop *models.Category

	// Fetch the category from the database
	if err := config.DB.First(&shop, shopID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Delete the category from the database
	if err := config.DB.Delete(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusNoContent, gin.H{"message": "Shop deleted successfully"})
}
