//go:build wireinject
// +build wireinject

package vatansoft_sms_service

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"vatansoft-sms-service/internal/app"
	"vatansoft-sms-service/internal/app/handler"
	"vatansoft-sms-service/internal/app/mobilisim"
	"vatansoft-sms-service/internal/app/orchestration"
	"vatansoft-sms-service/pkg/mobilisimclient"
)

var serviceProviders = wire.NewSet(
	mobilisim.NewMobilisimService,
)

var orchestratorProviders = wire.NewSet(
	orchestration.NewMobilisimOrchestrator,
)

var handlerProviders = wire.NewSet(
	handler.NewMobilisimHandler,
)

var allProviders = wire.NewSet(
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
)

func InitAll(
	l *logrus.Logger,
	mc mobilisimclient.Client,
) app.Router {
	wire.Build(allProviders, app.NewRoute)
	return nil
}
