package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"vatansoft-sms-service/configs/app"
	"vatansoft-sms-service/pkg/httpclient"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/rabbit"
)

func boot(l *logrus.Logger, appConf app.ApplicationConfigs) (*server, error) {
	var httpClient = httpclient.NewHTTPClient()
	var mobilisimClient = mobilisimclient.NewClient(initMobilisimClientConfig(appConf.Mobilisim), httpClient)
	var mqProducer = rabbit.NewMessagingClient(l)

	// Open broker connection
	mqProducer.ConnectToBroker(appConf.Rabbit.URL)

	return &server{
		logger:          l,
		mobilisimClient: mobilisimClient,
		mqProducer:      mqProducer,
	}, nil
}

func initMobilisimClientConfig(mobilisimConf app.MobilisimConfig) mobilisimclient.Config {
	return mobilisimclient.Config{
		MobilisimURL: mobilisimConf.URL,
		APIKey:       mobilisimConf.ApiKey,
	}
}

func initConfig() (*app.Configs, error) {
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var appConf app.Configs

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
