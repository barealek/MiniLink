package root

import (
	"github.com/barealek/minilink/utility"
	"github.com/gofiber/fiber/v2"
)

func redirectToMini(c *fiber.Ctx) error {
	miniId := c.Params("mini")

	mini, err := utility.GetMini(miniId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	go utility.IncrementClicks(miniId)

	return c.Redirect(mini.Redirect, fiber.StatusMovedPermanently)
}
