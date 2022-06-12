package event

import (
	"encoding/json"
	"vatansoft-sms-service/pkg/response"
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

func (e *OneToNEvent) Type() string {
	return e.EventType
}

func (e *OneToNEvent) Data() interface{} {
	return e
}

func (e *OneToNEvent) Response() interface{} {
	return response.NewMobilisimSuccessResponse(e.EventData.Message, e.EventData.MessageType)
}

func (e *OneToNEvent) Free() {
	e.EventData.Numbers = nil
}

func (e OneToNEvent) ToPrepareQueue(numbers []string) []byte {
	var event = OneToNEvent{
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
