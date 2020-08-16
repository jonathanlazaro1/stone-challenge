package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller/invoice"
	"github.com/jonathanlazaro1/stone-challenge/infra/router/middleware"
)

// addInvoiceHandler tells a mux.Router how to handle requests to /invoice
func addInvoiceHandler(r *mux.Router) {
	r.Use(middleware.AddJwtAuthentication)

	r.Methods("GET", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(invoice.GetManyHandler)

	r.HandleFunc("/{id:[0-9]+}", invoice.GetHandler).Methods("GET", "OPTIONS")

	r.Methods("POST", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(invoice.PostHandler)

	r.HandleFunc("/{id:[0-9]+}", invoice.UpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/{id:[0-9]+}", invoice.UpdateHandler).Methods("PATCH", "OPTIONS")

	r.HandleFunc("/{id:[0-9]+}", invoice.DeleteHandler).Methods("DELETE", "OPTIONS")
}
