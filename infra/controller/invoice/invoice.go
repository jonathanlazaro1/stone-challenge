package invoice

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/controller"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// GetManyHandler returns an array of Invoices
func GetManyHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := r.URL.Query()

	page, err := controller.ParsePageNumber(pathParams)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `Error parsing page No.`)
		return
	}

	service := service.BuildInvoiceService()

	filterByMap := make(map[string]string)
	sortByMap := make(map[string]bool)

	invoices, err := service.GetMany(10, page, filterByMap, sortByMap)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
