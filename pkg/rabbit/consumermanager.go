package rabbit

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ConsumerGroupHandler interface {
	Ready()
	Status() chan bool
}

type CustomHandler interface {
	Do(ctx context.Context, delivery amqp.Delivery) error
}

type IConsumerManager interface {
	Process(ctx context.Context, delivery amqp.Delivery) error
}

type ConsumerSessionMessage struct {
	Session sarama.ConsumerGroupSession
	Message *sarama.ConsumerMessage
}

type consumerManager struct {
	logger        *logrus.Logger
	customHandler CustomHandler
	//newRelicInstance nrclient.INewRelicInstance
}

func NewConsumerManager(l *logrus.Logger, ch CustomHandler /*ni nrclient.INewRelicInstance*/) IConsumerManager {
	return &consumerManager{
		logger:        l,
		customHandler: ch,
		//newRelicInstance: ni,
	}
}

func (cm *consumerManager) Process(ctx context.Context, delivery amqp.Delivery) error {
	txn := cm.newRelicInstance.Application().StartTransaction(fmt.Sprintf("rentId:%s", string(msg.Key)))
	txn.AddAttribute("event.topic", msg.Topic)
	txn.AddAttribute("event.partition", msg.Partition)
	txn.AddAttribute("event.offset", msg.Offset)

	defer txn.End()

	ctx = newrelic.NewContext(ctx, txn)
	if err := cm.customHandler.Do(ctx, delivery); err != nil {
		cm.logger.WithField("event", cm.prepareLogFields(delivery)).WithError(err).Error("processing error")
	}

	return nil
}

func (cm *consumerManager) prepareLogFields(delivery amqp.Delivery) logrus.Fields {
	return logrus.Fields{
		"topic":     msg.Topic,
		"key":       string(msg.Key),
		"partition": msg.Partition,
		"offset":    msg.Offset,
		"body":      string(msg.Value),
	}
}
