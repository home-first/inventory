package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/home-first/inventory/controllers"
)

func collectionRoute(superRoute *gin.RouterGroup) {
	superRoute.GET("/collections", controllers.CollectionController.ListCollection)
	superRoute.POST("/collections", controllers.CollectionController.CreateCollection)
	superRoute.GET("/collections/:id", controllers.CollectionController.GetCollection)
	superRoute.PUT("/collections/:id", controllers.CollectionController.UpdateCollection)
	superRoute.GET("/collections/:id/items", controllers.CollectionController.ListItems)
	superRoute.POST("/collections/:id/items", controllers.CollectionController.AddItems)
}
