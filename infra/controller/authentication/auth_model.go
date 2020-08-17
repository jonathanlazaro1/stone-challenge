package authentication

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// AuthRequestModel is the request body expected to a Auth request
type AuthRequestModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate verifies if AuthModel data supplied by the request is valid
func (a AuthRequestModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Name, validation.Required, validation.RuneLength(2, 100)))
}

// AuthResponseModel represents the JWT token generated after authenticating
type AuthResponseModel struct {
	Token string `json:"token"`
}
