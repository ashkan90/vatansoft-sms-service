package orchestration

import (
	"context"
	"vatansoft-sms-service/internal/app/dto/request"
	"vatansoft-sms-service/internal/app/mobilisim"
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

type MobilisimOrchestrator interface {
	OneToN(ctx context.Context, dto request.OneToN) (*model.ResourceOneToN, error)
}

type mobilisimOrchestrator struct {
	mobilisimService mobilisim.Service
}

func NewMobilisimOrchestrator(ms mobilisim.Service) MobilisimOrchestrator {
	return &mobilisimOrchestrator{
		mobilisimService: ms,
	}
}

func (m *mobilisimOrchestrator) OneToN(ctx context.Context, dto request.OneToN) (*model.ResourceOneToN, error) {
	if len(dto.Numbers) > constants.MaxMessageInTime {
		m.mobilisimService.OneToNEvent(ctx, dto.ToEvent())
		return nil, nil
	}
	return m.mobilisimService.OneToN(ctx, dto.ToPayload())
}
