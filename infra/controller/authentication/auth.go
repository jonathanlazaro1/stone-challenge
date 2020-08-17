package authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/domain"
	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const errCouldntParseAuthModel = "Couldn't parse auth model"
const errCouldntGenerateToken = "Couldn't generate auth token"

// HandleAuth processes requests to authentication
// @Summary Authenticate
// @Description Generates a JWT token that can be used to consume Invoice endpoints.
// @Tags auth
// @Accept json
// @Produce  json
// @Param authInfo body AuthRequestModel true "Auth Model. All fields are required."
// @Success 200 {object} AuthResponseModel "Response containing the JWT token that has been generated"
// @Failure 400 {string} string "Indicates a failure when parsing request body or a validation error, e.g. a required field is missing"
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /auth [post]
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var model AuthRequestModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParseAuthModel)
		return
	}

	err = model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprint(err))
		return
	}

	token, err := service.Authenticate(model.Email, model.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		io.WriteString(w, errCouldntGenerateToken)
		return
	}

	json.NewEncoder(w).Encode(&AuthResponseModel{Token: token})

}

// HandleAuthInfo handles requests to get the authenticated user info
// @Summary Get authentication Info
// @Description Get authentication info, according to the token passed in the request header.
// @Tags auth
// @Produce  json
// @Security JwtAuth
// @Success 200 {object} domain.AuthInfo
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /auth [get]
func HandleAuthInfo(w http.ResponseWriter, r *http.Request) {

	authInfo := r.Context().Value(middleware.RequestAuthInfo).(domain.AuthInfo)

	json.NewEncoder(w).Encode(domain.AuthInfo{Name: authInfo.Name, Email: authInfo.Email})
}
