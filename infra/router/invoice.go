package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller/invoice"
)

// addInvoiceHandler tells a mux.Router how to handle requests to /invoice
func addInvoiceHandler(r *mux.Router) {
	// Mux demands registering routes with and without trailing slash, so... yeah
	r.Methods("GET", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(invoice.GetManyHandler)

	r.HandleFunc("/{id:[0-9]+}", invoice.GetHandler).Methods("GET", "OPTIONS")

	// Same as above
	r.Methods("POST", "OPTIONS").Path("/").PathPrefix("").HandlerFunc(invoice.PostHandler)

	r.HandleFunc("/{id:[0-9]+}", invoice.UpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/{id:[0-9]+}", invoice.UpdateHandler).Methods("PATCH", "OPTIONS")

	r.HandleFunc("/{id:[0-9]+}", invoice.DeleteHandler).Methods("DELETE", "OPTIONS")
}
