package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jcblastor/restapi-fiber/routes"
)

func main() {
	app := fiber.New()

	routes.MoviesRoutes(app)

	app.Listen(":8080")
}
