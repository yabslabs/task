package minio

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
	"github.com/yabslabs/task/configuration"
	"github.com/yabslabs/task/storage"
)

type MinioStorage struct {
	Client *minio.Client
}

func NewClient(storageConfig *configuration.MinioConfig) (storage.Storage, error) {
	minioClient, err := minio.New(storageConfig.Endpoint, storageConfig.AccessKeyID, storageConfig.SecretAccessKey, true)
	if err != nil {
		log.Printf("Could not create minio client: %v", err)
		return nil, err
	}

	return &MinioStorage{
		Client: minioClient,
	}, nil
}

func (s *MinioStorage) CreateBucketIfNotExisting(config *configuration.StorageConfig) error {
	err := s.Client.MakeBucket(config.MinioConfig.BucketName, config.MinioConfig.Location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := s.Client.BucketExists(config.MinioConfig.BucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", config.MinioConfig.BucketName)
			return nil
		} else {
			log.Printf("Error in check for existing bucket, %v", err)
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", config.MinioConfig.BucketName)
		return nil
	}
}

func (s *MinioStorage) ListBuckets(config *configuration.StorageConfig) ([]storage.BucketInfo, error) {
	buckets, err := s.Client.ListBuckets()
	if err != nil {
		log.Printf("Error reading buckets, %v", err)
		return nil, err
	}
	bucketList := make([]storage.BucketInfo, 0)
	for _, bucket := range buckets {
		bucketList = append(bucketList, storage.BucketInfo{Name: bucket.Name, CreationDate: bucket.CreationDate})
	}
	return bucketList, err
}

func (s *MinioStorage) ListBucketObjects(config *configuration.StorageConfig) ([]storage.ObjectInfo, error) {
	doneCh := make(chan struct{})

	defer close(doneCh)

	isRecursive := true
	objectList := make([]storage.ObjectInfo, 0)
	objectCh := s.Client.ListObjectsV2(config.MinioConfig.BucketName, "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			log.Printf("Could not read objects of a bucket %v", object.Err)
			return nil, object.Err
		}
		fmt.Println(object)
		objectList = append(objectList, storage.ObjectInfo{Name: object.Key, LastModified: object.LastModified, ContentType: object.ContentType, Size: object.Size})
	}
	return objectList, nil
}

func (s *MinioStorage) UploadFileToBucket(config *configuration.StorageConfig) error {
	// Upload the zip file with FPutObject
	n, err := s.Client.FPutObject(config.MinioConfig.BucketName, config.FileConfig.ObjectName, config.FileConfig.FilePath, minio.PutObjectOptions{ContentType: config.FileConfig.ContentType})
	if err != nil {
		log.Printf("Could not upload file %v", err)
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", config.FileConfig.ObjectName, n)
	return nil
}

func (s *MinioStorage) DownloadFileFromBucket(config *configuration.StorageConfig) error {
	// Upload the zip file with FPutObject
	err := s.Client.FGetObject(config.MinioConfig.BucketName, config.FileConfig.ObjectName, "storage/tmp/downloads/"+config.FileConfig.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Could not get file %v", err)
		return err
	}

	log.Printf("Successfully donwloaded %s", config.FileConfig.ObjectName)
	return nil
}
