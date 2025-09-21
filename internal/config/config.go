package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		Storage  StorageConfig
	}

	ServerConfig struct {
		Port int
		Host string
	}
	DatabaseConfig struct {
		Path string
	}

	StorageConfig struct {
		Endpoint   string
		AccessKey  string
		SecretKey  string
		BucketName string
	}
)

func Init() (*Config, error) {
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfigFile() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
