package root

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	root := app.Group("/")
	root.Get("/", index)
}
