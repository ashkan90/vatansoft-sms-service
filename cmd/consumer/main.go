package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vatansoft-sms-service/configs/consumer"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/event/schema"
	"vatansoft-sms-service/pkg/rabbit"
)

var messagingClient rabbit.Client

func main() {
	fmt.Println("Starting " + constants.AppConsumerName + "...")

	conf, err := initConfig()
	if err != nil {
		panic("<consumer> " + err.Error())
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	l := initLogger()

	app, err := boot(l, conf.Consumer)

	sv := initServer(app)

	initializeMessaging(logrus.New(), conf.Consumer.Rabbit)

	// Makes sure connection is closed when service exits.
	handleSigterm(func() {
		if messagingClient != nil {
			messagingClient.Close()
		}
	})

	go app.mqInstance.Consume(context.Background(), schema.MobilisimQueueName)

	<-app.mqInstance.Handler().Status()

	go log.Fatal(sv.Listen(":1598"))

	graceful(sv)
}

func graceful(a *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	_, cancel := context.WithTimeout(context.Background(), constants.AppGracefulTimeout*time.Second)
	defer cancel()
	if err := a.Shutdown(); err != nil {
		log.Fatal(err)
	}
}

// The callback function that's invoked whenever we get a message on the "vipQueue"
func onMessage(delivery amqp.Delivery) {
	fmt.Printf("Got a message: %v\n", string(delivery.Body))
}

func initializeMessaging(l *logrus.Logger, mqConf consumer.RabbitConfig) {
	if mqConf.URL == "" {
		panic("No 'broker_url' set in configuration, cannot start")
	}
	messagingClient = rabbit.NewMessagingClient(l)
	messagingClient.ConnectToBroker(mqConf.URL)

	// Call the subscribe method with queue name and callback function
	err := messagingClient.SubscribeToQueue(schema.MobilisimQueueName, constants.AppConsumerName, onMessage)
	failOnError(err, "Could not start subscribe to "+schema.MobilisimQueueName)
}

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		log.Fatal(fmt.Sprintf("%s: %s", msg, err))
	}
}
