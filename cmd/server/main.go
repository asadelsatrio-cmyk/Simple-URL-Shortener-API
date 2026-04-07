package main

import (
	"log"

	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/database"
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Routes
	app.Post("/shorten", handlers.ShortenURL)
	app.Get("/:code", handlers.RedirectURL)
	app.Get("/stats/:code", handlers.GetURLStats)

	log.Println("Server is running on port 3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
