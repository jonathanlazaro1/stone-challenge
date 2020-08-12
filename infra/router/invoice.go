package router

import (
	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/controller/invoice"
)

// addInvoiceHandler tells a mux.Router how to handle requests to /invoice
func addInvoiceHandler(r *mux.Router) {
	r.HandleFunc("/", invoice.GetManyHandler).Methods("GET", "OPTIONS")
}
