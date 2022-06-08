package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	di "vatansoft-sms-service"
	"vatansoft-sms-service/internal/app"
	"vatansoft-sms-service/pkg/mobilisimclient"
)

const (
	bodyLimit = 1024 << 10 << 10
)

type server struct {
	logger          *logrus.Logger
	mobilisimClient mobilisimclient.Client
}

func initServer(sv *server) *fiber.App {
	fApp := fiber.New(fiber.Config{
		BodyLimit: bodyLimit,
	})

	route := di.InitAll(
		sv.logger,
		sv.mobilisimClient,
	)
	route.SetupRoutes(&app.RouteCtx{
		App: fApp,
	})

	return fApp
}

func initLogger() *logrus.Logger {
	return logrus.New()
}
