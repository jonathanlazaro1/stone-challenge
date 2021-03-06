package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// InvoiceDeleteHandler handles a request to delete an Invoice
func InvoiceDeleteHandler(svc service.InvoiceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errCouldntParseInvoiceID)
			return
		}

		deletedInvoice, err := svc.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Couldn't update Invoice")
			return
		} else if !deletedInvoice {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, errInvoiceNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
