package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"vatansoft-sms-service/internal/app/dto/request"
	"vatansoft-sms-service/internal/app/orchestration"
)

type MobilisimHandler interface {
	OneToN(c *fiber.Ctx) error
}

type mobilisimHandler struct {
	mobilisimOrchestrator orchestration.MobilisimOrchestrator
}

func NewMobilisimHandler(mo orchestration.MobilisimOrchestrator) MobilisimHandler {
	return &mobilisimHandler{
		mobilisimOrchestrator: mo,
	}
}

func (m *mobilisimHandler) OneToN(c *fiber.Ctx) error {
	var req request.OneToN
	if c.BodyParser(&req) != nil {
		return c.JSON(http.StatusBadRequest)
	}

	m.mobilisimOrchestrator.OneToN(c.Context(), req)

	return nil
}
