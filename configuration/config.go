package configuration

import (
	"git.workshop21.ch/go/abraxas/configuration/yaml"
)

type StorageConfig struct {
	MinioConfig  *MinioConfig
	GoogleConfig *GoogleConfig
	FileConfig   *FileConfig
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName string
	Location   string
}

type GoogleConfig stuct {
	ProjectID string
	BucketName string
}

type FileConfig struct {
	ObjectName  string
	FilePath    string
	ContentType string
}

func ReadConfig() (*StorageConfig, error) {
	config := &StorageConfig{}
	err := yaml.ReadConfig(config,
		"./configuration/config.yaml",
	)
	return config, err
}
