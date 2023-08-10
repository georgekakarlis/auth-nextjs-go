package routes

import (
	"submanager/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
			"status" : "success",
		})
	})  
	
	app.Post("/api/v1/signup", handlers.Signup)
	app.Post("/api/v1/login", handlers.Login)
	app.Get("/api/v1/profile", handlers.Profile)
	app.Post("/api/v1/logout", handlers.Logout)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}