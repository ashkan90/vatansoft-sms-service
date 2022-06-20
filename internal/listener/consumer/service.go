package consumer

import (
	"context"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/rabbit"
)

type Service interface {
	OneToN(ctx context.Context, req model.RequestOneToN) error
}

type mobilisimConsumerService struct {
	logger          *logrus.Logger
	mobilisimClient mobilisimclient.Client
	mqProducer      rabbit.Client
}

func NewMobilisimConsumerService(l *logrus.Logger, mc mobilisimclient.Client, mqp rabbit.Client) Service {
	return &mobilisimConsumerService{
		logger:          l,
		mobilisimClient: mc,
		mqProducer:      mqp,
	}
}

func (ms *mobilisimConsumerService) OneToN(ctx context.Context, req model.RequestOneToN) error {
	return ms.mobilisimClient.OneToN(ctx, req)
	//res, err := ms.mobilisimClient.OneToN(ctx, req)
	//if err != nil {
	//	ms.logger.Error(err)
	//	return nil, errors.New("something went wrong at <OneToN>(...) " + err.Error())
	//}
	//return res, nil
}
