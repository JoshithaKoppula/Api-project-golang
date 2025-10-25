package main

import (
	"API-project-go/config"
	"API-project-go/internal/logger"
	"API-project-go/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()

	app := fiber.New()

	db := config.ConnectDB()
	defer db.Close()

	logger.Log.Info("Connected to MySQL Database Successfully!")

	routes.SetupRoutes(app, db)

	logger.Log.Info(" Starting server on port 8080...")
	app.Listen(":8080")
}
