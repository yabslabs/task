package storage

import "github.com/yabslabs/task/configuration"

type Storage interface {
	CreateBucketIfNotExisting(*configuration.BucketConfig) error
	UploadFileToBucket(*configuration.BucketConfig, *configuration.FileConfig) error
	DownloadFileFromBucket(*configuration.BucketConfig, *configuration.FileConfig) error
}
