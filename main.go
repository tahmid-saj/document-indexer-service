package main

import (
	"document-indexer-service/dynamodb"
	"fmt"
)

func main() {
	// DynamoDB ops
	// CreateTable
	// createdTable, err := dynamodb.CreateTable("document-indexer-index-mapping")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(createdTable)

	// ListTables
	// tableNames, err := dynamodb.ListTables()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(tableNames)

	// AddItem
	// invertedIndex := dynamodb.InvertedIndex{
	// 	Term: "search",
	// 	DocumentTermMatrix: dynamodb.DocumentTermMatrix{
	// 		DocumentIDs: []string{"test1", "test2"},
	// 		DocumentTermFrequencies: []int{2, 4},
	// 		DocumentTermLocations: [][]int{[]int{1}, []int{2}},
	// 	},
	// }
	// addedItem, err := dynamodb.AddItem(invertedIndex, "document-indexer-index-mapping")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(addedItem)

	// ReadItem
	// itemRead, err := dynamodb.ReadItem("search", "document-indexer-index-mapping")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(itemRead)

	// UpdateItem
	// updatedInvertedIndex := dynamodb.InvertedIndex{
	// 	Term: "search",
	// 	DocumentTermMatrix: dynamodb.DocumentTermMatrix{
	// 		DocumentIDs: []string{"test", "test1"},
	// 		DocumentTermFrequencies: []int{3, 4},
	// 		DocumentTermLocations: [][]int{[]int{1}, []int{2}},
	// 	},
	// }
	// updatedItem, err := dynamodb.UpdateItem(updatedInvertedIndex, "document-indexer-index-mapping")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(updatedItem)

	// DeleteItem
	// term := "search"
	// deletedItem, err := dynamodb.DeleteItem(term, "document-indexer-index-mapping")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(deletedItem)
}