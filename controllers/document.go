package controllers

import (
	"github.com/gin-gonic/gin"
	db "github.com/home-first/inventory/database"
	"gorm.io/gorm"
)

type documentController struct{}

var DocumentController documentController

func (documentController) ListDocument(ctx *gin.Context) {
	var documents []db.Document
	result := db.DB.Find(&documents)
	ctx.JSON(200, gin.H{
		"count": result.RowsAffected,
		"data":  documents,
	})
}

func (documentController) CreateDocument(ctx *gin.Context) {
	var document db.Document
	err := ctx.BindJSON(&document)
	if err != nil {
		return
	}

	result := db.DB.Create(&document)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	ctx.JSON(201, document)
}

func (documentController) GetDocument(ctx *gin.Context) {
	var document db.Document
	id := ctx.Param("id")

	result := db.DB.First(&document, id)
	switch result.Error {
	case nil:
		ctx.JSON(200, document)
	case gorm.ErrRecordNotFound:
		ctx.JSON(404, gin.H{
			"error": "Document not found",
		})
	default:
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
	}
}

func (documentController) GetDocumentData(ctx *gin.Context) {
	var document db.Document
	id := ctx.Param("id")

	result := db.DB.First(&document, id)
	switch result.Error {
	case nil:
		ctx.JSON(200, document)
	case gorm.ErrRecordNotFound:
		ctx.JSON(404, gin.H{
			"error": "Document not found",
		})
	default:
		ctx.JSON(500, gin.H{
			"error": result.Error,
		})
	}
}
