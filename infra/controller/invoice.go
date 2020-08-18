package controller

import (
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/handler"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// InvoiceController groups handlers and injects services for invoice-related requests
type InvoiceController struct {
	service service.InvoiceService
}

// BuildInvoiceController builds an Invoice Controller, injecting an Invoice Service on it
func BuildInvoiceController() *InvoiceController {
	return &InvoiceController{service: *service.BuildInvoiceService()}
}

// GetMany handles a request to many Invoices
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
func (controller *InvoiceController) GetMany(w http.ResponseWriter, r *http.Request) {
	f := handler.InvoiceGetManyHandler(controller.service)
	f(w, r)
}

// Get handles a request to an Invoice by its id
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
func (controller *InvoiceController) Get(w http.ResponseWriter, r *http.Request) {
	f := handler.InvoiceGetHandler(controller.service)
	f(w, r)
}

// Post handles a request to post an Invoice
// @Summary New Invoice
// @Description Creates an invoice using the values supplied on the request body.
// @Tags invoices
// @Accept json
// @Produce  plain
// @Security JwtAuth
// @Param invoice body service.PostModel true "Post Invoice Model. All fields are required."
// @Success 201 {integer} integer "The new invoice Id"
// @Failure 400 {string} string "Indicates a failure when parsing request body or a validation error, e.g. a required field is missing"
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice [post]
func (controller *InvoiceController) Post(w http.ResponseWriter, r *http.Request) {
	f := handler.InvoicePostHandler(controller.service)
	f(w, r)
}

// Put handles a request to put an Invoice
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
func (controller *InvoiceController) Put(w http.ResponseWriter, r *http.Request) {
	f := handler.InvoiceUpdateHandler(controller.service)
	f(w, r)
}

// Patch handles a request to patch an Invoice
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
func (controller *InvoiceController) Patch(w http.ResponseWriter, r *http.Request) {
	f := handler.InvoiceUpdateHandler(controller.service)
	f(w, r)
}

// Delete handles a request to delete an Invoice
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
func (controller *InvoiceController) Delete(w http.ResponseWriter, r *http.Request) {
	f := handler.InvoiceDeleteHandler(controller.service)
	f(w, r)
}
