package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikraamdaanis/store-server/internal/api/handlers"
)

func AuthRoutes(app *fiber.App) {
	userRoutes := app.Group("/auth")

	userRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRoutes.Post("/sign-up", handlers.CreateUser)
}
