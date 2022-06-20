package event

import (
	"encoding/json"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/utils"
)

type OneToNEvent struct {
	EventType string
	EventData OneToNEventData
}

type OneToNEventData struct {
	Message     string   `json:"message"`
	MessageType string   `json:"message_type"`
	Sender      string   `json:"sender"`
	SendTime    string   `json:"send_time"`
	Numbers     []string `json:"numbers"`
}

type ResourceOneToNEvent struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	NumberCount int    `json:"numberCount"`
	SMSQuantity int    `json:"quantity"`
}

func (e *OneToNEvent) Type() string {
	return e.EventType
}

func (e *OneToNEvent) Data() interface{} {
	return e
}

func (e *OneToNEvent) ToAsyncPayload(numLn int) *ResourceOneToNEvent {
	return &ResourceOneToNEvent{
		Status:      constants.MobilisimSuccessStatus,
		Description: constants.MobilisimSuccessDescription,
		NumberCount: numLn,
		SMSQuantity: utils.GetMessageQuantity(e.EventData.Message, e.EventData.MessageType) * numLn,
	}
}

func (e *OneToNEvent) Free() {
	e.EventData.Numbers = nil
}

func (e OneToNEvent) ToPrepareQueue(numbers []string) []byte {
	var event = &OneToNEvent{
		EventType: e.EventType,
		EventData: OneToNEventData{
			Message:     e.EventData.Message,
			MessageType: e.EventData.MessageType,
			Sender:      e.EventData.Sender,
			SendTime:    e.EventData.SendTime,
			Numbers:     numbers,
		},
	}

	output, _ := json.Marshal(event)

	return output
}
