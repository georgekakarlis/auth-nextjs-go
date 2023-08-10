package handlers

import (
	"submanager/database"
	"submanager/models"
	"submanager/util"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	// NewUser struct
	type NewUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(NewUser)

	// Parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not parse the input. Please review your input.",
			"errors": err.Error(),
		})
	}

	// Validate user input
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"errors": err.Error(),
		})
	}

	// Hash password
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Internal server error",
			"errors": err.Error(),
		})
	}

	// Create user
	newUserRecord := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hash,
	}
	if err := db.Create(newUserRecord).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Internal server error",
			"errors": err.Error(),
		})
	}

	// Return JSON response the User
	newUserResponse := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	// Return JSON response
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Created user",
		"data": newUserResponse,
	})
}
