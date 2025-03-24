package configs

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CorsConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/cors
func CorsConfig() cors.Config {
	// Define cors settings.

	// Return Cors configuration.
	return cors.Config{
		AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
		AllowMethods:     os.Getenv("ALLOW_METHODS"),
		AllowHeaders:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: false,
		ExposeHeaders:    "Retry-After",
		MaxAge:           12 * 60 * 60,
	}
}
