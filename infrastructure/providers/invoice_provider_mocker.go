package providers

import (
	"github.com/jonathanlazaro1/stone-challenge/core/models"
	pv "github.com/jonathanlazaro1/stone-challenge/interface/providers"
)

type mockedInvoiceProvider struct {
	invoices []models.Invoice
}

func (mp *mockedInvoiceProvider) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]models.Invoice, error) {
	return mp.invoices, nil
}

// MockInvoiceProvider creates a mocked implementation of an InvoiceProvider
func MockInvoiceProvider() pv.InvoiceProvider {
	mockedProvider := &mockedInvoiceProvider{}
	return mockedProvider
}

func (mp *mockedInvoiceProvider) PopulateMockedProvider() {
	for i := 0; i < 10; i++ {
		mp.invoices = append(mp.invoices, models.MakeFakeInvoice())
	}
}
