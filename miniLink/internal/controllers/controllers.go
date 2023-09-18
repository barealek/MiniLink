package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/barealek/minilink/internal/controllers/minis"
	"github.com/barealek/minilink/internal/controllers/root"
)

func RegisterGroups(app *fiber.App) {
	api := app.Group("/api")
	root.RegisterControllers(api)
	minis.RegisterControllers(api.Group("/mini"))
}
