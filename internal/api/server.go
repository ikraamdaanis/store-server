package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Println("Server started on port " + port + ".")
	log.Fatal(app.Listen(":" + port))
}
