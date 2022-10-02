package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/home-first/inventory/database"
	"gorm.io/gorm"
)

type collectionController struct{}

var CollectionController collectionController

func (collectionController) ListCollection(ctx *gin.Context) {
	var collections []db.Collection
	result := db.DB.Model(&collections).Preload("Items").Find(&collections)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	minimalCollections := make([]CollectionMinimal, len(collections))

	for i, collection := range collections {
		minimalCollections[i] = MinimizeCollection(collection)
	}

	ctx.JSON(200, minimalCollections)
}

type CreateCollectionRequest struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

type CollectionMinimal struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

func MinimizeCollection(collection db.Collection) CollectionMinimal {
	items := make([]string, len(collection.Items))
	for i, item := range collection.Items {
		items[i] = item.ID.String()
	}
	return CollectionMinimal{
		ID:    collection.ID.String(),
		Name:  collection.Name,
		Items: items,
	}
}

type CollectionDetail struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	CreatedAt int      `json:"created_at"`
	UpdatedAt int      `json:"updated_at"`
	Items     []string `json:"items"`
}

func (collectionController) CreateCollection(ctx *gin.Context) {
	var createRequestBody CreateCollectionRequest
	err := ctx.BindJSON(&createRequestBody)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	collection := db.Collection{
		Name: createRequestBody.Name,
	}
	var items []db.Item
	for _, itemID := range createRequestBody.Items {
		var item db.Item
		result := db.DB.First(&item, "id = ?", itemID)
		if result.Error != nil {
			ctx.JSON(500, gin.H{
				"error": result.Error,
			})
			return
		}
		items = append(items, item)
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&collection)
		if result.Error != nil {
			return result.Error
		}
		err := tx.Model(&collection).Association("Items").Append(items)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(201, &collection)
}

func (collectionController) GetCollection(ctx *gin.Context) {
	var collection db.Collection

	idString := ctx.Param("id")

	_, err := uuid.Parse(idString)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	result := db.DB.Model(&collection).Preload("Items").First(&collection, "id = ?", idString)
	switch result.Error {
	case nil:
		ctx.JSON(200, MinimizeCollection(collection))
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

func (collectionController) UpdateCollection(ctx *gin.Context) {

	var collection db.Collection
	idString := ctx.Param("id")

	_, err := uuid.Parse(idString)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	result := db.DB.Model(&collection).First(&collection, "id = ?", idString)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	var updateRequestBody struct {
		Name string `json:"name"`
	}
	err = ctx.BindJSON(&updateRequestBody)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	collection.Name = updateRequestBody.Name
	result = db.DB.Save(&collection)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	ctx.JSON(200, collection)
}

// List the Items in a Collection
func (collectionController) ListItems(ctx *gin.Context) {
	var collection db.Collection
	collectionIDString := ctx.Param("id")
	_, err := uuid.Parse(collectionIDString)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	result := db.DB.Model(&collection).Preload("Items").Preload("Items.Collections").First(&collection, "id = ?", collectionIDString).First(&collection)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	items := make([]ItemMinimal, len(collection.Items))

	for i, item := range collection.Items {
		collections := make([]string, len(item.Collections))
		for j, collection := range item.Collections {
			collections[j] = collection.ID.String()
		}
		items[i] = ItemMinimal{
			ID:          item.ID.String(),
			Name:        item.Name,
			Collections: collections,
		}
	}

	ctx.JSON(200, items)
}

//Add an Item to a Collection
func (collectionController) AddItems(ctx *gin.Context) {
	// var collection db.Collection
	collectionIDString := ctx.Param("id")
	_, err := uuid.Parse(collectionIDString)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	// Check if all of the items exist
	var newItemIds []string

	ctx.BindJSON(&newItemIds)

	var items []db.Item
	db.DB.Where("id IN (?)", newItemIds).Find(&items)
	if len(items) != len(newItemIds) {
		ctx.JSON(400, gin.H{
			"error": "One or more items do not exist",
		})
		return
	}

	// Get the collection
	var collection db.Collection
	result := db.DB.Model(&collection).First(&collection, "id = ?", collectionIDString)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	err = db.DB.Debug().Model(&collection).Association("Items").Append(items)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, collection)
}
