package rabbit

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"time"
	"vatansoft-sms-service/pkg/constants"
)

type Client interface {
	ConnectToBroker(connectionString string) error
	PublishOnQueue(msg []byte, queueName, eventType string) error
	Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error
	SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error
	Close()

	Sync
}

type Sync interface {
	Consume(ctx context.Context, queue string, handler ConsumerGroupHandler) error
}

type MessagingClient struct {
	logger *logrus.Logger
	conn   *amqp.Connection
}

func NewMessagingClient(l *logrus.Logger) Client {
	return &MessagingClient{
		logger: l,
	}
}

func (m *MessagingClient) ConnectToBroker(connectionString string) error {
	if connectionString == "" {
		m.logger.Fatal("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
		return errors.New("empty connection string given")
	}

	var err error
	m.conn, err = amqp.Dial(brokerURL(connectionString))
	if err != nil {
		m.logger.Fatal("Failed to connect to AMQP compatible broker at: " + connectionString)
		return err
	}

	return nil
}

func (m *MessagingClient) PublishOnQueue(body []byte, queueName, eventType string) error {
	if m.conn == nil {
		m.logger.Panic("Tried to send message before connection was initialized. Don't do that.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	defer ch.Close()

	// Declare a queue that will be created if not exists with some args
	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	// Publishes a message onto the queue.
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Type:        eventType,
			Timestamp:   time.Now(),
			Body:        body,
		})
	m.logger.WithTime(time.Now()).Printf("A message was sent to queue %v: %v", queueName, body)

	return err
}

func (m *MessagingClient) Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	failOnError(m.logger, err, "Failed to open a channel")
	// defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	failOnError(m.logger, err, "Failed to register an Exchange")

	log.Printf("declared Exchange, declaring Queue (%s)", "")
	queue, err := ch.QueueDeclare(
		"",    // name of the queue
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(m.logger, err, "Failed to register an Queue")

	log.Printf("declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
		queue.Messages, queue.Consumers, exchangeName)

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("queue Bind: %s", err.Error())
	}

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(m.logger, err, "Failed to register a consumer")

	go consumeLoop(msgs, handlerFunc)
	return nil
}

func (m *MessagingClient) SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	failOnError(m.logger, err, "Failed to open a channel")

	m.logger.Printf("Declaring Queue (%s)", queueName)
	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	failOnError(m.logger, err, "Failed to register an Queue")

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(m.logger, err, "Failed to register a consumer")

	go consumeLoop(msgs, handlerFunc)
	return nil
}

func (m *MessagingClient) Consume(ctx context.Context, queue string, handler ConsumerGroupHandler) error {
	for {
		ch, err := m.conn.Channel()
		if err != nil {
			return err
		}

		deliveries, err := ch.Consume(
			queue,
			constants.AppConsumerName,
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
		if err != nil {
			return err
		}

		handler.ConsumeClaim(ctx, deliveries)
	}
}

func (m *MessagingClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}

func brokerURL(cString string) string {
	return fmt.Sprintf("%s/", cString)
}

func failOnError(l *logrus.Logger, err error, msg string) {
	if err != nil {
		l.Printf("%s: %s", msg, err)
		l.Panic(fmt.Sprintf("%s: %s", msg, err))
		return
	}
}
