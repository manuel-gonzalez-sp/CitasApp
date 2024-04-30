package internal

import (
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	CloudinaryCloudName string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey    string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret string `mapstructure:"CLOUDINARY_API_SECRET"`
}

func NewEnvConfig() (*Config, error) {
	var config Config

	viper.AutomaticEnv()

	t := reflect.TypeOf(config)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("mapstructure")
		viper.BindEnv(field)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
