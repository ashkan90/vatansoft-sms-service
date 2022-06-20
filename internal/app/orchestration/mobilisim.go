package orchestration

import (
	"context"
	"vatansoft-sms-service/internal/app/dto/request"
	"vatansoft-sms-service/internal/app/mobilisim"
	"vatansoft-sms-service/pkg/event"
)

type MobilisimOrchestrator interface {
	OneToN(ctx context.Context, dto request.OneToN) (*event.ResourceOneToNEvent, error)
}

type mobilisimOrchestrator struct {
	mobilisimService mobilisim.Service
}

func NewMobilisimOrchestrator(ms mobilisim.Service) MobilisimOrchestrator {
	return &mobilisimOrchestrator{
		mobilisimService: ms,
	}
}

func (m *mobilisimOrchestrator) OneToN(ctx context.Context, dto request.OneToN) (*event.ResourceOneToNEvent, error) {
	return m.mobilisimService.OneToNEvent(ctx, dto.ToEvent()), nil

}
