package service

import (
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql/repository"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice"
)

// BuildInvoiceService builds a new InvoiceInteractor with the specified repository
func BuildInvoiceService() *invoice.Interactor {
	repo := repository.GetInvoiceRepository()
	service := invoice.NewInteractor(repo)

	return service
}
