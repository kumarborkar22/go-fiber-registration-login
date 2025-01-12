// package main

// import (
// 	"go-fiber-app/database"
// 	"go-fiber-app/routes"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Load environment variables
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Connect to database
// 	database.ConnectDB()

// 	// Initialize Fiber app
// 	app := fiber.New()

// 	// Set up routes
// 	routes.SetupRoutes(app)

// 	// Start server
// 	log.Fatal(app.Listen(":3000"))
// }

// main.go

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
