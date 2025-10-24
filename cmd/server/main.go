package main

import (
	"API-project-go/internal/logger"
	"API-project-go/internal/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	defer logger.Logger.Sync()

	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	logger.Logger.Info("Server running on :8080")
	fmt.Println("Server running on :8080")

	app.Listen(":8080")
}
