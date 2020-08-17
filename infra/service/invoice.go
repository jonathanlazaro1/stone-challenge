package service

import (
	"github.com/jonathanlazaro1/stone-challenge/domain"
	"github.com/jonathanlazaro1/stone-challenge/usecase"
)

// InvoiceService represents a service that handles view layer calls to fetch and modify Invoices
type InvoiceService struct {
	interactor usecase.InvoiceInteractor
}

// NewInvoiceService builds an instance of InvoiceService using the specified InvoiceRepository
func NewInvoiceService(repo usecase.InvoiceRepository) *InvoiceService {
	return &InvoiceService{interactor: *usecase.NewInvoiceInteractor(repo)}
}

// GetMany returns an array of Invoices, according to the given arguments
func (svc *InvoiceService) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error) {
	return svc.interactor.GetMany(itemsPerPage, page, filterBy, sortBy)
}

// Get an Invoice, given its Id
func (svc InvoiceService) Get(id int) (*domain.Invoice, error) {
	return svc.interactor.Get(id)
}

// Add creates a new Invoice
func (svc InvoiceService) Add(invoice domain.Invoice) (int, error) {
	return svc.interactor.Add(invoice)
}

// Update replaces the values on the Invoice held by id, using the new ones that came from request.
// It ignores server-controled ones, such as CreatedAt, IsActive and UpdatedAt
func (svc InvoiceService) Update(id int, postModel PostModel) (*domain.Invoice, error) {
	currentInvoice, err := svc.interactor.Get(id)
	if err != nil {
		return nil, err
	} else if currentInvoice == nil {
		return nil, nil
	}
	updatedInvoice := postModel.ToInvoice(currentInvoice)

	rowCount, err := svc.interactor.Update(updatedInvoice)
	if err != nil {
		return nil, err
	} else if rowCount < 1 {
		return nil, nil
	}

	currentInvoice, err = svc.interactor.Get(id)
	if err != nil {
		return nil, err
	}
	return currentInvoice, nil
}

// Delete makes an Invoice inactive, which makes it unable to be retrieved/maintained
func (svc InvoiceService) Delete(id int) (bool, error) {
	invoice, err := svc.interactor.Get(id)
	if err != nil {
		return false, err
	} else if invoice == nil {
		return false, nil
	}

	rowCount, err := svc.interactor.Delete(*invoice)
	if err != nil {
		return false, err
	} else if rowCount < 1 {
		return false, nil
	}

	return true, nil
}
