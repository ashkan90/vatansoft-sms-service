package request

import (
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/utils"
)

type OneToN struct {
	Message     string   `json:"message"`
	MessageType string   `json:"message_type"`
	Sender      string   `json:"sender"`
	SendTime    string   `json:"send_time"`
	Numbers     []string `json:"numbers"`
}

func (otn OneToN) ToPayload() model.RequestOneToN {
	var dto model.RequestOneToN

	dto.Messages.SetMessage(model.OneToNMessage{
		From:            otn.Sender,
		Text:            otn.Message,
		CallbackData:    "cb",
		Transliteration: otn.MessageType,
		ValidityPeriod:  2880,
		Language:        model.MessageLanguage{},
	})

	utils.Chunk(otn.Numbers, utils.DefaultChunkSize)
	dto.Messages.AddReceiver()

	return dto
}
