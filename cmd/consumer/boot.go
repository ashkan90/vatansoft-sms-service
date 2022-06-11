package main

import (
	"github.com/spf13/viper"
	"vatansoft-sms-service/configs/consumer"
)

func initConfig() (*consumer.Config, error) {
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var appConf consumer.Config

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&appConf)
	if err != nil {
		return nil, err
	}

	return &appConf, nil
}
