package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	di "vatansoft-sms-service"
	"vatansoft-sms-service/internal/app"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/rabbit"
)

type server struct {
	logger          *logrus.Logger
	mobilisimClient mobilisimclient.Client
	mqProducer      rabbit.Client
}

func initServer(sv *server) *fiber.App {
	fApp := fiber.New(fiber.Config{
		BodyLimit: constants.AppRequestBodyLimit,
	})

	route := di.InitAll(
		sv.logger,
		sv.mobilisimClient,
		sv.mqProducer,
	)
	route.SetupRoutes(&app.RouteCtx{
		App: fApp,
	})

	return fApp
}

func initLogger() *logrus.Logger {
	return logrus.New()
}
