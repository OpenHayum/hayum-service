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

	// CollectionSession collection name for Session
	CollectionSession = "sessions"

	// ExternalConfigFilePath external config file path
	ExternalConfigFilePath = "/opt/conf/hayum"
)

// App contains all config globally
var App *viper.Viper

// Detail stores the configuration file details
type Detail struct {
	Path string
	Name string
	Type string
}

// NewDetail creates a new Detail
func NewDetail(path string, name string) *Detail {
	return &Detail{path, name, "yaml"}
}

// LoadConfig loads any yaml config file
func LoadConfig(details ...*Detail) (*viper.Viper, error) {
	v := viper.New()

	for _, configDetail := range details {
		v.AddConfigPath(configDetail.Path)
		v.SetConfigName(configDetail.Name)
		v.SetConfigType(configDetail.Type)

		if err := v.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
		}
	}

	App = v

	return v, nil
}
