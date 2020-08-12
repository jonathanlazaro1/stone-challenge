package service

import (
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice/test"
)

// BuildInvoiceService builds a new InvoiceInteractor with the specified repository
func BuildInvoiceService() *invoice.Interactor {
	repo := test.MockInvoiceRepository(50)
	service := invoice.NewInteractor(repo)

	return service
}
