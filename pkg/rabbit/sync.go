package rabbit

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type syncHandler struct {
	logger          *logrus.Logger
	consumerManager ConsumerManager

	ready chan bool
}

func NewSyncHandler(l *logrus.Logger, cm ConsumerManager) ConsumerGroupHandler {
	return &syncHandler{
		logger:          l,
		consumerManager: cm,
		ready:           make(chan bool),
	}
}

func (s *syncHandler) ConsumeClaim(ctx context.Context, queue <-chan amqp.Delivery) {
	errCh := make(chan error)
	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Context done...")
			return
		case delivery := <-queue:
			s.logger.Info("Got a delivery... " + string(delivery.Body))
			go func() {
				errCh <- s.consumerManager.Process(ctx, delivery)
			}()

			select {
			case err := <-errCh:
				if err != nil {
					delivery.Nack(false, true)
				}
			}
		}
	}
}

func (s *syncHandler) Ready() {
	s.ready = make(chan bool)
}

func (s *syncHandler) Status() chan bool {
	return s.ready
}
