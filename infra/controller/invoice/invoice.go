package invoice

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// IndexHandler returns NotFound
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	service := service.BuildInvoiceService()

	filterByMap := make(map[string]string)
	sortByMap := make(map[string]bool)

	invoices, err := service.GetMany(10, 1, filterByMap, sortByMap)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `Error fetching invoices`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
