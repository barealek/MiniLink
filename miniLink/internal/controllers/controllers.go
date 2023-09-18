package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/barealek/minilink/internal/controllers/minis"
	"github.com/barealek/minilink/internal/controllers/root"
)

func RegisterGroups(app *fiber.App) {
	api := app.Group("/a")
	// Since i have a reverse proxy in front of this and also have to serve a front end, I'm using the /a prefix for the api.
	// This is not necessary if you are not using a reverse proxy and are not serving a front end. You can just use app.Group("/").

	root.RegisterControllers(api)
	minis.RegisterControllers(api.Group("/mini"))
}
