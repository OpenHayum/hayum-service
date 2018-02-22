package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// CollectionUser collection name for users
	CollectionUser = "users"

	// CollectionArtist collection name for artists
	CollectionArtist = "artists"

	// CollectionAudio collection name for audios
	CollectionAudio = "audios"

	// CollectionMessage collection name for messages
	CollectionMessage = "messages"

	// CollectionS3Document collection name for S3 Documents
	CollectionS3Document = "s3Documents"
)

// LoadConfig loads any yaml config file
func LoadConfig(configPaths ...string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	return v, nil
}
