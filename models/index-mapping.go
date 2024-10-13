package models

import (
	"document-indexer-service/dynamodb"
	"document-indexer-service/indexer"
)

type ReadInvertedIndexInput struct {
	TableName string `json:"tableName"`
}

type UpdateInvertedIndexMappingsFromDocumentsInput struct {
	TableName string `json:"tableName"`
	BucketName string `json:"bucketName"`
}

type UpdateInvertedIndexInput struct {
	TableName          string `json:"tableName"`
	InvertedIndex dynamodb.InvertedIndex `json:"invertedIndex"`
}

type DeleteInvertedIndexInput struct {
	TableName          string `json:"tableName"`
}

type Response struct {
	Ok bool
	Response interface{}
}

func ReadInvertedIndex(term string, readInvertedIndexInput ReadInvertedIndexInput) (*Response, error) {
	resInvertedIndex, err := indexer.ReadInvertedIndex(term, readInvertedIndexInput.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: resInvertedIndex,
	}, nil
}

func UpdateInvertedIndexMappingsFromDocuments(updateInvertedIndexMappingsFromDocumentsInput UpdateInvertedIndexMappingsFromDocumentsInput) (*Response, error) {
	updatedInvertedIndexMappings, err := indexer.UpdateInvertedIndexMappingsFromDocuments(
		updateInvertedIndexMappingsFromDocumentsInput.TableName, 
		updateInvertedIndexMappingsFromDocumentsInput.BucketName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: updatedInvertedIndexMappings,
	}, nil
}

func UpdateInvertedIndex(updateInvertedIndexInput UpdateInvertedIndexInput) (*Response, error) {
	resUpdatedInvertedIndex, err := indexer.UpdateInvertedIndex(updateInvertedIndexInput.InvertedIndex, updateInvertedIndexInput.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: resUpdatedInvertedIndex,
	}, nil
}

func DeleteInvertedIndex(term string, deleteInvertedIndex DeleteInvertedIndexInput) (*Response, error) {
	resDeletedInvertedIndex, err := indexer.DeleteInvertedIndex(term, deleteInvertedIndex.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: resDeletedInvertedIndex,
	}, nil
}