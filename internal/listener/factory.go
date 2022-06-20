package listener

import (
	"vatansoft-sms-service/internal/listener/consumer"
	"vatansoft-sms-service/internal/listener/handler"
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/event/schema"
	"vatansoft-sms-service/pkg/eventmanager"
)

type eventHandlerFactory struct {
	service consumer.Service
}

func NewEventHandlerFactory(cs consumer.Service) eventmanager.IEventHandlerFactory {
	return &eventHandlerFactory{
		service: cs,
	}
}

func (eh *eventHandlerFactory) Make(e event.Factory) (eventmanager.EventHandler, error) {
	if e.Type() == schema.MobilisimOneToNEventType {
		return handler.NewOneToNConsumerHandler(eh.service, e.Data().(*event.OneToNEvent)), nil
	}

	return nil, schema.MobilisimUnexpectedHandlerType
}
