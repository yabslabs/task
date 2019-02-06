package storage

import (
	"time"

	"github.com/yabslabs/task/configuration"
)

type Storage interface {
	ListBuckets(*configuration.StorageConfig) ([]BucketInfo, error)
	CreateBucketIfNotExisting(*configuration.StorageConfig) error
	ListBucketObjects(*configuration.StorageConfig) ([]ObjectInfo, error)
	UploadObjectToBucket(*configuration.StorageConfig) error
	GetObjectFromBucket(*configuration.StorageConfig) error
}

type BucketInfo struct {
	Name         string
	CreationDate time.Time
}

type ObjectInfo struct {
	Name         string
	LastModified time.Time
	ContentType  string
	Size         int64
}
