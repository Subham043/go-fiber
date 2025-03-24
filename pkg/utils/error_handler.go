package utils

import (
	"errors"
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomFiberError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Implement the error interface
func (e *CustomFiberError) Error() string {
	return e.Message
}

// Helper function to create a new CustomFiberError
func NewCustomFiberError(code int, message string) error {
	return &CustomFiberError{Code: code, Message: message}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	path := c.Path()
	response := fiber.Map{
		"path":       path,
		"statusCode": fiber.StatusInternalServerError,
		"message":    err.Error(),
	}

	var customErr *CustomFiberError
	var fiberErr *fiber.Error
	var validationErr validation.Errors

	switch {
	case errors.As(err, &customErr):
		response["statusCode"] = customErr.Code
		response["message"] = customErr.Message
		return c.Status(customErr.Code).JSON(response)

	case errors.As(err, &fiberErr):
		response["statusCode"] = fiberErr.Code
		response["message"] = fiberErr.Message
		return c.Status(fiberErr.Code).JSON(response)

	case errors.As(err, &validationErr):
		response["statusCode"] = fiber.StatusBadRequest
		response["message"] = fiber.ErrBadRequest.Error()
		response["errors"] = err.(validation.Errors).Filter()
		return c.Status(fiber.StatusBadRequest).JSON(response)

	case errors.Is(err, gorm.ErrRecordNotFound):
		response["statusCode"] = fiber.StatusNotFound
		return c.Status(fiber.StatusNotFound).JSON(response)

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
}

func LimitReachedHandler(c *fiber.Ctx) error {
	// Get the Retry-After value from response headers
	retryAfterStr := c.GetRespHeader(fiber.HeaderRetryAfter, "0")

	// Convert the string to an integer (seconds)
	retryAfterSeconds, _ := strconv.Atoi(retryAfterStr)

	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"message": fmt.Sprintf("Too many requests, try again in %v seconds", retryAfterSeconds),
	})
}

func JWTErrorHandler(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": fiber.ErrUnauthorized.Error(),
	})
}
