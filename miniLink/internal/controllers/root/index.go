package root

import (
	"github.com/gofiber/fiber/v2"
)

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
