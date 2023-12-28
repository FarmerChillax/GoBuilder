package config

import "github.com/spf13/viper"

type Config struct {
	Platform   []string `json:"platform,omitempty"`
	SourcePath string   `json:"source_path,omitempty"`
}

var config *Config

func Get() *Config {
	return config
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
