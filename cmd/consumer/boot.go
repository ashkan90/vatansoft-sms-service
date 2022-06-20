package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"vatansoft-sms-service/configs/consumer"
	"vatansoft-sms-service/internal/listener"
	consumerservice "vatansoft-sms-service/internal/listener/consumer"
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/eventmanager"
	"vatansoft-sms-service/pkg/httpclient"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/rabbit"
)

func boot(l *logrus.Logger, appConf consumer.Application) (*server, error) {
	var httpClient = httpclient.NewHTTPClient()
	var mobilisimClient = mobilisimclient.NewClient(initMobilisimClientConfig(appConf.Mobilisim), httpClient)
	var mqConsumer = rabbit.NewMessagingClient(l)

	var (
		consumerService     = consumerservice.NewMobilisimConsumerService(l, mobilisimClient, mqConsumer)
		eventHandlerFactory = listener.NewEventHandlerFactory(consumerService)
		eventManager        = eventmanager.NewEventManager(eventHandlerFactory, event.NewEventCreator())
		customHandler       = listener.NewCustomHandler(l, eventManager)
		consumerGroup       = rabbit.NewConsumerManager(l, customHandler)
	)

	// Open broker connection
	err := mqConsumer.ConnectToBroker(appConf.Rabbit.URL)
	if err != nil {
		return nil, err
	}

	return &server{
		logger:          l,
		mobilisimClient: mobilisimClient,
		mqProducer:      mqProducer,
	}, nil
}

func initMobilisimClientConfig(mobilisimConf consumer.MobilisimConfig) mobilisimclient.Config {
	return mobilisimclient.Config{
		MobilisimURL: mobilisimConf.URL,
		APIKey:       mobilisimConf.ApiKey,
	}
}

func initConfig() (*consumer.Configs, error) {
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var appConf consumer.Configs

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
