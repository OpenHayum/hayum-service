package config

import (
	"fmt"
	"hayum/core_apis/logger"
	"log"
	"os"

	"github.com/spf13/viper"
)

const externalConfigPath = "/opt/conf/hayum"

// Detail stores the configuration file details
type Detail struct {
	Path string
	Name string
	Type string
}

func newDetail(path string, name string) *Detail {
	return &Detail{path, name, "yaml"}
}

func newConfig(details ...*Detail) (*viper.Viper, error) {
	v := viper.New()

	for _, configDetail := range details {
		v.AddConfigPath(configDetail.Path)
		v.SetConfigName(configDetail.Name)
		v.SetConfigType(configDetail.Type)

		if err := v.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read the configuration file: %s", err)
		}
	}

	return v, nil
}

func getConfigFilename(env string) string {
	return fmt.Sprintf("config.%s", env)
}

func New() *viper.Viper {
	env := os.Getenv("GO_ENV")
	logger.Log.Info("GO_ENV:", env)

	internalConfig := newDetail("config", getConfigFilename(env))
	externalConfig := newDetail(externalConfigPath, getConfigFilename(env))
	v, err := newConfig(internalConfig, externalConfig)

	if err != nil {
		log.Fatal(err)
	}

	return v
}
