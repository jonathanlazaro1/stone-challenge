package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infrastructure/router/home"
)

// Router returns a instance of mux.Router ready to listen and respond to requests
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", home.Index).Methods("GET", "OPTIONS")

	return router
}
