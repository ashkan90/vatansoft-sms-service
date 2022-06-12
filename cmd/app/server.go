package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	di "vatansoft-sms-service"
	"vatansoft-sms-service/internal/app"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/rabbit"
	"vatansoft-sms-service/pkg/utils"
	"vatansoft-sms-service/pkg/validation"
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

	sv.initCommonMiddlewares(fApp)

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

func (s *server) initCommonMiddlewares(app *fiber.App) {
	validator := validation.InitValidator()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(utils.Validator, validator)
		return c.Next()
	})
}
