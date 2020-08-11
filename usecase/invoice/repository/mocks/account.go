package mocks

import (
	in "github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	rp "github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"
)

type mockedInvoiceRepository struct {
}

func (mp *mockedInvoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) (*[]in.Invoice, error) {
	invoices := make([]in.Invoice, 1)
	return &invoices, nil
}

// MockInvoiceRepository creates a mocked implementation of an InvoiceProvider
func MockInvoiceRepository() rp.Invoice {
	mockedProvider := &mockedInvoiceRepository{}
	return mockedProvider
}
