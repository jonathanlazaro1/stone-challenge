package invoice

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const errInvoiceNotFound = "Couldn't find invoice"

// GetHandler handles a request to an Invoice by its id
func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errInvoiceNotFound)
		return
	}

	service := service.BuildInvoiceService()
	invoice, err := service.Get(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if invoice == nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errInvoiceNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
}
