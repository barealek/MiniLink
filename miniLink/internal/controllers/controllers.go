package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/barealek/minilink/internal/controllers/minis"
	"github.com/barealek/minilink/internal/controllers/root"
)

func RegisterGroups(app *fiber.App) {
	root.RegisterControllers(app.Group("/"))
	minis.RegisterControllers(app.Group("/mini"))
}
