package rabbit

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ConsumerInstance interface {
	Handler() ConsumerGroupHandler
	Consume(ctx context.Context, queue string)
}

type consumerInstance struct {
	logger  *logrus.Logger
	client  Client
	handler ConsumerGroupHandler
}

func NewConsumerInstance(l *logrus.Logger, c Client, h ConsumerGroupHandler) ConsumerInstance {
	return &consumerInstance{
		logger:  l,
		client:  c,
		handler: h,
	}
}

func (k *consumerInstance) Handler() ConsumerGroupHandler {
	return k.handler
}

func (k *consumerInstance) Consume(ctx context.Context, queue string) {
	for {
		if err := k.client.Consume(ctx, queue, k.handler); err != nil {
			k.logger.Fatalf("consume: %v", err)
		}

		if ctx.Err() != nil {
			return
		}

		k.handler.Ready()
	}
}
