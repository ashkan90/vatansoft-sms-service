package app

import (
	"github.com/gofiber/fiber/v2"
	"vatansoft-sms-service/internal/app/handler"
)

type RouteCtx struct {
	App *fiber.App
}

type Router interface {
	SetupRoutes(r *RouteCtx)
}

type route struct {
	mobilisimHandler handler.MobilisimHandler
}

func NewRoute(mobilisimHandler handler.MobilisimHandler) Router {
	return &route{
		mobilisimHandler: mobilisimHandler,
	}
}

func (r *route) SetupRoutes(rc *RouteCtx) {
	v1 := rc.App.Group("/v1")

	r.mobilisimRoutes(v1)
}

func (r *route) mobilisimRoutes(gr fiber.Router) {
	mobilisim := gr.Group("/mobilisim")

	mobilisim.Post("/oneToN", nil)
}
