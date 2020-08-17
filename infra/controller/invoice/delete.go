package invoice

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// DeleteHandler handles a request to delete an Invoice
// @Summary Delete Invoice
// @Description Deactivates an Invoice, which means it will still exist on the server, but won't be capable of being retrieved|updated anymore.
// @Tags invoices
// @Produce  plain
// @Security JwtAuth
// @Param id path int true "Id of the invoice to delete"
// @Success 204 "Invoice was successfully deleted."
// @Failure 400 {string} string "Indicates a failure when parsing Invoice Id."
// @Failure 404 {string} string "Indicates that no Invoice with given Id was found, or Invoice is deactivated"
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice/{id} [delete]
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParseInvoiceID)
		return
	}

	svc := service.BuildInvoiceService()
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
