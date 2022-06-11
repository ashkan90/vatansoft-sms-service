package request

import (
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

type OneToN struct {
	Message     string   `json:"message"`
	MessageType string   `json:"message_type"`
	Sender      string   `json:"sender"`
	SendTime    string   `json:"send_time"`
	Numbers     []string `json:"numbers"`
}

func (otn OneToN) ToPayload() model.RequestOneToN {
	var dto = model.RequestOneToN{
		Messages: []model.OneToNMessage{
			{
				From:            otn.Sender,
				Text:            otn.Message,
				CallbackData:    "cb",
				Transliteration: otn.MessageType,
				ValidityPeriod:  2880,
				Language: model.MessageLanguage{
					LanguageCode: "TR",
					SingleShift:  false,
					LockingShift: true,
				},
			},
		},
	}

	for _, number := range otn.Numbers {
		dto.Messages[0].Destinations = append(dto.Messages[0].Destinations, model.MessageDestination{
			To:        number,
			MessageID: "1",
		})
	}

	return dto
}

func (otn OneToN) ToEvent() event.OneToNEvent {
	return event.OneToNEvent{
		Message:     otn.Message,
		MessageType: otn.MessageType,
		Sender:      otn.Sender,
		SendTime:    otn.SendTime,
		Numbers:     otn.Numbers,
	}
}
