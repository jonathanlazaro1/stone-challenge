package invoice

import (
	"github.com/jonathanlazaro1/stone-challenge/domain/models"
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
func (interactor *Interactor) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) (*[]models.Invoice, error) {
	return interactor.repository.GetMany(itemsPerPage, page, filterBy, sortBy)
}
