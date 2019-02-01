package storage

import (
	"log"

	"github.com/minio/minio-go"
	"github.com/yabslabs/task/configuration"
)

type MinioStorage struct {
	Client *minio.Client
}

func NewClient(storageConfig *configuration.MinioConfig) Storage {
	minioClient, err := minio.New(storageConfig.Endpoint, storageConfig.AccessKeyID, storageConfig.SecretAccessKey, true)
	if err != nil {
		log.Fatalln(err)
	}

	return &MinioStorage{
		Client: minioClient,
	}
}

func (s *MinioStorage) CreateBucketIfNotExisting(bucketConfig *configuration.BucketConfig) {
	err := s.Client.MakeBucket(bucketConfig.BucketName, bucketConfig.Location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := s.Client.BucketExists(bucketConfig.BucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketConfig.BucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketConfig.BucketName)
	}
}

func (s *MinioStorage) UploadFileToBucket(bucketConfig *configuration.BucketConfig, fileConfig *configuration.FileConfig) {
	// Upload the zip file with FPutObject
	n, err := s.Client.FPutObject(bucketConfig.BucketName, fileConfig.ObjectName, fileConfig.FilePath, minio.PutObjectOptions{ContentType: fileConfig.ContentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileConfig.ObjectName, n)
}
