package minis

import (
	"github.com/barealek/minilink/utility"

	"github.com/gofiber/fiber/v2"
)

func getAllMinis(c *fiber.Ctx) error {
	minis, err := utility.GetAllMinis()
	if err != nil {
		panic(err)
	}
	if len(minis) == 0 {
		return c.JSON([]string{})
	}
	return c.JSON(minis)

}
