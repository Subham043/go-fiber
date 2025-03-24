package configs

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/subham043/go-fiber/pkg/utils"
	"github.com/subham043/go-fiber/platform/redis"
)

// DefaultLimiterConfig func for configuration Fiber app.
// See: "github.com/gofiber/fiber/v2/middleware/limiter"
func DefaultLimiterConfig() limiter.Config {
	// Define limiter storage
	storage := redis.NewRedisStorage(redis.REDISCLIENT)
	// Return Limiter configuration.
	return limiter.Config{
		Max:          100,
		Expiration:   60 * time.Second,
		Storage:      storage,
		LimitReached: utils.LimitReachedHandler,
	}
}

// LimiterConfig func for configuration Fiber app.
// See: "github.com/gofiber/fiber/v2/middleware/limiter"
func LowLimiterConfig() limiter.Config {
	// Define limiter storage
	storage := redis.NewRedisStorage(redis.REDISCLIENT)
	// Return Limiter configuration.
	return limiter.Config{
		Max:          3,
		Expiration:   60 * time.Second,
		Storage:      storage,
		LimitReached: utils.LimitReachedHandler,
	}
}
