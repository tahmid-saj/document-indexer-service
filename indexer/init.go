package indexer

import (
	"document-indexer-service/conf"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitS3() *s3.Client {
	var s3Client, err = conf.ConfigureS3()
	if err != nil {
		os.Exit(1)
	}

	return s3Client
}