package service

import (
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
)

// BuildInvoiceService builds an instance of Invoice
func BuildInvoiceService() *InvoiceService {
	repo := pgsql.GetInvoiceRepository()
	return NewInvoiceService(repo)
}
