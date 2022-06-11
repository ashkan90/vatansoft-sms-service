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
	if bErr := c.BodyParser(&req); bErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Body parser gave error. For detail: " + bErr.Error(),
		})
	}

	// TODO: Implement request validator.

	res, err := m.mobilisimOrchestrator.OneToN(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(res)
}
