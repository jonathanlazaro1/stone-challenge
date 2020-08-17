package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller"
	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
)

// addInvoiceHandler tells a mux.Router how to handle requests to /invoice
func addInvoiceHandler(r *mux.Router) {
	controller := controller.BuildInvoiceController()

	r.Use(middleware.AddJwtAuthentication)

	r.Methods("GET", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(controller.GetMany)

	r.HandleFunc("/{id:[0-9]+}", controller.Get).Methods("GET", "OPTIONS")

	r.Methods("POST", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(controller.Post)

	r.HandleFunc("/{id:[0-9]+}", controller.Put).Methods("PUT", "OPTIONS")

	r.HandleFunc("/{id:[0-9]+}", controller.Patch).Methods("PATCH", "OPTIONS")

	r.HandleFunc("/{id:[0-9]+}", controller.Delete).Methods("DELETE", "OPTIONS")
}
