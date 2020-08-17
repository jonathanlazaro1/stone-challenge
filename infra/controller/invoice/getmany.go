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
// @Summary List invoices
// @Description Fetch invoices according to query. Only active invoices can be fetched.
// @Tags invoices
// @Produce  json
// @Security JwtAuth
// @Param itemsperpage query int false "Number of items per page" minimum(1) maximum(50) default(50)
// @Param p query int false "Page to fetch, given a number of items per page" minimum(1) default(1)
// @Param filter query string false "Filter results by one or more of comma-separated queries. A query has the following structure: [filter_name]:[value]. Possible filters are: Reference Year = value (reference_year:value), Reference Month = value (reference_month:value) and Document contains value (document:value). Queries are inclusive."
// @Param sort query string false "Sort results by one or more of comma-separated sort items. A sort item has the sort field name, followed by (optionally) a boolean indicating if the sort is in descending order. Sort items have the following structure: [sort_name]:[descending]. Possible sort fields are: Reference Year (reference_year:bool), Reference Month (reference_month:bool) and Document (document:bool). Sorts are inclusive."
// @Success 200 {object} getManyResult "Returns an object containing the array of Invoices found, among an integer indicating the total number of items for the query made."
// @Failure 400 {string} string "Indicates a failure when parsing query params, or a itemsperpage query param greater than max value"
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice [get]
func GetManyHandler(w http.ResponseWriter, r *http.Request) {
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

	svc := service.BuildInvoiceService()

	filterBy := parseFilterByToMap(pathParams)
	sortBy := parseSortByToMap(pathParams)

	invoices, total, err := svc.GetMany(itemsPerPage, page, filterBy, sortBy)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Unable to fetch invoices")
		return
	}

	json.NewEncoder(w).Encode(getManyResult{
		Items:      invoices,
		TotalItems: total,
	})
}
