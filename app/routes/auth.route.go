package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/go-fiber/app/controllers"
	middleware "github.com/subham043/go-fiber/pkg/middlewares"
)

// AuthRoutes func for describe group of auth routes.
func AuthRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1/auth")

	// route.Use(limiter.New(configs.LowLimiterConfig()))

	// Routes for POST method:
	route.Post("/register", controllers.UserSignUp) // register a new user
	route.Post("/login", controllers.UserSignIn)    // auth, return Access & Refresh tokens

	route.Get("/users", middleware.JWTProtected(), controllers.GetAllUsers) // register a new user
}
