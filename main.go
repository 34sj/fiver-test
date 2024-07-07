package main

import (
	"log"

	"github.com/34sj/fiver-test/book"
	"github.com/gofiber/fiber/v3"
)

func helloWorld(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Setup routes.
	setupRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
