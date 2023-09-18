package minis

import "github.com/gofiber/fiber/v2"

func RegisterControllers(app fiber.Router) {
	// Register the root controller
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This endpoint is for managing minilinks.")
	})

	app.Get("/:mini", getMiniData)
	app.Get("/all", getAllMinis)

	app.Post("/create", createMiniLink)

}
