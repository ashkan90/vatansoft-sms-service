package main

import (
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/httpclient"
	"vatansoft-sms-service/pkg/mobilisimclient"
)

func boot(l *logrus.Logger) (*server, error) {
	var httpClient = httpclient.NewHTTPClient()
	var mobilisimClient = mobilisimclient.NewClient(initMobilisimClientConfig(), httpClient)

	return &server{
		logger:          l,
		mobilisimClient: mobilisimClient,
	}, nil
}

func initMobilisimClientConfig() mobilisimclient.Config {
	return mobilisimclient.Config{}
}
