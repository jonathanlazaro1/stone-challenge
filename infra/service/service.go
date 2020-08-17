package service

import (
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
)

// BuildInvoiceService builds an instance of Invoice
func BuildInvoiceService() *Invoice {
	repo := pgsql.GetInvoiceRepository()
	return newInvoiceService(repo)
}
