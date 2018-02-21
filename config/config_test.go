package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		configPaths []string
	}
	tests := []struct {
		name    string
		args    args
		want    *viper.Viper
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.configPaths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
