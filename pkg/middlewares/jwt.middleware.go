package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/subham043/go-fiber/pkg/utils"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/contrib/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "user", // used in private routes
		ErrorHandler: utils.JWTErrorHandler,
	}

	return jwtMiddleware.New(config)
}

func JWTRefreshProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_REFRESH_KEY"))},
		ContextKey:   "user", // used in private routes
		ErrorHandler: utils.JWTErrorHandler,
	}

	return jwtMiddleware.New(config)
}
