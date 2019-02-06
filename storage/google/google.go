package google

// import (
//         "context"
//         "fmt"
//         "log"

//         "cloud.google.com/go/storage"
// )

// type GoogleStorage struct {
// 	Client *storage.Client
// 	Bucket *storage.BucketHandle
// }

// func NewClient(storageConfig *configuration.GoogleConfig) (Storage, error) {
// 	ctx := context.Background()

// 	googleClient, err := storage.NewClient(ctx)
// 	if err != nil {
// 			log.Printf("Failed to create client: %v", err)
// 			return err
// 	}

// 	return &GoogleStorage{
// 		Client: googleClient,
// 	}
// }

// func (s *GoogleStorage) CreateBucketIfNotExisting(config *configuration.StorageConfig) error {
// 	s.Bucket = s.Client.Bucket(config.GoogleConfig.BucketName)

// 	if err := s.Bucket.Create(ctx, config.GoogleConfig.ProjectID, nil); err != nil {
// 		log.Printf("Failed to create bucket: %v", err)
// 		return err
// 	}

// 	log.Printf("Bucket %v created.\n", config.GoogleConfig.BucketName)
// 	return nil
// }

// func (s *GoogleStorage) UploadFileToBucket(config *configuration.StorageConfig) error {
// 	f, err := os.Open(config.FileConfig.FilePath)
// 	if err != nil {
// 		log.Printf("Failed to open file: %v", err)
// 		return err
// 	}
// 	defer f.Close()

// 	wc := s.Client.Bucket(s.Bucket).Object(object).NewWriter(context.Background())
// 	if _, err = io.Copy(wc, f); err != nil {
// 		log.Printf("Failed to copy file: %v", err)
// 		return err
// 	}
// 	if err := wc.Close(); err != nil {
// 		log.Printf("Failed to close writer: %v", err)
// 		return err
// 	}
// 	return nil
// }

// func (s *GoogleStorage) DownloadFileFromBucket(config *configuration.StorageConfig) error {
// 	// Upload the zip file with FPutObject
// 	// err := s.Client.FGetObject(bucketConfig.BucketName, fileConfig.ObjectName, "storage/tmp/downloads/"+fileConfig.ObjectName, minio.GetObjectOptions{})
// 	// if err != nil {
// 	// 	log.Printf("Could not get file %v", err)
// 	// 	return err
// 	// }

// 	// log.Printf("Successfully donwloaded %s", fileConfig.ObjectName)
// 	return nil
// }
