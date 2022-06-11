package mobilisim

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/rabbit"
)

type Service interface {
	OneToN(ctx context.Context, req model.RequestOneToN) (*model.ResourceOneToN, error)
	OneToNEvent(ctx context.Context, e event.OneToNEvent)
	Test(ctx context.Context) error
}

type mobilisimService struct {
	logger          *logrus.Logger
	mobilisimClient mobilisimclient.Client
	mqProducer      rabbit.Client
}

func NewMobilisimService(l *logrus.Logger, mc mobilisimclient.Client, mqp rabbit.Client) Service {
	return &mobilisimService{
		logger:          l,
		mobilisimClient: mc,
		mqProducer:      mqp,
	}
}

func (ms *mobilisimService) OneToN(ctx context.Context, req model.RequestOneToN) (*model.ResourceOneToN, error) {
	res, err := ms.mobilisimClient.OneToN(ctx, req)
	if err != nil {
		ms.logger.Error(err)
		return nil, errors.New("something went wrong at <OneToN>() " + err.Error())
	}
	return res, nil
}

func (ms *mobilisimService) OneToNEvent(ctx context.Context, e event.OneToNEvent) {

}

func (ms *mobilisimService) Test(ctx context.Context) error {
	err := ms.mqProducer.PublishOnQueue([]byte(`{"data": "selam"}`), "sms-queue")
	if err != nil {
		panic("something went wrong while publishing to queue" + err.Error())
	}

	return nil
}
