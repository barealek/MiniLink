package root

import (
	"github.com/gofiber/fiber/v2"
)

func index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Name":        "MiniLink",
		"Author":      "Aleksander (barealek)",
		"Github":      "github.com/barealek/minilink",
		"VersionType": "Alpha",
	})
}
