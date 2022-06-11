package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"vatansoft-sms-service/internal/app/dto/request"
	"vatansoft-sms-service/internal/app/orchestration"
)

type MobilisimHandler interface {
	OneToN(c *fiber.Ctx) error
	Test(c *fiber.Ctx) error
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
		logrus.Error("Body parser gave error. For detail: " + bErr.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	res, err := m.mobilisimOrchestrator.OneToN(c.Context(), req)
	if err != nil {
		logrus.Error(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(res)
}

func (m *mobilisimHandler) Test(c *fiber.Ctx) error {

	return c.JSON(map[string]string{
		"status": "ok",
	})
}
