package controllers

import (
	"github.com/gin-gonic/gin"
	db "github.com/home-first/inventory/database"
	"gorm.io/gorm"
)

type collectionController struct{}

var CollectionController collectionController

func (collectionController) ListCollection(ctx *gin.Context) {
	var collections []db.Collection
	result := db.DB.Find(&collections)
	ctx.JSON(200, gin.H{
		"count": result.RowsAffected,
	})
}

func (collectionController) CreateCollection(ctx *gin.Context) {
	var collection db.Collection
	err := ctx.BindJSON(&collection)
	if err != nil {
		return
	}
	result := db.DB.Create(&collection)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	ctx.JSON(201, collection)
}

func (collectionController) GetCollection(ctx *gin.Context) {
	var collection db.Collection
	id := ctx.Param("id")

	result := db.DB.First(&collection, id)
	switch result.Error {
	case nil:
		ctx.JSON(200, collection)
	case gorm.ErrRecordNotFound:
		ctx.JSON(404, gin.H{
			"error": "Collection not found",
		})
	default:
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
	}
}
