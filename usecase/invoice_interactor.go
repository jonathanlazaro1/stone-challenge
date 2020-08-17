package usecase

import (
	"github.com/jonathanlazaro1/stone-challenge/domain"
)

// InvoiceInteractor handles Invoices requests from and gives responses to the outer layers
type InvoiceInteractor struct {
	repository InvoiceRepository
}

// NewInvoiceInteractor returns a new instance of InvoiceInteractor with the injected repository
func NewInvoiceInteractor(invoiceRepository InvoiceRepository) *InvoiceInteractor {
	return &InvoiceInteractor{
		repository: invoiceRepository,
	}
}

// GetMany returns an array of Invoices, according to the given arguments
func (interactor *InvoiceInteractor) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error) {
	return interactor.repository.GetMany(itemsPerPage, page, filterBy, sortBy)
}

// Get an Invoice, given its Id
func (interactor *InvoiceInteractor) Get(id int) (*domain.Invoice, error) {
	return interactor.repository.Get(id)
}

// Add creates a new Invoice
func (interactor *InvoiceInteractor) Add(invoice domain.Invoice) (int, error) {
	return interactor.repository.Add(invoice)
}

// Update gives new values to an existent Invoice
func (interactor *InvoiceInteractor) Update(invoice domain.Invoice) (int64, error) {
	return interactor.repository.Update(invoice)
}

// Delete deactivates the current Invoice
func (interactor *InvoiceInteractor) Delete(invoice domain.Invoice) (int64, error) {
	invoice.Deactivate()
	return interactor.repository.Update(invoice)
}
