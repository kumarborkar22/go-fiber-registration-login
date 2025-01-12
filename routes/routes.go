// The routes.go file defines the SetupRoutes function, which sets up two POST routes for user registration and login

package routes

import (
	"go-fiber-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
}
