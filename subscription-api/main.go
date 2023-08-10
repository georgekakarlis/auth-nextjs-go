package main

import (
	"fmt"
	"log"
	"os"

	"submanager/database"
	"submanager/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init () {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	database.ConnectDB()
	// Initialize the session store
	database.InitSessionStore()
}


func main() {
	
	 

	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "http://localhost:3000",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	routes.InitRoutes(app)

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	fmt.Printf(" 🤖Starting up on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}



