package handlers

import (
	"submanager/database"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	// Fetch the session
	sess, err := database.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Internal server error",
		})
	}

	// Destroy the session
	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Failed to log out",
		})
	}

	// Optionally clear the session cookie on the client side
	c.ClearCookie("session_id")

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Logged out successfully",
	})
}