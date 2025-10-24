package handler

import (
	"API-project-go/internal/logger"
	"API-project-go/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var users []models.User
var nextID = 1

// Calculate age from DOB
func calculateAge(dob string) int {
	t, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return 0
	}
	now := time.Now()
	age := now.Year() - t.Year()
	if now.YearDay() < t.YearDay() {
		age--
	}
	return age
}

// POST /users
func CreateUser(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		logger.Logger.Error("Failed to parse request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	logger.Logger.Info("User created", zap.String("name", newUser.Name))
	return c.Status(fiber.StatusCreated).JSON(newUser)
}

// GET /users
func ListUsers(c *fiber.Ctx) error {
	var result []models.User
	for _, u := range users {
		userWithAge := u
		userWithAge.Age = calculateAge(u.DOB)
		result = append(result, userWithAge)
	}
	return c.JSON(result)
}

// GET /users/:id
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	for _, u := range users {
		if u.ID == id {
			u.Age = calculateAge(u.DOB)
			return c.JSON(u)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}

// PUT /users/:id
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var updated models.User
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	for i, u := range users {
		if u.ID == id {
			users[i].Name = updated.Name
			users[i].DOB = updated.DOB
			return c.JSON(users[i])
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}

// DELETE /users/:id
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}
