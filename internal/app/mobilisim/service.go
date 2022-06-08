package mobilisim

import (
	"context"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

type Service interface {
	OneToN(ctx context.Context, req model.RequestOneToN) (*model.ResourceOneToN, error)
}

type mobilisimService struct {
	logger          *logrus.Logger
	mobilisimClient mobilisimclient.Client
}

func NewMobilisimService(l *logrus.Logger, mc mobilisimclient.Client) Service {
	return &mobilisimService{
		logger:          l,
		mobilisimClient: mc,
	}
}

func (ms *mobilisimService) OneToN(ctx context.Context, req model.RequestOneToN) (*model.ResourceOneToN, error) {
	return nil, nil
}
