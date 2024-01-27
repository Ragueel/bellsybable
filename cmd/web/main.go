package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New())

	// app.Post("/<token>/generate", handlers.HandleGenerate)

	app.Listen(":6969")
}
