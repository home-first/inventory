package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/home-first/inventory/controllers"
)

func documentRoute(superRoute *gin.RouterGroup) {
	superRoute.GET("/documents", controllers.DocumentController.ListDocument)
	superRoute.POST("/documents", controllers.DocumentController.CreateDocument)
	superRoute.GET("/documents/:id", controllers.DocumentController.GetDocument)
	superRoute.GET("/documents/:id/data", controllers.DocumentController.GetDocumentData)
}
