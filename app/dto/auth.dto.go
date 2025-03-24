package dto

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (payload SignUpPayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(&payload.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&payload.Email, validation.Required, is.Email, validation.Length(1, 255)),
		validation.Field(&payload.Password, validation.Required, validation.Length(1, 255)),
	)
}

func (payload *SignUpPayload) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	payload.Password = string(hashedPassword)
	return nil
}

type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (payload SignInPayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(&payload.Email, validation.Required, is.Email, validation.Length(1, 255)),
		validation.Field(&payload.Password, validation.Required, validation.Length(1, 255)),
	)
}
