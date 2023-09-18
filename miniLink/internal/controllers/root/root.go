package root

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterControllers(app fiber.Router) {
	root := app.Group("/")
	root.Get("/", index)
	root.Get("/:mini", redirectToMini)
}
