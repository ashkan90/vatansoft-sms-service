package event

import (
	"encoding/json"
	"vatansoft-sms-service/pkg/event/schema"
)

type Factory interface {
	Type() string
	Data() interface{}
}

type Creation interface {
	Make(t string, data []byte) (Factory, error)
}

type eventCreator struct{}

func NewEventCreator() Creation {
	return &eventCreator{}
}

func (ec *eventCreator) Make(t string, data []byte) (Factory, error) {
	event := ec.getEventByType(t)
	if event == nil {
		return nil, schema.MobilisimUnexpectedEventType
	}

	if err := json.Unmarshal(data, &event); err != nil {
		return nil, err
	}

	return event, nil
}

func (ec *eventCreator) getEventByType(t string) Factory {
	switch t {
	case schema.MobilisimOneToNEventType:
		return &OneToNEvent{}

	default:
		return nil
	}
}
