package minio

import (
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
