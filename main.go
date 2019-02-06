package main

import (
	"log"

	config "github.com/yabslabs/task/configuration"
	"github.com/yabslabs/task/storage/minio"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	storage, err := minio.NewClient(config.MinioConfig)
	if err != nil {
		log.Fatalf("Could not create storage client: %v", err)
	}
	err = storage.CreateBucketIfNotExisting(config)
	if err == nil {
		storage.UploadFileToBucket(config)
		storage.DownloadFileFromBucket(config)
	}
}
