package handler

import (
	"context"
	"vatansoft-sms-service/internal/listener/consumer"
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/eventmanager"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/utils"
)

type oneToNConsumerHandler struct {
	service consumer.Service
	event   *event.OneToNEvent
}

func NewOneToNConsumerHandler(service consumer.Service, event *event.OneToNEvent) eventmanager.EventHandler {
	return &oneToNConsumerHandler{
		service: service,
		event:   event,
	}
}

func (o *oneToNConsumerHandler) Handle(ctx context.Context) error {
	return o.service.OneToN(ctx, o.preparePayload())
}

func (o *oneToNConsumerHandler) preparePayload() model.RequestOneToN {
	var dto = model.RequestOneToN{
		Messages: []model.OneToNMessage{
			{
				From:             o.event.EventData.Sender,
				Text:             utils.RecomposeMessage(o.event.EventData.Message, o.event.EventData.MessageType),
				CallbackData:     "cb",
				LanguageEncoding: utils.GetMessageType(o.event.EventData.MessageType),
				ValidityPeriod:   2880,
			},
		},
	}

	for _, number := range o.event.EventData.Numbers {
		dto.Messages[0].Destinations = append(dto.Messages[0].Destinations, model.MessageDestination{
			To: utils.CleanupPhone(number),
		})
	}

	return dto
}
