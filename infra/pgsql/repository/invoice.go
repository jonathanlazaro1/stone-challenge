package repository

import (
	"fmt"
	"log"

	"github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"
)

type invoiceRepository struct {
}

// GetInvoiceRepository returns an implementation of InvoiceRepository that relies on a PGSQL DB
func GetInvoiceRepository() repository.Invoice {
	return &invoiceRepository{}
}

// GetMany fetches all invoices found on DB table invoice, according to the parameters given
func (repo *invoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]invoice.Invoice, error) {
	db := pgsql.CreateConnection()
	defer db.Close()

	invoices := []invoice.Invoice{}

	sqlStatement := fmt.Sprintf("SELECT id, reference_year, reference_month, document, description, amount, is_active, created_at, deactivated_at FROM invoice LIMIT %v OFFSET %v;", itemsPerPage, (itemsPerPage * (page - 1)))

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var invoice invoice.Invoice
		err = rows.Scan(
			&invoice.ID,
			&invoice.ReferenceYear,
			&invoice.ReferenceMonth,
			&invoice.Document,
			&invoice.Description,
			&invoice.Amount,
			&invoice.IsActive,
			&invoice.CreatedAt,
			&invoice.DeactivatedAt)

		if err != nil {
			log.Printf("Unable to fetch invoices. %v", err)
			return nil, err
		}

		invoices = append(invoices, invoice)
	}
	return invoices, err
}
