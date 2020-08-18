package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/domain"
	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
)

// AuthGetInfoHandler handles requests to get the authenticated user info
func AuthGetInfoHandler(w http.ResponseWriter, r *http.Request) {

	authInfo := r.Context().Value(middleware.RequestAuthInfo).(domain.AuthInfo)

	json.NewEncoder(w).Encode(domain.AuthInfo{Name: authInfo.Name, Email: authInfo.Email})
}
