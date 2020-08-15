package invoice

import (
	domain "github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"
)

// Interactor handles Invoices requests from and gives responses to the outer layers
type Interactor struct {
	repository repository.Invoice
}

// NewInteractor returns a new instance of InvoiceInteractor with the injected repository
func NewInteractor(invoiceRepository repository.Invoice) *Interactor {
	return &Interactor{
		repository: invoiceRepository,
	}
}

// GetMany returns an array of Invoices, according to the given arguments
func (interactor *Interactor) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error) {
	return interactor.repository.GetMany(itemsPerPage, page, filterBy, sortBy)
}

// Get an Invoice, given its Id
func (interactor *Interactor) Get(id int) (*domain.Invoice, error) {
	return interactor.repository.Get(id)
}

// Add creates a new Invoice
func (interactor *Interactor) Add(invoice domain.Invoice) (int, error) {
	return interactor.repository.Add(invoice)
}

// Update gives new values to an existent Invoice
func (interactor *Interactor) Update(invoice domain.Invoice) (int64, error) {
	return interactor.repository.Update(invoice)
}

// Delete deactivates the current Invoice
func (interactor *Interactor) Delete(invoice domain.Invoice) (int64, error) {
	invoice.Deactivate()
	return interactor.repository.Update(invoice)
}
