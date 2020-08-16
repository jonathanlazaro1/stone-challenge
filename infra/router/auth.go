package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller/authentication"
	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
)

func addAuthHandler(r *mux.Router) {
	r.Methods("POST", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(authentication.HandleAuth)

	authInfo := r.Path("/").PathPrefix("").Methods("GET", "OPTIONS").Subrouter()
	authInfo.Use(middleware.AddJwtAuthentication)
	authInfo.Path("/").PathPrefix("").HandlerFunc(authentication.HandleAuthInfo)
}
