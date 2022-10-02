package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/home-first/inventory/controllers"
)

func itemRoute(superRoute *gin.RouterGroup) {
	superRoute.GET("/items", controllers.ItemController.ListItems)
	superRoute.POST("/items", controllers.ItemController.CreateItem)
	superRoute.GET("/items/:id", controllers.ItemController.GetItem)
	superRoute.PUT("/items/:id", controllers.ItemController.UpdateItem)
}
