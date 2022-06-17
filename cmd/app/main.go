package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"time"
	"vatansoft-sms-service/pkg/constants"
)

func main() {
	logger := initLogger()

	conf, cErr := initConfig()
	if cErr != nil {
		logger.Error(cErr)
		return
	}

	app, err := boot(logger, conf.Application)
	if err != nil {
		logger.Fatalf("Something went wrong while utilizing the server. %v", err)
	}
	sv := initServer(app)

	go log.Fatal(sv.Listen(":8080"))

	graceful(logger, sv)
}

func graceful(l *logrus.Logger, a *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	_, cancel := context.WithTimeout(context.Background(), constants.AppGracefulTimeout*time.Second)
	defer cancel()
	if err := a.Shutdown(); err != nil {
		l.Fatal(err)
	}
}
