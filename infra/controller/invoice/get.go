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
const errCouldntParseInvoiceID = "Couldn't parse invoice Id"

// GetHandler handles a request to an Invoice by its id
// @Summary Get an invoice
// @Description Get an invoice, given its Id. Only an active Invoice can be fetched.
// @Tags invoices
// @Produce  json
// @Security JwtAuth
// @Param id path int true "Id of the invoice to fetch"
// @Success 200 {object} domain.Invoice
// @Failure 400 {string} string "Indicates a failure when parsing invoice Id"
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 404 {string} string "Indicates that no Invoice with given Id was found, or Invoice is deactivated"
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice/{id} [get]
func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParseInvoiceID)
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

	json.NewEncoder(w).Encode(invoice)
}
