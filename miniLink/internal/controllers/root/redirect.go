package root

import (
	"github.com/barealek/minilink/utility"
	"github.com/gofiber/fiber/v2"
)

func redirectToMini(c *fiber.Ctx) error {
	miniId := c.Params("mini")

	mini, err := utility.GetMini(miniId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("That minlink wasn't found.")
	}

	go utility.IncrementClicks(miniId)

	return c.Redirect(mini.Redirect, fiber.StatusMovedPermanently)
}
