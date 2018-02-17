package config

import (
	"fmt"

	"github.com/spf13/viper"
)

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
