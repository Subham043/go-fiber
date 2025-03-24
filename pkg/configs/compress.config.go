package configs

import (
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// CompressConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/compress
func CompressConfig() compress.Config {
	// Define compress settings.

	// Return compress configuration.
	return compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}
}
