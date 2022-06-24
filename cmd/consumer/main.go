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
	"vatansoft-sms-service/pkg/event/schema"
	"vatansoft-sms-service/pkg/rabbit"
)

func main() {
	conf, err := initConfig()
	if err != nil {
		panic("<consumer> " + err.Error())
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	l := initLogger()

	l.Println("Starting " + constants.AppConsumerName + "...")

	app, err := boot(l, conf.Consumer)

	sv := initServer(app)

	go app.mqInstance.Consume(context.Background(), schema.MobilisimQueueName)

	<-app.mqInstance.Handler().Status()

	go log.Fatal(sv.Listen(":1598"))

	graceful(sv, app.mqInstance)
}

func graceful(a *fiber.App, mq rabbit.ConsumerInstance) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	_, cancel := context.WithTimeout(context.Background(), constants.AppGracefulTimeout*time.Second)
	defer cancel()
	defer mq.Close()

	if err := a.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
