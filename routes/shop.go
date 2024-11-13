package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ShopRoutes(router *gin.Engine) {
	shop := router.Group("/api/shops")
	{
		shop.POST("/", middlewares.AuthMiddleware(), middlewares.CheckIfAdmin(), controllers.AddBrand)
		shop.GET("", controllers.GetBrands)
		shop.PUT("/:id/", middlewares.AuthMiddleware(), middlewares.CheckIfAdmin(), controllers.UpdateBrand)
		shop.DELETE("/:id/", middlewares.AuthMiddleware(), middlewares.CheckIfAdmin(), controllers.DeleteBrand)
	}
}
