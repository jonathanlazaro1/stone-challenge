package providers

import (
	"github.com/jonathanlazaro1/stone-challenge/domain/models"
	pv "github.com/jonathanlazaro1/stone-challenge/interface/providers"
)

type mockedInvoiceRepository struct {
	invoices []models.Invoice
}

func (mp *mockedInvoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]models.Invoice, error) {
	return mp.invoices, nil
}

// MockInvoiceRepository creates a mocked implementation of an InvoiceProvider
func MockInvoiceRepository() pv.InvoiceProvider {
	mockedProvider := &mockedInvoiceRepository{}
	return mockedProvider
}

// Populate puts i invoices inside the mocked repository
func (mp *mockedInvoiceRepository) Populate(i int) {
	for i := 1; i < i; i++ {
		mp.invoices = append(mp.invoices, models.MakeFakeInvoice())
	}
}
