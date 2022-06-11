package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/rabbit"
)

var messagingClient rabbit.Client

func main() {
	fmt.Println("Starting " + constants.AppName + "...")

	initializeMessaging(logrus.New())

	// Makes sure connection is closed when service exits.
	handleSigterm(func() {
		if messagingClient != nil {
			messagingClient.Close()
		}
	})
	//service.StartWebServer(viper.GetString("server_port"))
}

// The callback function that's invoked whenever we get a message on the "vipQueue"
func onMessage(delivery amqp.Delivery) {
	fmt.Printf("Got a message: %v\n", string(delivery.Body))
}

func initializeMessaging(l *logrus.Logger) {
	if !viper.IsSet("amqp_server_url") {
		panic("No 'broker_url' set in configuration, cannot start")
	}
	messagingClient = rabbit.NewMessagingClient(l)
	messagingClient.ConnectToBroker(viper.GetString("amqp_server_url"))

	// Call the subscribe method with queue name and callback function
	err := messagingClient.SubscribeToQueue("vip_queue", constants.AppName, onMessage)
	failOnError(err, "Could not start subscribe to vip_queue")

	err = messagingClient.Subscribe(viper.GetString("config_event_bus"), "topic", constants.AppName, func(delivery amqp.Delivery) {

	})
	failOnError(err, "Could not start subscribe to "+viper.GetString("config_event_bus")+" topic")
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
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
