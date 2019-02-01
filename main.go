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
	storage.CreateBucketIfNotExisting(config.BucketConfig)
	storage.UploadFileToBucket(config.BucketConfig, config.FileConfig)
}
