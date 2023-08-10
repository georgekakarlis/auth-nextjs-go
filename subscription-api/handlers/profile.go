package handlers

import (
	"submanager/database"

	"github.com/gofiber/fiber/v2"
)


func Profile(c *fiber.Ctx) error {
	sess, err := database.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Internal server error",
		})
	}

	userID := sess.Get("user_id")
	username := sess.Get("username")

	if userID == nil || username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"message": "Unauthorized",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"user": fiber.Map{
			"id":       userID,
			"username": username,
		},
	})
}