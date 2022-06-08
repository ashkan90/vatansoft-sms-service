// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitAll(l *logrus.Logger, mc mobilisimclient.Client) app.Router {
	mobilisimOrchestrator := orchestration.NewMobilisimOrchestrator()
	mobilisimHandler := handler.NewMobilisimHandler(mobilisimOrchestrator)
	router := app.NewRoute(mobilisimHandler)
	return router
}

// wire.go:

var serviceProviders = wire.NewSet(mobilisim.NewMobilisimService)

var orchestratorProviders = wire.NewSet(orchestration.NewMobilisimOrchestrator)

var handlerProviders = wire.NewSet(handler.NewMobilisimHandler)

var allProviders = wire.NewSet(
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
)
