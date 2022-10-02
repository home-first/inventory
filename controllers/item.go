package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/home-first/inventory/database"
	"gorm.io/gorm"
)

type itemController struct{}

var ItemController itemController

type ItemMinimal struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Collections []string `json:"collections"`
}

type ItemDetail struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	CreatedAt   int      `json:"created_at"`
	UpdatedAt   int      `json:"updated_at"`
	Description string   `json:"description"`
	Collections []string `json:"collections"`
}

func (itemController) ListItems(ctx *gin.Context) {
	var items []db.Item
	result := db.DB.Model(&items).Preload("Collections").Find(&items)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	minimalItems := make([]ItemMinimal, len(items))

	for i, item := range items {
		collectionIds := make([]string, len(item.Collections))
		for j, collection := range item.Collections {
			collectionIds[j] = collection.ID.String()
		}
		minimalItems[i] = ItemMinimal{
			ID:          item.ID.String(),
			Name:        item.Name,
			Collections: collectionIds,
		}
	}

	ctx.JSON(200, minimalItems)
}

func (itemController) CreateItem(ctx *gin.Context) {
	var item db.Item
	err := ctx.BindJSON(&item)
	if err != nil {
		return
	}
	result := db.DB.Create(&item).Preload("Collections")

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	collectionIds := make([]string, len(item.Collections))
	for j, collection := range item.Collections {
		collectionIds[j] = collection.ID.String()
	}
	ctx.JSON(201, ItemDetail{
		ID:          item.ID.String(),
		Name:        item.Name,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		Description: item.Description,
		Collections: collectionIds,
	})
}

func (itemController) GetItem(ctx *gin.Context) {
	var item db.Item

	idString := ctx.Param("id")

	_, err := uuid.Parse(idString)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	result := db.DB.Model(&item).Preload("Collections").First(&item, "id = ?", idString)
	switch result.Error {
	case nil:
		collectionIds := make([]string, len(item.Collections))
		for j, collection := range item.Collections {
			collectionIds[j] = collection.ID.String()
		}
		ctx.JSON(200, ItemDetail{
			ID:          item.ID.String(),
			Name:        item.Name,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			Description: item.Description,
			Collections: collectionIds,
		})
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

func (itemController) UpdateItem(ctx *gin.Context) {
	var item db.Item

	idString := ctx.Param("id")

	_, err := uuid.Parse(idString)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	result := db.DB.First(&item, "id = ?", idString)
	switch result.Error {
	case nil:
		err := ctx.BindJSON(&item)
		if err != nil {
			return
		}
		result := db.DB.Save(&item)
		if result.Error != nil {
			ctx.JSON(500, gin.H{
				"error": result.Error,
			})
			return
		}
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
