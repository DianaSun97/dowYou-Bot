package config

import "github.com/spf13/viper"

type Config struct {
	Token string `json:"token"`
	Port  string `json:"port"`
}

var cfg Config

func ParseConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}

func GetConfig() Config {
	return cfg
}
