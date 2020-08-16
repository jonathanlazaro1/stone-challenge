package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller/home"
)

// Router returns a instance of mux.Router ready to listen and respond to requests
func Router() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", home.IndexHandler).Methods("GET", "OPTIONS")

	apiv1 := router.PathPrefix("/api/v1/").Subrouter()

	addAuthHandler(apiv1.PathPrefix("/auth").Subrouter())
	addInvoiceHandler(apiv1.PathPrefix("/invoice").Subrouter())

	return router
}
