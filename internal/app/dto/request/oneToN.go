package request

import (
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/event/schema"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/utils"
)

type OneToN struct {
	Message     string   `json:"message" validate:"required"`
	MessageType string   `json:"message_type" validate:"required"`
	Sender      string   `json:"sender" validate:"required"`
	SendTime    string   `json:"send_time"`
	Numbers     []string `json:"numbers" validate:"required"`
}

func (otn OneToN) ToPayload() model.RequestOneToN {
	var dto = model.RequestOneToN{
		Messages: []model.OneToNMessage{
			{
				From:             otn.Sender,
				Text:             utils.RecomposeMessage(otn.Message, otn.MessageType),
				CallbackData:     "cb",
				LanguageEncoding: utils.GetMessageType(otn.MessageType),
				ValidityPeriod:   2880,
			},
		},
	}

	for _, number := range otn.Numbers {
		dto.Messages[0].Destinations = append(dto.Messages[0].Destinations, model.MessageDestination{
			To: utils.CleanupPhone(number),
		})
	}

	return dto
}

func (otn OneToN) ToEvent() *event.OneToNEvent {
	return &event.OneToNEvent{
		EventType: schema.MobilisimOneToNEventType,
		EventData: event.OneToNEventData{
			Message:     otn.Message,
			MessageType: otn.MessageType,
			Sender:      otn.Sender,
			SendTime:    otn.SendTime,
			Numbers:     otn.Numbers,
		},
	}
}
