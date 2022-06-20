package rabbit

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ConsumerInstance interface {
	Handler() ConsumerGroupHandler
	Consume(ctx context.Context, topics []string)
}

type consumerInstance struct {
	logger  *logrus.Logger
	client  *MessagingClient
	handler ConsumerGroupHandler
}

func NewConsumerInstance(l *logrus.Logger, c *MessagingClient, h ConsumerGroupHandler) ConsumerInstance {
	return &consumerInstance{
		logger:  l,
		client:  c,
		handler: h,
	}
}

func (k *consumerInstance) Handler() ConsumerGroupHandler {
	return k.handler
}

func (k *consumerInstance) Consume(ctx context.Context, topics []string) {
	k.client.SubscribeToQueue()
	for {

		if err := k.client.Consume(ctx, topics, k.handler); err != nil {
			k.logger.Fatalf("consume: %v", err)
		}

		if ctx.Err() != nil {
			return
		}

		k.handler.Ready()
	}
}
