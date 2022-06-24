package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/rabbit"
)

type server struct {
	logger          *logrus.Logger
	mobilisimClient mobilisimclient.Client
	mqInstance      rabbit.ConsumerInstance
}

func initServer(sv *server) *fiber.App {
	fApp := fiber.New(fiber.Config{
		BodyLimit: constants.AppRequestBodyLimit,
	})

	return fApp
}

func initLogger() *logrus.Logger {
	return logrus.New()
}
