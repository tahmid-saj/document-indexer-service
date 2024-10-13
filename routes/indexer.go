package routes

import (
	"document-indexer-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getInvertedIndex(context *gin.Context) {
	term := context.Param("term")
	var readInvertedIndexInput models.ReadInvertedIndexInput

	err := context.ShouldBindJSON(&readInvertedIndexInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.ReadInvertedIndex(term, readInvertedIndexInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch inverted index mapping"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func updateInvertedIndexMappings(context *gin.Context) {
	var updateInvertedIndexMappingsFromDocumentsInput models.UpdateInvertedIndexMappingsFromDocumentsInput

	err := context.ShouldBindJSON(&updateInvertedIndexMappingsFromDocumentsInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.UpdateInvertedIndexMappingsFromDocuments(updateInvertedIndexMappingsFromDocumentsInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update inverted index mappings"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func updateInvertedIndex(context *gin.Context) {
	var updateInvertedIndexInput models.UpdateInvertedIndexInput

	err := context.ShouldBindJSON(&updateInvertedIndexInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.UpdateInvertedIndex(updateInvertedIndexInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update inverted index"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func deleteInvertedIndex(context *gin.Context) {
	term := context.Param("term")

	var deleteInvertedIndexInput models.DeleteInvertedIndexInput

	err := context.ShouldBindJSON(&deleteInvertedIndexInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.DeleteInvertedIndex(term, deleteInvertedIndexInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete inverted index"})
		return
	}

	context.JSON(http.StatusOK, res)
}