package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ikraamdaanis/store-server/internal/api/models"
	"github.com/ikraamdaanis/store-server/internal/database"
	"github.com/ikraamdaanis/store-server/pkg/utils"
)

func CreateUser(c *fiber.Ctx) error {
	// Parse request body and create a new user
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {

		message := err.Error()

		if strings.Contains(message, "duplicate key") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created user.",
		"data":    user,
	})
}
