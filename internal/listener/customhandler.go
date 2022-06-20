package listener

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
	"vatansoft-sms-service/pkg/eventmanager"
	"vatansoft-sms-service/pkg/rabbit"
)

type customHandler struct {
	logger       *logrus.Logger
	eventManager eventmanager.EventManager
}

func NewCustomHandler(l *logrus.Logger, em eventmanager.EventManager) rabbit.CustomHandler {
	return &customHandler{
		logger:       l,
		eventManager: em,
	}
}

func (c *customHandler) Do(ctx context.Context, delivery amqp.Delivery) error {
	attr, err := c.eventManager.Handle(ctx, delivery)
	if err != nil {
		return fmt.Errorf("consume: %v", err)
	}

	if attr != nil {
		c.logger.WithContext(ctx).WithField("event", logrus.Fields{
			"key":         "input.Key",
			"type":        attr.Type,
			"createdAt":   attr.CreatedAt,
			"deliveredAt": time.Now(),
		}).Info("message delivered")
	}

	return nil
}
