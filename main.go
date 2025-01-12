// The main.go file initializes a Fiber web application, loads environment variables, connects to a MySQL database, defines routes for user registration and login, and starts the server on port 3000.

package main

import (
	"go-fiber-app/database"
	"go-fiber-app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	database.ConnectDB()

	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)

	// Start the app on port 3000
	log.Fatal(app.Listen(":3000"))
}
