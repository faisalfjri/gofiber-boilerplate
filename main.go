package main

import (
	"github.com/faisalfjri/gofiber-boilerplate/models"
	"github.com/faisalfjri/gofiber-boilerplate/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()
	app := fiber.New()

	routers.Api(app)
	routers.NotFound(app)

	app.Listen(":8080")
}
