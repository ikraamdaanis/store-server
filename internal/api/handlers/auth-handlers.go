package handlers

import (
	"strings"
	"time"

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

func LoginUser(c *fiber.Ctx) error {
	// Parse request body and create a new loginReq
	loginReq := &models.User{}
	foundUser := &models.User{}

	if err := c.BodyParser(loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	err := database.DB.Where("email = ?", loginReq.Email).First(&models.User{}).Scan(&foundUser)

	if err.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email.",
		})
	}

	passwordIsCorrect := utils.CheckPasswordHash(loginReq.Password, foundUser.Password)

	if !passwordIsCorrect {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password.",
		})
	}

	sessionToken, error := utils.CreateSession(foundUser.ID)

	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create session.",
		})
	}

	session := &models.Session{
		UserID:     foundUser.ID,
		Token:      sessionToken,
		UserAgent:  c.Get("User-Agent"),
		IP_Address: c.Context().RemoteAddr().String(),
		ExpiresAt:  utils.TokenExpiry,
	}

	createdSession := &models.Session{}

	err = database.DB.Create(&session).Scan(&createdSession)

	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create session.",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "session"
	cookie.Value = session.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully logged in.",
		"data":    foundUser,
		"session": createdSession,
	})
}
