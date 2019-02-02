package main

import (
	"log"

	config "github.com/yabslabs/task/configuration"
	"github.com/yabslabs/task/storage"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	storage := storage.NewClient(config.MinioConfig)
	err = storage.CreateBucketIfNotExisting(config.BucketConfig)
	if err == nil {
		storage.UploadFileToBucket(config.BucketConfig, config.FileConfig)
		storage.DownloadFileFromBucket(config.BucketConfig, config.FileConfig)
	}
}
