package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ShopRoutes(router *gin.Engine) {
	shop := router.Group("/api/shops")
	{
		shop.POST("/", middlewares.AuthMiddleware(), middlewares.CheckIfAdmin(), controllers.AddShop)
		shop.GET("", controllers.GetShops)
		shop.PUT("/:id/", middlewares.AuthMiddleware(), middlewares.CheckIfAdmin(), controllers.UpdateShop)
		shop.DELETE("/:id/", middlewares.AuthMiddleware(), middlewares.CheckIfAdmin(), controllers.DeleteShop)
	}
}
