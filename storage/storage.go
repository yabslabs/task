package storage

import "github.com/yabslabs/task/configuration"

type Storage interface {
	CreateBucketIfNotExisting(*configuration.StorageConfig) error
	UploadFileToBucket(*configuration.StorageConfig) error
	DownloadFileFromBucket(*configuration.StorageConfig) error
}
