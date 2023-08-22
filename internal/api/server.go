package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ikraamdaanis/store-server/internal/api/routes"
)

func RunServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	app := fiber.New()

	routes.AuthRoutes(app)

	log.Println("Server started on port " + port + ".")
	log.Fatal(app.Listen(":" + port))
}
