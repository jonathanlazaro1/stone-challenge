package invoice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/domain"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const maxItemsPerPage = 50
const errParsingPageNumber = "Error parsing page number"
const errParsingItemsPerPage = "Error parsing items per page"
const errMaxItemsPerPageAllowed = "Max items per page allowed: %v"
const errParsingSortParams = "Error parsing sort params"

type getManyResult struct {
	Items      []domain.Invoice `json:"items"`
	TotalItems int64            `json:"totalItems"`
}

// GetManyHandler handles a request to many Invoices
func GetManyHandler(svc service.Invoice) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParams := r.URL.Query()

		page, err := parseParamToInt(pathParams, "p", 1)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errParsingPageNumber)
			return
		}

		itemsPerPage, err := parseParamToInt(pathParams, "itemsperpage", maxItemsPerPage)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errParsingItemsPerPage)
			return
		}

		if itemsPerPage > maxItemsPerPage {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprintf(errMaxItemsPerPageAllowed, maxItemsPerPage))
			return
		}

		filterBy := parseFilterByToMap(pathParams)
		sortBy := parseSortByToMap(pathParams)

		invoices, total, err := svc.GetMany(itemsPerPage, page, filterBy, sortBy)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Unable to fetch invoices")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getManyResult{
			Items:      invoices,
			TotalItems: total,
		})
	}
}
