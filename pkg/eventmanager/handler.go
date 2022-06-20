package eventmanager

import (
	"context"
	"vatansoft-sms-service/pkg/event"
)

type EventHandler interface {
	Handle(ctx context.Context) error
}

type IEventHandlerFactory interface {
	Make(e event.Factory) (EventHandler, error)
}
