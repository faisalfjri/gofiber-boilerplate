package routers

import (
	BookController "github.com/faisalfjri/gofiber-boilerplate/controllers"
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	// home
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// books
	api := app.Group("/api")
	book := api.Group("/books")
	book.Get("/", BookController.Index)
	book.Post("/", BookController.Store)
	book.Get("/:id", BookController.Show)
	book.Put("/:id", BookController.Update)
	book.Delete("/:id", BookController.Destroy)
}
