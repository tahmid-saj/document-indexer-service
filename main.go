package main

import (

)

func main() {
	// S3 ops
	// s3Client, err := conf.ConfigureS3()
	// if err != nil {
	// 	return
	// }

	// ListBuckets
	// buckets, err := bucket.ListBuckets(s3Client)
	// if err != nil {
	// 	return
	// }
	// PrintBuckets
	// bucket.PrintBuckets(buckets)

	// BucketExists
	// _, err = bucket.BucketExists(s3Client, "document-indexer-service-test-documents")
	// if err != nil {
	// 	return
	// }

	// ListObjectsV2
	// objects, err := object.ListObjects(s3Client, "document-indexer-service-test-documents")
	// if err != nil {
	// 	return
	// }
	// PrintObjects
	// object.PrintObjects(objects)

	// DownloadObject
	// err = object.DownloadObject(s3Client, "document-indexer-service-test-documents", "lorem_ipsum_1.txt", "downloaded_obj.txt")
	// if err != nil {
	// 	return
	// }

	// ReadObject
	// resObject, err := object.ReadObject(s3Client, "document-indexer-service-test-documents", "lorem_ipsum_1.json")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(resObject.Content)

	// ListObjectVersions
	// objectVersions, err := object.ListObjectVersions(s3Client, "document-indexer-service-test-documents")
	// if err != nil {
	// 	return
	// }
	// object.PrintObjectVersions(objectVersions)

	// UploadObject
	// err = object.UploadObject(s3Client, "document-indexer-service-test-documents", "lorem_ipsum_5.txt", "downloaded_obj.txt")
	// if err != nil {
	// 	return
	// }

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

	// indexer ops
	// invertedIndexMappings, err := indexer.UpdateInvertedIndexMappingsFromDocuments("document-indexer-index-mapping", "document-indexer-service-test-documents")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(invertedIndexMappings)
}