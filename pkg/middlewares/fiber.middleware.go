package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/subham043/go-fiber/pkg/configs"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(configs.CorsConfig()),
		// Add recover middleware.
		recover.New(configs.RecoverConfig()),
		// Add helmet middleware.
		helmet.New(),
		// Add limiter middleware.
		limiter.New(configs.DefaultLimiterConfig()),
		// Add compress middleware.
		compress.New(configs.CompressConfig()),
		// Add simple logger.
		logger.New(),
	)
}
