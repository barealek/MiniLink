package minis

import (
	"github.com/barealek/minilink/utility"
	"github.com/gofiber/fiber/v2"
)

func getMiniData(c *fiber.Ctx) error {
	miniID := c.Params("mini")
	miniLink, err := utility.GetMini(miniID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("MiniLink not found.")
	}

	return c.JSON(miniLink)
}
