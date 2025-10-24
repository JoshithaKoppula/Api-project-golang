package main

import (
	"API-project-go/internal/logger"
	"API-project-go/internal/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Logger.Sync()
	app := fiber.New()
	routes.SetupRoutes(app)

	logger.Logger.Info("Server running on :8080")
	fmt.Println("Server running on :8080")

	if err := app.Listen(":8080"); err != nil {
		logger.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
