package mobilisim

import (
	"context"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/pkg/event"
	"vatansoft-sms-service/pkg/event/schema"
	"vatansoft-sms-service/pkg/mobilisimclient"
	"vatansoft-sms-service/pkg/rabbit"
	"vatansoft-sms-service/pkg/utils"
)

type Service interface {
	OneToNEvent(ctx context.Context, e *event.OneToNEvent) *event.ResourceOneToNEvent
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

func (ms *mobilisimService) OneToNEvent(ctx context.Context, e *event.OneToNEvent) *event.ResourceOneToNEvent {
	var (
		batch   = utils.Chunk(e.EventData.Numbers, utils.DefaultChunkSize)
		batchLn = 0
		err     error
	)
	e.Free()

	for _, numbers := range batch {
		batchLn += len(numbers)
		err = ms.mqProducer.PublishOnQueue(
			e.ToPrepareQueue(numbers),
			schema.MobilisimQueueName,
			schema.MobilisimOneToNEventType,
		)
		if err != nil {
			ms.logger.WithContext(ctx).WithField("userID", "x").Error("something went wrong at <OneToNEvent>(...) " + err.Error())
		}
	}

	ms.logger.WithField("queued_message_length", batchLn).Info("length")

	return e.ToAsyncPayload(batchLn)
}
