package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/home-first/inventory/database"
	"gorm.io/gorm"
)

type itemController struct{}

var ItemController itemController

func (itemController) ListItems(ctx *gin.Context) {
	var items []db.Item
	result := db.DB.Find(&items)
	ctx.JSON(200, gin.H{
		"count": result.RowsAffected,
	})
}

func (itemController) CreateItem(ctx *gin.Context) {
	var item db.Item
	err := ctx.BindJSON(&item)
	if err != nil {
		return
	}
	result := db.DB.Create(&item)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	ctx.JSON(201, item)
}

func (itemController) GetItem(ctx *gin.Context) {
	var item db.Item
	id := ctx.Param("id")

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid item id",
		})
		return
	}

	result := db.DB.First(&item, id)
	switch result.Error {
	case nil:
		ctx.JSON(200, item)
	case gorm.ErrRecordNotFound:
		ctx.JSON(404, gin.H{
			"error": "Item not found",
		})
	default:
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
	}

}
