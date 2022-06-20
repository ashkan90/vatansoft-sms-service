package rabbit

import "github.com/streadway/amqp"

type MessageAttribute struct {
	Type      string
	CreatedAt int64
}

func GetEventAttributes(delivery *amqp.Delivery) *MessageAttribute {
	return &MessageAttribute{
		Type:      delivery.Type,
		CreatedAt: delivery.Timestamp.Unix(),
	}
}
