package authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const errCouldntParseAuthModel = "Couldn't parse auth model"
const errCouldntGenerateToken = "Couldn't generate auth token"

// HandleAuth processes requests to authentication
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var model service.AuthModel
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

	token, err := service.Authenticate(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		io.WriteString(w, errCouldntGenerateToken)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: token})

}

// HandleAuthInfo handles requests to get the authenticated user info
func HandleAuthInfo(w http.ResponseWriter, r *http.Request) {
	type AuthInfo = struct {
		sub  string
		name string
	}
	authInfo := r.Context().Value(middleware.RequestAuthInfo).(middleware.ContextAuthClaims)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{Name: authInfo.Name, Email: authInfo.Email})
}
