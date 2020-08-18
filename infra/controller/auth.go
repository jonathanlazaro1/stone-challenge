package controller

import (
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/handler"
)

// AuthController groups handlers and injects services for auth-related requests
type AuthController struct {
}

// Authenticate processes requests to authentication
// @Summary Authenticate
// @Description Generates a JWT token that can be used to consume Invoice endpoints.
// @Tags auth
// @Accept json
// @Produce  json
// @Param authInfo body handler.AuthRequestModel true "Auth Model. All fields are required."
// @Success 200 {object} handler.AuthResponseModel "Response containing the JWT token that has been generated"
// @Failure 400 {string} string "Indicates a failure when parsing request body or a validation error, e.g. a required field is missing"
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /auth [post]
func (controller *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	handler.AuthPostTokenHandler(w, r)
}

// GetAuthInfo handles requests to get the authenticated user info
// @Summary Get authentication Info
// @Description Get authentication info, according to the token passed in the request header.
// @Tags auth
// @Produce  json
// @Security JwtAuth
// @Success 200 {object} domain.AuthInfo
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /auth [get]
func (controller *AuthController) GetAuthInfo(w http.ResponseWriter, r *http.Request) {
	handler.AuthGetInfoHandler(w, r)
}

const errCouldntParseAuthModel = "Couldn't parse auth model"
const errCouldntGenerateToken = "Couldn't generate auth token"
