package routers

import "github.com/gofiber/fiber/v2"

func NotFound(app *fiber.App) {
	// Register new special route.
	app.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "not found",
			})
		},
	)
}
