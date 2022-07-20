package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	superRoute.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})
	itemRoute(superRoute)
	collectionRoute(superRoute)
	documentRoute(superRoute)
}
