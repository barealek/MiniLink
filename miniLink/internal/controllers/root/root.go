package root

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterControllers(app *fiber.App) {
	root := app.Group("/")
	root.Get("/", index)
}
