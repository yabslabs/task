package storage

import "github.com/yabslabs/task/configuration"

type Storage interface {
	CreateBucketIfNotExisting(*configuration.BucketConfig)
	UploadFileToBucket(*configuration.BucketConfig, *configuration.FileConfig)
}
