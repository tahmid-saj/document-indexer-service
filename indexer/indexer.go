package indexer

import (
	"document-indexer-service/dynamodb"
	"document-indexer-service/object"
	"strings"
)

func UpdateInvertedIndexMappingsFromDocuments(tableName, bucketName string) (*dynamodb.InvertedIndexMappings, error) {
	s3Client := InitS3()
	
	generatedInvertedIndexMappings := make(dynamodb.InvertedIndexMappings)

	objectKeys, err := object.ListObjects(s3Client, bucketName)
	if err != nil {
		return nil, err
	}

	for _, objectKey := range objectKeys {
		document, err := ParseDocument(bucketName, *objectKey.Key)
		if err != nil {
			return nil, err
		}

		err = UpdateInvertedIndexMappingsFromDocument(*objectKey.Key, document, generatedInvertedIndexMappings)
		if err != nil {
			return nil, err
		}
	}

	// add the inverted index mapping to dynamodb if they do not exist
	// otherwise update the inverted index mapping
	for term, invertedIndex := range generatedInvertedIndexMappings {
		updatedInvertedIndex := dynamodb.InvertedIndex{
			Term: string(term),
			DocumentTermMatrix: dynamodb.DocumentTermMatrix{
				DocumentIDs: invertedIndex.DocumentIDs,
				DocumentTermFrequencies: invertedIndex.DocumentTermFrequencies,
				DocumentTermLocations: invertedIndex.DocumentTermLocations,
			},
		}

		_, err := dynamodb.ReadItem(string(term), tableName)
		
		// the term doesn't exist, so add the inverted index mapping to dynamodb
		if err != nil {
			_, err := dynamodb.AddItem(updatedInvertedIndex, tableName)
			if err != nil {
				return nil, err
			}
		} else {
			_, err := dynamodb.UpdateItem(updatedInvertedIndex, tableName)
			if err != nil {
				return nil, err
			}
		}
	}

	return &generatedInvertedIndexMappings, nil
}

func UpdateInvertedIndexMappingsFromDocument(objectKey string, document *object.Document, generatedInvertedIndexMappings dynamodb.InvertedIndexMappings) error {

	// traverse through the title
	titleTerms := strings.Split(document.Title, " ")
	err := parseTextUpdateDocumentTermMatrix(titleTerms, objectKey, generatedInvertedIndexMappings)
	if err != nil {
		return err
	}

	// traverse through the content
	contentTerms := strings.Split(document.Content, " ")
	err = parseTextUpdateDocumentTermMatrix(contentTerms, objectKey, generatedInvertedIndexMappings)
	if err != nil {
		return err
	}

	return nil
}


// Manual operations

func ReadInvertedIndex(term, tableName string) (*dynamodb.InvertedIndex, error) {
	resInvertedIndex, err := dynamodb.ReadItem(term, tableName)
	if err != nil {
		return nil, err
	}

	return resInvertedIndex, nil
}

func UpdateInvertedIndex(updatedInvertedIndex dynamodb.InvertedIndex, tableName string) (bool, error) {
	_, err := dynamodb.UpdateItem(updatedInvertedIndex, tableName)
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteInvertedIndex(deleteTerm, tableName string) (bool, error) {
	_, err := dynamodb.DeleteItem(deleteTerm, tableName)
	if err != nil {
		return false, err
	}

	return true, nil
}

// helper functions

func parseTextUpdateDocumentTermMatrix(terms []string, objectKey string, generatedInvertedIndexMappings dynamodb.InvertedIndexMappings) error {
	for termIndex, term := range terms {
		documentTermMatrix, exists := generatedInvertedIndexMappings[dynamodb.Term(term)];
		
		// add a new document term matrix if the term doesn't exist
		if !exists {
			generatedInvertedIndexMappings[dynamodb.Term(term)] = dynamodb.DocumentTermMatrix{
				DocumentIDs: []string{objectKey},
				DocumentTermFrequencies: []int{1},
				DocumentTermLocations: [][]int{[]int{termIndex}},
			}
		}

		// update the document term matrix if the term exists
		if exists {
			updatedDocumentIDs := documentTermMatrix.DocumentIDs
			updatedDocumentTermFrequencies := documentTermMatrix.DocumentTermFrequencies
			updatedDocumentTermLocations := documentTermMatrix.DocumentTermLocations

			updatedDocumentIDs, documentIndex, termExistsInDocument, err := updateDocumentIDInDocumentTermMatrix(objectKey, updatedDocumentIDs)
			if err != nil {
				return err
			}

			if !termExistsInDocument {
				updatedDocumentTermFrequencies = append(updatedDocumentTermFrequencies, 0)
				updatedDocumentTermLocations = append(updatedDocumentTermLocations, []int{})
			}

			updatedDocumentTermFrequencies, err = updateDocumentTermFrequencyInDocumentTermMatrix(documentIndex, updatedDocumentTermFrequencies)
			if err != nil {
				return err
			}

			updatedDocumentTermLocations, err = updateDocumentTermLocationInDocumentTermMatrix(documentIndex, termIndex, updatedDocumentTermLocations)
			if err != nil {
				return err
			}

			generatedInvertedIndexMappings[dynamodb.Term(term)] = dynamodb.DocumentTermMatrix{
				DocumentIDs: updatedDocumentIDs,
				DocumentTermFrequencies: updatedDocumentTermFrequencies,
				DocumentTermLocations: updatedDocumentTermLocations,
			}
		}
	}

	return nil
}

func updateDocumentIDInDocumentTermMatrix(documentID string, documentIDs []string) ([]string, int, bool, error){
	for docIndex, docID := range documentIDs {
		if docID == documentID {
			return documentIDs, docIndex, true, nil
		}
	}

	documentIDs = append(documentIDs, documentID)
	return documentIDs, len(documentIDs) - 1, false, nil
}

func updateDocumentTermFrequencyInDocumentTermMatrix(documentIndex int, documentTermFrequencies []int) ([]int, error) {
	documentTermFrequencies[documentIndex] += 1
	return documentTermFrequencies, nil
}

func updateDocumentTermLocationInDocumentTermMatrix(documentIndex int, termLocation int, documentTermLocations [][]int) ([][]int, error) {
	documentTermLocations[documentIndex] = append(documentTermLocations[documentIndex], termLocation)
	return documentTermLocations, nil
}