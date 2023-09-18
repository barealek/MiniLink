package minis

import (
	"github.com/barealek/minilink/utility"

	"github.com/gofiber/fiber/v2"
)

var (
	createMiniLinkBody struct {
		Url string `json:"url"`
	}
)

func createMiniLink(c *fiber.Ctx) error {
	// parse body
	if err := c.BodyParser(&createMiniLinkBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Missing body parameter. Check the docs for more info.")
	}

	// create minilink
	miniLink, err := utility.CreateMini(createMiniLinkBody.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something went wrong on the server side.")
	}

	// return minilink
	return c.JSON(miniLink)
}
