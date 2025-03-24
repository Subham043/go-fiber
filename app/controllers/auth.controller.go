package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/subham043/go-fiber/app/dto"
	"github.com/subham043/go-fiber/app/models"
	"github.com/subham043/go-fiber/app/services"
)

func UserSignIn(c *fiber.Ctx) error {
	// Validate the request payload.
	var payload dto.SignInPayload

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := payload.Validate(); err != nil {
		return err
	}

	// Find the user account in the database by email.
	result, err := services.FindUserByEmail(payload.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if !result.ValidatePassword(payload.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	access_token, err := result.GenerateAccessToken()

	if err != nil {
		return err
	}

	refresh_token, err := result.GenerateRefreshToken()

	if err != nil {
		return err
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged in successfully",
		"data":    result,
		"token": fiber.Map{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	})
}

func UserSignUp(c *fiber.Ctx) error {
	// Validate the request payload.
	var payload dto.SignUpPayload

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := payload.Validate(); err != nil {
		return err
	}

	if err := payload.HashPassword(); err != nil {
		return err
	}

	// Register the user account in the database.
	result, err := services.RegisterUser(models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return err
	}

	access_token, err := result.GenerateAccessToken()

	if err != nil {
		return err
	}

	refresh_token, err := result.GenerateRefreshToken()

	if err != nil {
		return err
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registered successfully",
		"data":    result,
		"token": fiber.Map{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	users := c.Locals("user").(*jwt.Token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users.Claims,
	})
}
