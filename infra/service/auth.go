package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// AuthModel is the request body expected to a Auth request
type AuthModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate verifies if AuthModel data supplied by the request is valid
func (a AuthModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Name, validation.Required, validation.RuneLength(2, 100)))
}

// Authenticate generates a JWT token based on given data
func Authenticate(auth AuthModel) (string, error) {
	token, err := GenerateJWT(auth.Email, auth.Name)

	return token, err
}
