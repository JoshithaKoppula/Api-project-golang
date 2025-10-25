package routes

import (
	"API-project-go/internal/handler"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running successfully")
	})

	app.Post("/users", handler.CreateUser)
	app.Get("/users", handler.ListUsers)
	app.Get("/users/:id", handler.GetUser)
	app.Put("/users/:id", handler.UpdateUser)
	app.Delete("/users/:id", handler.DeleteUser)
}
