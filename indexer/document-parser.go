package indexer

import "document-indexer-service/object"

func ParseDocument(bucketName, objectKey string) (*object.Document, error) {
	s3Client := InitS3()

	document, err := object.ReadObject(s3Client, bucketName, objectKey)
	if err != nil {
		return nil, err
	}

	return document, nil
}