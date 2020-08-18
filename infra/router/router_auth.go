package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller"
	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
)

func addAuthHandler(r *mux.Router) {
	controller := controller.AuthController{}
	r.Methods("POST", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(controller.Authenticate)

	authInfo := r.Path("/").PathPrefix("").Methods("GET", "OPTIONS").Subrouter()
	authInfo.Use(middleware.AddJwtAuthentication)
	authInfo.Path("/").PathPrefix("").HandlerFunc(controller.GetAuthInfo)
}
