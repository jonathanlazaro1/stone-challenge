package invoice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

func updateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParseInvoiceID)
		return
	}

	var model service.PostModel
	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParsePostModel)
		return
	}

	if r.Method == "PUT" {
		err = model.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprint(err))
			return
		}
	}

	newInvoice, err := service.Put(id, model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Couldn't update Invoice")
		return
	} else if newInvoice == nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errInvoiceNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PutHandler handles a request to put an Invoice
// @Summary Update Invoice
// @Description Updates an Invoice under the supplied Id. All values on the Invoice will be updated.
// @Tags invoices
// @Accept json
// @Produce  plain
// @Security JwtAuth
// @Param id path int true "Id of the invoice to update"
// @Param invoice body service.PostModel true "Update Invoice Model. All fields are required."
// @Success 204 "Invoice was successfully updated."
// @Failure 400 {string} string "Indicates a failure when parsing Invoice Id|request body or a validation error, e.g. a required field is missing"
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 404 {string} string "Indicates that no Invoice with given Id was found, or Invoice is deactivated"
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice/{id} [put]
func PutHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r)
}

// PatchHandler handles a request to patch an Invoice
// @Summary Update Invoice
// @Description Updates an Invoice under the supplied Id. Only the supplied field values will be applied to the Invoice.
// @Tags invoices
// @Accept json
// @Produce  plain
// @Security JwtAuth
// @Param id path int true "Id of the invoice to update"
// @Param invoice body service.PostModel true "Update Invoice Model. All fields are optional."
// @Success 204 "Invoice was successfully updated."
// @Failure 400 {string} string "Indicates a failure when parsing Invoice Id|request body."
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 404 {string} string "Indicates that no Invoice with given Id was found, or Invoice is deactivated"
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice/{id} [patch]
func PatchHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r)
}
