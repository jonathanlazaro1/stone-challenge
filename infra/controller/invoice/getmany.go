package invoice

import (
	"encoding/json"
	"io"
	"net/http"

	helpers "github.com/jonathanlazaro1/stone-challenge/infra/controller"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// maxItemsPerPage defines the max allowed number of items per page in a query
const maxItemsPerPage = 50

// GetManyHandler returns an array of Invoices
func GetManyHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := r.URL.Query()

	page, err := helpers.ParseParamToInt(pathParams, "p", 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `Error parsing page number`)
		return
	}

	itemsPerPage, err := helpers.ParseParamToInt(pathParams, "ipp", maxItemsPerPage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `Error parsing items per page`)
		return
	}

	// if itemsPerPage > maxItemsPerPage {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, fmt.Sprintf("Max allowed items per page: %v", maxItemsPerPage))
	// 	return
	// }

	service := service.BuildInvoiceService()

	filterByMap := make(map[string]string)
	sortByMap := make(map[string]bool)

	invoices, err := service.GetMany(itemsPerPage, page, filterByMap, sortByMap)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
