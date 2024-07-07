package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func helloWorld(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Setup routes.
	setupRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
