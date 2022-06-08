package orchestration

import (
	"context"
	"vatansoft-sms-service/internal/app/dto/request"
	"vatansoft-sms-service/internal/app/mobilisim"
)

type MobilisimOrchestrator interface {
	OneToN(ctx context.Context, req request.OneToN) error
}

type mobilisimOrchestrator struct {
	mobilisimService mobilisim.Service
}

func NewMobilisimOrchestrator(ms mobilisim.Service) MobilisimOrchestrator {
	return &mobilisimOrchestrator{
		mobilisimService: ms,
	}
}

func (m *mobilisimOrchestrator) OneToN(ctx context.Context, req request.OneToN) error {
	m.mobilisimService.OneToN(ctx, req.ToPayload())
	return nil
}
