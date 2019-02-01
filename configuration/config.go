package configuration

import (
	"git.workshop21.ch/go/abraxas/configuration/yaml"
)

type Config struct {
	MinioConfig  *MinioConfig
	BucketConfig *BucketConfig
	FileConfig   *FileConfig
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type BucketConfig struct {
	BucketName string
	Location   string
}

type FileConfig struct {
	ObjectName  string
	FilePath    string
	ContentType string
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	err := yaml.ReadConfig(config,
		"./configuration/config.yaml",
	)
	return config, err
}
