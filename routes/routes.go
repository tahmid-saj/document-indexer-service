package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// indexer routes
	server.POST("/inverted-index/:term/retrieve", getInvertedIndex) // ReadInvertedIndex
	server.POST("/inverted-index", updateInvertedIndexMappings) // UpdateInvertedIndexMappingsFromDocuments
	server.PUT("/inverted-index", updateInvertedIndex) // UpdateInvertedIndex
	server.POST("/inverted-index/:term/delete", deleteInvertedIndex) // DeleteInvertedIndex
}