package handlers

import (
	"submanager/database"
	"submanager/models"
	"submanager/util"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)



func Login(c *fiber.Ctx) error {
	//login user struct
	type LoginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	db := database.DB
	user := new(LoginUser)

	//parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not parse the input. Please review your input.",
			"errors": err.Error(),
		})
	}

	//validate user input
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not validate the input. Please review your input.",
			"errors": err.Error(),
		})
	}

	// find user by username
	dbUser := new(models.User)
	db.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"message": "User not found.",
		})
	}

	// Check password hash
	if !util.CheckPasswordHash(user.Password, dbUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"message": "Invalid credentials",
		})
	}

	// After verifying the password, create a session
	sess, err := database.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Internal server error",
		})
	}
	sess.Set("user_id", dbUser.ID)
	sess.Set("username", dbUser.Username)
	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Failed to save session",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Logged in successfully",
	})
}