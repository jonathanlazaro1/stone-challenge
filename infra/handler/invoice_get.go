package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const errInvoiceNotFound = "Couldn't find invoice"
const errCouldntParseInvoiceID = "Couldn't parse invoice Id"

// InvoiceGetHandler handles a request to an Invoice by its id
func InvoiceGetHandler(svc service.InvoiceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errCouldntParseInvoiceID)
			return
		}
		invoice, err := svc.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if invoice == nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, errInvoiceNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoice)
	}
}
