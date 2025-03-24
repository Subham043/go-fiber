package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/go-fiber/app/routes"
	"github.com/subham043/go-fiber/pkg/configs"
	middleware "github.com/subham043/go-fiber/pkg/middlewares"
	"github.com/subham043/go-fiber/pkg/utils"
	"github.com/subham043/go-fiber/platform/database"
	"github.com/subham043/go-fiber/platform/redis"

	// _ "github.com/create-go-app/fiber-go-template/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// connect to database.
	database.ConnectDB()

	// connect to redis.
	redis.RedisConnection()

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Serve static files from the "public" directory
	app.Static("/public", "./public", fiber.Static{
		Compress: true,
		Download: false,
		Browse:   true,
	})

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	// routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.AuthRoutes(app) // Register a public routes for app.
	// routes.PrivateRoutes(app) // Register a private routes for app.

	// Start server with graceful shutdown.
	utils.StartServerWithGracefulShutdown(app)
}
