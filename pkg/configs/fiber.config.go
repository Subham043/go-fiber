package configs

import (
	"os"
	"strconv"
	"time"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/go-fiber/pkg/utils"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		JSONEncoder:              json.Marshal,
		JSONDecoder:              json.Unmarshal,
		ServerHeader:             "Fiber",
		AppName:                  "Test App v1.0.1",
		CaseSensitive:            true,
		EnablePrintRoutes:        false,
		EnableSplittingOnParsers: true,
		Prefork:                  false,
		ReadTimeout:              time.Second * time.Duration(readTimeoutSecondsCount),
		ErrorHandler:             utils.ErrorHandler,
	}
}
