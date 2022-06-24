package rabbit

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
	"vatansoft-sms-service/pkg/constants"
)

type Client interface {
	ConnectToBroker(connectionString string) error
	PublishOnQueue(msg []byte, queueName, eventType string) error
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

	return err
}

func (m *MessagingClient) Consume(ctx context.Context, queue string, handler ConsumerGroupHandler) error {
	ch, chErr := m.conn.Channel()
	if chErr != nil {
		return chErr
	}

	queueDeclaration, qErr := ch.QueueDeclare(
		queue, // name of the queue
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if qErr != nil {
		return qErr
	}

	for {
		deliveries, cErr := ch.Consume(
			queueDeclaration.Name,
			constants.AppConsumerName,
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
		if cErr != nil {
			return cErr
		}

		handler.ConsumeClaim(ctx, deliveries)
	}
}

func (m *MessagingClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func brokerURL(cString string) string {
	return fmt.Sprintf("%s/", cString)
}
