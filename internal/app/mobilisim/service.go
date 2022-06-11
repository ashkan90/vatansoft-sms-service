package mobilisim

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/event/schema"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/rabbit"
	"vatansoft-sms-service/pkg/utils"
)

type Service interface {
	OneToN(ctx context.Context, req model.RequestOneToN) (*model.ResourceOneToN, error)
	OneToNEvent(ctx context.Context, e *event.OneToNEvent)
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
		return nil, errors.New("something went wrong at <OneToN>(...) " + err.Error())
	}
	return res, nil
}

func (ms *mobilisimService) OneToNEvent(ctx context.Context, e *event.OneToNEvent) {
	batch := utils.Chunk(e.EventData.Numbers, utils.DefaultChunkSize)
	e.Free()

	var err error

	for _, numbers := range batch {
		err = ms.mqProducer.PublishOnQueue(e.ToPrepareQueue(numbers), schema.MobilisimQueueName)
		if err != nil {
			ms.logger.WithContext(ctx).WithField("userID", "x").Error("something went wrong at <OneToNEvent>(...) " + err.Error())
		}
	}

	// TODO: will be moved to another helper or util package idk rn.
	var lenFunc = func() int {
		ln := len(batch) - 1

		// First segment of length calculation.
		fSegment := ln * utils.DefaultChunkSize // [0 -> 9][...] -> [0 - 8][...] * utils.DefaultChunkSize

		// Last segment of length calculation.
		lSegment := len(batch[ln])

		return fSegment + lSegment
	}

	ms.logger.WithField("queued_message_length", lenFunc()).Info("length")
}
