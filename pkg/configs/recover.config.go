package configs

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

// RecoverConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/recover
func RecoverConfig() recover.Config {
	// Define recover settings.

	// Return recover configuration.
	return recover.Config{
		EnableStackTrace: os.Getenv("NODE_ENV") != "production",
	}
}
